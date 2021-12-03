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

type Replacement struct {
	from string
	to   string
}

func NewReplacement(f, t string) Replacement {
	return Replacement{from: f, to: t}
}

type SimpleTransformer struct {
	replacements []Replacement
}

func NewSimpleTransformer(r []Replacement) *SimpleTransformer {
	return &SimpleTransformer{replacements: r}
}

func (t *SimpleTransformer) Transform(data string) string {
	res := data
	for _, r := range t.replacements {
		res = strings.Replace(res, r.from, r.to, -1)
	}
	return res
}
