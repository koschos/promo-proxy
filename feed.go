package main

import (
	"fmt"
	"sync"
)

type FeedStorage interface {
	Update(feed string) error
	Get() (string, error)
}

type MemoryFeedStorage struct {
	lock *sync.Mutex
	data string
}

func NewMemoryStorage(l *sync.Mutex) *MemoryFeedStorage {
	return &MemoryFeedStorage{lock: l}
}

func (m *MemoryFeedStorage) Update(feed string) error {
	m.lock.Lock()
	m.data = feed
	m.lock.Unlock()
	return nil
}

func (m *MemoryFeedStorage) Get() (string, error) {
	var res string
	m.lock.Lock()
	res = m.data
	m.lock.Unlock()

	if res == "" {
		return "", fmt.Errorf("not found")
	}

	return res, nil
}
