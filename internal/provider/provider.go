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
	sourceURL       string
	frequency       time.Duration
	feedTransformer FeedTransformer
	feedStorage     storage.FeedStorage
}

func NewFeedProvider(s string, i time.Duration, t FeedTransformer, f storage.FeedStorage) (*FeedProvider, error) {
	if s == "" {
		return nil, fmt.Errorf("source url must be provided")
	}
	if t == nil {
		return nil, fmt.Errorf("feed transformer must be provided")
	}
	if f == nil {
		return nil, fmt.Errorf("feed storage must be provided")
	}

	return &FeedProvider{sourceURL: s, frequency: i, feedTransformer: t, feedStorage: f}, nil
}

func (s FeedProvider) Start() error {
	err := s.iterate()
	if err != nil {
		return err
	}
	return nil
}

func (s FeedProvider) Run(ctx context.Context) {
	for {
		select {
		case <-time.After(s.frequency):
			err := s.iterate()
			if err != nil {
				log.Println("error: %w", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s FeedProvider) iterate() error {
	log.Printf("source loading from %s ...", s.sourceURL)

	source, err := s.load()
	if err != nil {
		return fmt.Errorf("failed to load source: %w", err)
	}

	res := s.feedTransformer.Transform(source)

	err = s.feedStorage.Save(res)
	if err != nil {
		return fmt.Errorf("failed to write source: %w", err)
	}

	log.Println("transformed feed saved to a storage")

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
