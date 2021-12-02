package storage

import "strings"

type FeedTransformer interface {
	Transform(data string) string
}

type DummyFeedTransformer struct {
}

func (t *DummyFeedTransformer) Transform(data string) string {
	return strings.Replace(data, "ukr-net1", "ukrnet", 100)
}
