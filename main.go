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
	sourceURL = "https://intertool.ua/xml_output/yandex_market.xml"
	frequency = 60 * time.Minute
	port      = "3000"
)

var feedStorage storage.FeedProxyStorage

func init() {
	lock := &sync.Mutex{}
	replacements := []storage.Replacement{
		storage.NewReplacement("available=\"false\"", "available=\"\""),
		storage.NewReplacement("available-kiev=\"false\"", "available-kiev=\"\""),
		storage.NewReplacement("available-kharkov=\"false\"", "available-kharkov=\"\""),
	}

	feedStorage = storage.NewMemoryStorage(lock, storage.NewSimpleTransformer(replacements))
}

func main() {
	ctx := context.Background()

	feedProvider, err := provider.NewFeedProvider(sourceURL, feedStorage, frequency)
	if err != nil {
		log.Fatalln("failed to create feed provider")
	}

	startErr := feedProvider.Start()
	if startErr != nil {
		log.Fatalln("failed to start feed provider")
	}

	go feedProvider.Run(ctx)

	feedHandler, err := feedhttp.NewFeedHandler(feedStorage)
	if err != nil {
		log.Fatalf("failed to create feed handler: %w", err)
	}

	http.HandleFunc("/", feedhttp.Wrapper(feedHandler))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
