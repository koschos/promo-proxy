package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/koschos/promo-proxy/internal/storage"
)

func Wrapper(f *FeedHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f.Handle(w, r)
	}
}

type FeedHandler struct {
	feedStorage storage.FeedStorage
}

func NewFeedHandler(s storage.FeedStorage) (*FeedHandler, error) {
	if s == nil {
		return nil, fmt.Errorf("feed storage must be provided")
	}

	return &FeedHandler{feedStorage: s}, nil
}

func (f *FeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	log.Printf("start fetching feed")

	res, err := f.feedStorage.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("feed storage is empty: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(res))

	log.Printf("feed fetched in %v", time.Since(t))
}
