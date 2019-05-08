package database

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	c := NewCache()
	err := c.Close()
	assert.Nil(t, err)
}

func TestCache_Save(t *testing.T) {
	c := NewCache()
	err := c.Close()
	assert.Nil(t, err)

	c.Save(
		context.Background(),
		&pubsub.Message{},
	)
}

func TestCache_Read(t *testing.T) {
}
