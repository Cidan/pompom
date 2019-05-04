package database

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestNewPubsub(t *testing.T) {
	p, err := NewPubsub(
		context.Background(),
		"test",
		"test",
	)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestPubsub_Save(t *testing.T) {
	p, err := NewPubsub(
		context.Background(),
		"test",
		"test",
	)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	m := &pubsub.Message{
		Data: []byte("test"),
	}

	err = p.Save(context.Background(), m)
	assert.Nil(t, err)
}

func TestPubsub_Read(t *testing.T) {
}
