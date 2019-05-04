package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlatFile(t *testing.T) {
	ff, err := NewFlatFile("path/to/file", false)
	assert.Nil(t, err)
	assert.NotNil(t, ff)
}

func TestFlatFile_Save(t *testing.T) {
}

func TestFlatFile_Read(t *testing.T) {
}

func TestFlatFile_doRead(t *testing.T) {
}

func TestFlatFile_doWatch(t *testing.T) {
}

func TestFlatFile_readFile(t *testing.T) {
}
