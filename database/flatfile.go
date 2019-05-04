package database

import (
	"context"
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

func (f *FlatFile) Save(ctx context.Context, m *pubsub.Message) error {
	return nil
}

// Ingest a flat file and create a pubsub message out of it
// with metadata set as attributes. The file is not interpreted
// in any way what so ever.
func (f *FlatFile) Read() (chan *pubsub.Message, error) {
	c := make(chan *pubsub.Message, 1)
	if f.watch {
		go f.doWatch(c)
	} else {
		go f.doRead(c)
	}
	return c, nil
}

func (f *FlatFile) doRead(c chan *pubsub.Message) {
	m, err := f.readFile(f.file)
	if err != nil {
		log.Panic().Err(err).Str("file", f.file).Msg("unable to read created file")
	}
	c <- m
}

func (f *FlatFile) doWatch(c chan *pubsub.Message) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic().Err(err).Msg("unable to watch for new files")
		return
	}
	defer watcher.Close()
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Panic().Msg("filesystem watcher was unable to start watching")
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				m, err := f.readFile(event.Name)
				if err != nil {
					log.Panic().Err(err).Str("file", event.Name).Msg("unable to read created file")
				}
				c <- m
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				log.Panic().Msg("filesystem watcher was unable to listen for errors")
				return
			}
			log.Panic().Err(err).Msg("error while watching for new files")
		}
	}
}

func (f *FlatFile) readFile(file string) (*pubsub.Message, error) {
	data, err := ioutil.ReadFile(file)
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
