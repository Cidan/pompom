package database

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFlatFile(t *testing.T) {
	ff, err := NewFlatFile("path/to/file", false)
	assert.Nil(t, err)
	assert.NotNil(t, ff)
}

func TestFlatFile_Save(t *testing.T) {
	// TODO: implement
}

func TestFlatFile_Read(t *testing.T) {
	ff, err := NewFlatFile("../fixtures/fake-file", false)
	assert.Nil(t, err)
	assert.NotNil(t, ff)

	c, err := ff.Read()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	m := <-c
	assert.Equal(t, string(m.Data), "fake-data")
	assert.Equal(t, m.Attributes["filename"], "fake-file")

	// Test watcher
	ff, err = NewFlatFile("../fixtures", true)
	assert.Nil(t, err)
	assert.NotNil(t, ff)

	c, err = ff.Read()
	assert.Nil(t, err)
	assert.NotNil(t, ff)

	// TODO: wait for watcher to be ready, don't sleep
	time.Sleep(time.Second * 1)
	err = createFile("../fixtures/watched")
	assert.Nil(t, err)

	m = <-c
	assert.Equal(t, m.Attributes["filename"], "watched")
}

func TestFlatFile_doRead(t *testing.T) {
}

func TestFlatFile_doWatch(t *testing.T) {
}

func TestFlatFile_readFile(t *testing.T) {
}

// Create a file in /tmp and then move it to our dest
func createFile(file string) error {

	now := time.Now().UnixNano()
	name := fmt.Sprintf("../tmp/watched-%d", now)

	if _, err := os.Stat("../tmp"); os.IsNotExist(err) {
		err = os.Mkdir("../tmp", os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if _, err := os.Stat(file); err == nil {
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}

	err = os.Rename(name, file)
	if err != nil {
		return err
	}

	return nil
}
