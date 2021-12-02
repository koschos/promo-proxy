package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	feedhttp "github.com/koschos/promo-proxy/internal/http"
	"github.com/koschos/promo-proxy/internal/provider"
	"github.com/koschos/promo-proxy/internal/storage"
)

const (
	sourceURL = "https://ukr-net1.webnode.com.ua/rss/novini.xml"
	frequency = 10 * time.Minute
	port      = "3000"
)

var feedStorage storage.FeedStorage

func init() {
	lock := &sync.Mutex{}
	feedStorage = storage.NewMemoryStorage(lock)
}

func main() {
	feedHandler, err := feedhttp.NewFeedHandler(feedStorage)
	if err != nil {
		log.Fatalf("failed to create feed handler: %w", err)
	}

	feedProvider, err := provider.NewFeedProvider(sourceURL, feedStorage, frequency)
	if err != nil {
		log.Fatalln("failed to init feed provider")
	}

	ctx := context.Background()

	go feedProvider.Run(ctx)

	http.HandleFunc("/", feedhttp.Wrapper(feedHandler))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
