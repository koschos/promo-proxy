package main

import (
	"encoding/xml"
	"net/http"
)

func rssHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
