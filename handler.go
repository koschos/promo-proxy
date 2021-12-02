package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func Wrapper(f *FeedHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f.Handle(w, r)
	}
}

type FeedHandler struct {
	storage FeedStorage
}

func NewFeedHandler(s FeedStorage) (*FeedHandler, error) {
	if s == nil {
		return nil, fmt.Errorf("storage must be provided")
	}

	return &FeedHandler{storage: s}, nil
}

func (f *FeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("feed called")

	res, err := f.storage.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("storage is empty, no feed to provide")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(res))
}

func (f *FeedHandler) buildFeed() ([]byte, error) {
	type Item struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
	}

	type rss struct {
		Version     string `xml:"version,attr"`
		Description string `xml:"channel>description"`
		Link        string `xml:"channel>link"`
		Title       string `xml:"channel>title"`

		Item []Item `xml:"channel>item"`
	}

	articles := []Item{
		{"foo", "http://mywebsite.com/foo", "lorem ipsum"},
		{"foo2", "http://mywebsite.com/foo2", "lorem ipsum2"},
	}

	feed := &rss{
		Version:     "2.0",
		Description: "My super website",
		Link:        "http://mywebsite.com",
		Title:       "Mywebsite",
		Item:        articles,
	}

	x, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return nil, err
	}

	return x, nil
}
