package database

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	c := NewCache("/tmp/badger.new")
	err := c.Close()
	assert.Nil(t, err)
}

func TestCache_Save(t *testing.T) {
	c := NewCache("/tmp/badger.save")
	assert.NotNil(t, c)

	m := &pubsub.Message{
		Data: []byte("this data"),
	}

	c.Save(context.Background(), m)

	err := c.Close()
	assert.Nil(t, err)
}

func TestCache_Read(t *testing.T) {
	c := NewCache("/tmp/badger.read")
	assert.NotNil(t, c)
	m := &pubsub.Message{
		Data: []byte("this data"),
	}
	c.Save(context.Background(), m)
	ch, erch := c.Read()
	assert.NotNil(t, erch)
	assert.NotNil(t, ch)
	mr := <-ch
	assert.Equal(t, mr.Data, m.Data)
}
