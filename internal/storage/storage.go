package storage

import (
	"fmt"
	"sync"
)

type FeedStorage interface {
	Save(feed string) error
	Get() (string, error)
}

type MemoryFeedStorage struct {
	lock *sync.Mutex
	feed string
}

func NewMemoryStorage(l *sync.Mutex) *MemoryFeedStorage {
	return &MemoryFeedStorage{lock: l}
}

func (m *MemoryFeedStorage) Save(feed string) error {
	m.lock.Lock()
	m.feed = feed
	m.lock.Unlock()
	return nil
}

func (m *MemoryFeedStorage) Get() (string, error) {
	var res string
	m.lock.Lock()
	res = m.feed
	m.lock.Unlock()

	if res == "" {
		return "", fmt.Errorf("feed not found")
	}

	return res, nil
}
