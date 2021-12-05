package storage

import (
	"fmt"
	"sync/atomic"
)

type FeedStorage interface {
	Save(feed string) error
	Get() (string, error)
}

type MemoryFeedStorage struct {
	value atomic.Value
}

func NewMemoryStorage() *MemoryFeedStorage {
	return &MemoryFeedStorage{}
}

func (m *MemoryFeedStorage) Save(feed string) error {
	m.value.Store(feed)
	return nil
}

func (m *MemoryFeedStorage) Get() (string, error) {
	res := m.value.Load()

	if res == nil {
		return "", fmt.Errorf("feed not found")
	}

	return res.(string), nil
}
