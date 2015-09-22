package surl

import (
	"fmt"
	"net/url"
)

// Service bundles together URL shortening and persistent mapping
// of short to long URLs
type Service struct {
	ticketer Ticketer // Used to create a unique id for an incoming URL
	store    Store    // Used to store short -> long url mappings
}

// ErrorGettingURL indicates issues resolving a URL for key.
type ErrorGettingURL struct {
	Key   string // key for which the get operation failed
	Inner error  // error reported by the underlying store implementation
}

// Error pretty prints an ErrorGettingURL instance and implements go's error interface.
func (self ErrorGettingURL) Error() string {
	return fmt.Sprint("Could not get URL from store for key: " + self.Key + self.Inner.Error())
}

// ErrorPuttingURL indicates issues with storing a URL under a key.
type ErrorPuttingURL struct {
	Key   string // key for which the put operation failed
	Inner error  // error reported by the underlying store implementation
}

// Error pretty prints an ErrorPuttingURL instance and implements go's error interface.
func (self ErrorPuttingURL) Error() string {
	return fmt.Sprint("Could not put URL into store: " + self.Key + self.Inner.Error())
}

// Get resolves the long URL corresponding to the URL s.
func (self Service) Get(key string) (*url.URL, error) {
	if l, err := self.store.Get(key); err != nil {
		return nil, ErrorGettingURL{Key: key, Inner: err}
	} else {
		return l, nil
	}
}

// Put shortens the URL l and persists the mapping from short url to l.
func (self Service) Put(l *url.URL) (string, error) {
	s := self.ticketer.Next()
	if err := self.store.Put(s, l); err != nil {
		return "", ErrorPuttingURL{Key: s, Inner: err}
	}

	return s, nil
}

func NewService(ticketer Ticketer, store Store) *Service {
	return &Service{ticketer: ticketer, store: store}
}
