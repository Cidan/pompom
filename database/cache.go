package database

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/dgraph-io/badger"
	"github.com/rs/zerolog/log"
)

type Cache struct {
	db     *badger.DB
	closed bool
}

func NewCache() *Cache {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/"
	opts.ValueDir = "/tmp/"
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
		err = txn.Set([]byte("yup, sure do"), b.Bytes())
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

func (c *Cache) Read() (chan *pubsub.Message, error) {
	if c.closed {
		return nil, errors.New("database is closed")
	}
	ch := make(chan *pubsub.Message, 1)
	return ch, nil
}

func (c *Cache) Close() error {
	c.closed = true
	return c.db.Close()
}
