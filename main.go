package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	url  = "https://ukr-net1.webnode.com.ua/rss/novini.xml"
	port = "3000"
)

var storage FeedStorage

func init() {
	lock := &sync.Mutex{}
	storage = NewMemoryStorage(lock)
}

func main() {
	feedHandler, err := NewFeedHandler(storage)
	if err != nil {
		log.Fatalf("failed to create feed handler: %w", err)
	}

	sourceLoader, err := NewSourceLoader(url, storage, 10*time.Second)
	if err != nil {
		log.Fatalln("failed to init source loader")
	}

	ctx := context.Background()

	go sourceLoader.Run(ctx)

	http.HandleFunc("/", Wrapper(feedHandler))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
