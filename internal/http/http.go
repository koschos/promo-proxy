package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/koschos/promo-proxy/internal/storage"
)

func Wrapper(f *FeedHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f.Handle(w, r)
	}
}

type FeedHandler struct {
	feedStorage storage.FeedProxyStorage
}

func NewFeedHandler(s storage.FeedProxyStorage) (*FeedHandler, error) {
	if s == nil {
		return nil, fmt.Errorf("feed storage must be provided")
	}

	return &FeedHandler{feedStorage: s}, nil
}

func (f *FeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("feed called")

	res, err := f.feedStorage.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("feedStorage is empty, no feed to provide")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(res))
}
