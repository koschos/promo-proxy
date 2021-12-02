package storage

import (
	"fmt"
	"sync"
)

type FeedProxyStorage interface {
	Update(feed string) error
	Get() (string, error)
}

type MemoryFeedStorage struct {
	lock        *sync.Mutex
	transformer FeedTransformer
	source      string
	target      string
}

func NewMemoryStorage(l *sync.Mutex, t FeedTransformer) *MemoryFeedStorage {
	return &MemoryFeedStorage{lock: l, transformer: t}
}

func (m *MemoryFeedStorage) Update(feed string) error {
	m.lock.Lock()
	m.source = feed
	m.target = m.transformer.Transform(feed)
	m.lock.Unlock()
	return nil
}

func (m *MemoryFeedStorage) Get() (string, error) {
	var res string
	m.lock.Lock()
	res = m.target
	m.lock.Unlock()

	if res == "" {
		return "", fmt.Errorf("target not found")
	}

	return res, nil
}
