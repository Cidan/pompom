package database

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// TODO: implement error channel for read
type Database interface {
	Start(Database)
	Save(context.Context, *pubsub.Message) error
	Read(context.Context) (chan *pubsub.Message, chan error)
}
