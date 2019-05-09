package database

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"

	"cloud.google.com/go/pubsub"
	"github.com/fsnotify/fsnotify"
)

// Read from a flat file, optionally with watch, and construct a pubsub message.

type FlatFile struct {
	file  string
	watch bool
}

// TODO: possibly check for file/path here? check for watch bool path?
func NewFlatFile(file string, watch bool) (*FlatFile, error) {
	return &FlatFile{
		file:  file,
		watch: watch,
	}, nil
}

func (f *FlatFile) Start(d Database) {
	c, erch := f.Read()

	for {
		select {
		case m := <-c:
			log.Debug().Msg("saving flat file message to pubsub")
			err := d.Save(context.Background(), m)
			if err != nil {
				log.Panic().Err(err).Msg("error saving message to database")
			}
		case err := <-erch:
			log.Panic().Err(err).Msg("error reading file from disk")
		}
	}
}

func (f *FlatFile) Save(ctx context.Context, m *pubsub.Message) error {
	return nil
}

// Ingest a flat file and create a pubsub message out of it
// with metadata set as attributes. The file is not interpreted
// in any way what so ever.
func (f *FlatFile) Read() (chan *pubsub.Message, chan error) {
	c := make(chan *pubsub.Message, 1)
	erch := make(chan error, 1)
	if f.watch {
		go f.doWatch(c, erch)
	} else {
		go f.doRead(c, erch)
	}
	return c, erch
}

func (f *FlatFile) doRead(c chan *pubsub.Message, erch chan error) {
	m, err := f.readFile(f.file)
	if err != nil {
		erch <- errors.New("unable to read created file")
	}
	c <- m
}

func (f *FlatFile) doWatch(c chan *pubsub.Message, erch chan error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		erch <- errors.New("unable to watch for new files")
		return
	}
	defer watcher.Close()
	watcher.Add(f.file)
	log.Debug().Str("location", f.file).Msg("watching for files on disk")
	for {
		select {
		case event, ok := <-watcher.Events:
			log.Debug().Interface("event", event).Msg("got watch event")
			if !ok {
				erch <- errors.New("filesystem watcher was unable to start watching")
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				m, err := f.readFile(event.Name)
				if err != nil {
					erch <- errors.New("unable to read created file")
				}
				c <- m
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				erch <- errors.New("filesystem watcher was unable to listen for errors")
				return
			}
			erch <- err
		}
	}
}

func (f *FlatFile) readFile(file string) (*pubsub.Message, error) {
	data, err := ioutil.ReadFile(file)
	log.Debug().Str("filename", file).Msg("reading file from disk")
	if err != nil {
		log.Panic().Err(err).Str("file", file).Msg("unable to read created file")
	}
	desc, err := os.Open(file)
	if err != nil {
		log.Panic().Err(err).Str("file", file).Msg("unable to read created file")
	}
	stat, err := desc.Stat()
	if err != nil {
		log.Panic().Err(err).Str("file", file).Msg("unable to read file stats")
	}
	desc.Close()
	return &pubsub.Message{
		Attributes: map[string]string{
			"filename": stat.Name(),
			"size":     string(stat.Size()),
			"modtime":  stat.ModTime().String(),
		},
		Data: data,
	}, nil
}
