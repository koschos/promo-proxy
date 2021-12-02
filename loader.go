package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SourceLoader struct {
	sourceURL string
	storage   FeedStorage
	interval  time.Duration
}

func NewSourceLoader(s string, f FeedStorage, i time.Duration) (*SourceLoader, error) {
	if s == "" {
		return nil, fmt.Errorf("source url must be provided")
	}
	if f == nil {
		return nil, fmt.Errorf("storage must be provided")
	}

	return &SourceLoader{sourceURL: s, storage: f, interval: i}, nil
}

func (s SourceLoader) Run(ctx context.Context) {
	for {
		err := s.iterate()
		if err != nil {
			log.Println("error: %w", err)
		}

		select {
		case <-time.After(s.interval):
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (s SourceLoader) iterate() error {
	log.Printf("source loading from %s", s.sourceURL)

	source, err := s.load()
	if err != nil {
		return fmt.Errorf("failed to load source: %w", err)
	}

	err = s.storage.Update(source)
	if err != nil {
		return fmt.Errorf("failed to write source: %w", err)
	}

	log.Println("source saved to storage")

	return nil
}

func (s SourceLoader) load() (string, error) {
	res, err := http.Get(s.sourceURL)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("unable to close body")
		}
	}()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %s", string(b))
	}

	return string(b), nil
}
