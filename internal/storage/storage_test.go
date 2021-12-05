package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	storage := NewMemoryStorage()

	noRes, err := storage.Get()
	assert.Equal(t, "", noRes)
	assert.Errorf(t, err, "feed not found")

	feed := "<xml>feed</xml>"
	saveErr := storage.Save(feed)
	assert.NoError(t, saveErr)

	res, getErr := storage.Get()
	assert.NoError(t, getErr)
	assert.Equal(t, feed, res)
}
