package main

import (
	"context"
	"log"
	"net/http"
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

var feedStorage storage.FeedStorage
var feedTransformer provider.FeedTransformer

func init() {
	feedStorage = storage.NewMemoryStorage()

	replacements := []provider.Replacement{
		provider.NewReplacement("available=\"false\"", "available=\"\""),
		provider.NewReplacement("available-kiev=\"false\"", "available-kiev=\"\""),
		provider.NewReplacement("available-kharkov=\"false\"", "available-kharkov=\"\""),
	}
	feedTransformer = provider.NewSimpleTransformer(replacements)
}

func main() {
	feedProvider, err := provider.NewFeedProvider(sourceURL, frequency, feedTransformer, feedStorage)
	if err != nil {
		log.Fatalln("failed to create feed provider")
	}

	startErr := feedProvider.Start()
	if startErr != nil {
		log.Fatalln("failed to start feed provider: %w", startErr)
	}

	ctx := context.Background()

	go feedProvider.Run(ctx)

	feedHandler, err := feedhttp.NewFeedHandler(feedStorage)
	if err != nil {
		log.Fatalf("failed to create feed handler: %v", err)
	}

	http.HandleFunc("/", feedhttp.Wrapper(feedHandler))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
