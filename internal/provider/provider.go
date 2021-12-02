package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/koschos/promo-proxy/internal/storage"
)

type FeedProvider struct {
	sourceURL   string
	feedStorage storage.FeedProxyStorage
	frequency   time.Duration
}

func NewFeedProvider(s string, f storage.FeedProxyStorage, i time.Duration) (*FeedProvider, error) {
	if s == "" {
		return nil, fmt.Errorf("source url must be provided")
	}
	if f == nil {
		return nil, fmt.Errorf("feedStorage must be provided")
	}

	return &FeedProvider{sourceURL: s, feedStorage: f, frequency: i}, nil
}

func (s FeedProvider) Run(ctx context.Context) {
	for {
		err := s.iterate()
		if err != nil {
			log.Println("error: %w", err)
		}

		select {
		case <-time.After(s.frequency):
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (s FeedProvider) iterate() error {
	log.Printf("source loading from %s", s.sourceURL)

	source, err := s.load()
	if err != nil {
		return fmt.Errorf("failed to load source: %w", err)
	}

	err = s.feedStorage.Update(source)
	if err != nil {
		return fmt.Errorf("failed to write source: %w", err)
	}

	log.Println("source saved to feedStorage")

	return nil
}

func (s FeedProvider) load() (string, error) {
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
