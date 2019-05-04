package database

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Read from a flat file, optionally with watch, and construct a pubsub message.

type FlatFile struct {
	file  string
	watch bool
}

func NewFlatFile(file string, watch bool) (*FlatFile, error) {
	return &FlatFile{
		file:  file,
		watch: watch,
	}, nil
}

func (f *FlatFile) Save(ctx context.Context, m *pubsub.Message) error {
	return nil
}

// Ingest a flat file, new line delimited
func (f *FlatFile) Read() (chan *pubsub.Message, error) {
	return nil, nil
}
