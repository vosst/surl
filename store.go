package surl

import "net/url"

// Store models surl's assumptions towards a persistent key-value storage component
type Store interface {
	// Get returns the URL known under key, or an error if the lookup failed.
	Get(key string) (*url.URL, error)
	// Put makes url known under key in the store.
	Put(key string, url *url.URL) error
}
