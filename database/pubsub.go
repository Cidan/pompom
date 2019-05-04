package database

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Pubsub struct for connecting and publishing to Pubsub
type Pubsub struct {
	client  *pubsub.Client
	topic   *pubsub.Topic
	project string
}

// NewPubsub returns a pubsub struct
func NewPubsub(ctx context.Context, project, topic string) (*Pubsub, error) {
	client, err := pubsub.NewClient(context.Background(), project)

	if err != nil {
		return nil, err
	}

	// Create our topic
	t := client.Topic(topic)
	exists, err := t.Exists(context.Background())

	if err != nil {
		return nil, err
	}

	if !exists {
		_, err = client.CreateTopic(context.Background(), topic)
		if err != nil {
			return nil, err
		}
	}

	return &Pubsub{
		client:  client,
		topic:   t,
		project: project,
	}, nil
}

// Save a message to pubsub
func (p *Pubsub) Save(ctx context.Context, m *pubsub.Message) error {
	res := p.topic.Publish(ctx, m)

	_, err := res.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Read a message from Pubsub
func (p *Pubsub) Read() (chan *pubsub.Message, error) {
	return nil, nil
}
