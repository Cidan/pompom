package database

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/dgraph-io/badger"
	"github.com/rs/zerolog/log"
)

type Cache struct {
	db     *badger.DB
	closed bool
}

func NewCache(location string) *Cache {
	opts := badger.DefaultOptions
	opts.Dir = location
	opts.ValueDir = location
	db, err := badger.Open(opts)
	if err != nil {
		log.Panic().Err(err).Msg("unable to open cache")
	}
	return &Cache{
		db:     db,
		closed: false,
	}
}

func (c *Cache) Save(ctx context.Context, m *pubsub.Message) error {
	if c.closed {
		return errors.New("database is closed")
	}
	err := c.db.Update(func(txn *badger.Txn) error {
		var b bytes.Buffer
		e := gob.NewEncoder(&b)
		err := e.Encode(m)
		if err != nil {
			return err
		}
		bt := make([]byte, 8)
		binary.LittleEndian.PutUint64(bt, uint64(time.Now().UnixNano()))
		err = txn.Set(bt, b.Bytes())
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// TODO: clean this up a bit
func (c *Cache) Read() (chan *pubsub.Message, chan error) {
	okch := make(chan *pubsub.Message, 1)
	erch := make(chan error)
	go func() {
		if c.closed {
			erch <- errors.New("database is closed")
			return
		}
		err := c.db.Update(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			opts.PrefetchSize = 10
			it := txn.NewIterator(opts)
			defer it.Close()
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				v, err := item.Value()
				if err != nil {
					return err
				}
				var m pubsub.Message
				var b bytes.Buffer
				b.Write(v)
				d := gob.NewDecoder(&b)
				d.Decode(&m)
				okch <- &m
				txn.Delete(item.Key())
			}
			return nil
		})
		if err != nil {
			erch <- err
			return
		}
	}()
	return okch, erch
}

func (c *Cache) Close() error {
	c.closed = true
	return c.db.Close()
}
