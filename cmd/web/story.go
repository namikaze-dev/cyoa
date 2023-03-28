package main

import (
	"encoding/json"
	"io"
)

type Story struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func NewStoryStore(r io.Reader) (map[string]Story, error) {
	store := map[string]Story{}
	err := json.NewDecoder(r).Decode(&store)
	return store, err
}
