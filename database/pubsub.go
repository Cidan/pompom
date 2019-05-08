package database

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
)

// Pubsub struct for connecting and publishing to Pubsub
type Pubsub struct {
	client  *pubsub.Client
	topic   *pubsub.Topic
	project string
	cache   *Cache
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
		cache:   NewCache(),
	}, nil
}

func (p *Pubsub) Start(d Database) {

}

// Save a message to pubsub
func (p *Pubsub) Save(ctx context.Context, m *pubsub.Message) error {
	res := p.topic.Publish(ctx, m)
	log.Debug().Msg("publishing message to pubsub")

	_, err := res.Get(ctx)
	if err != nil {
		log.Debug().Msg("unable to publish message, storing in cache")
		p.cache.Save(context.Background(), m)
		return err
	}
	log.Debug().Msg("message sent")
	return nil
}

// Read a message from Pubsub
func (p *Pubsub) Read(ctx context.Context) (chan *pubsub.Message, error) {
	return nil, nil
}
