package surl

import (
	"errors"
	"net/url"
)

type InMemoryStore struct {
	mapping map[string]*url.URL
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{mapping: make(map[string]*url.URL)}
}

func (self InMemoryStore) Get(key string) (*url.URL, error) {
	return self.mapping[key], nil
}

func (self InMemoryStore) Put(key string, url *url.URL) error {
	if url == nil {
		return errors.New("Cannot map to empty URL")
	}

	self.mapping[key] = url
	return nil
}
