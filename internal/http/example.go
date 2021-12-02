package http

import "encoding/xml"

func buildFeed() ([]byte, error) {
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
