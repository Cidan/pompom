package database

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Database interface {
	Save(context.Context, *pubsub.Message) error
	Read(context.Context) (chan *pubsub.Message, error)
}
