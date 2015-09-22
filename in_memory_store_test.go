package surl

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGetForKnownKeyReturnsURL(t *testing.T) {
	s := NewInMemoryStore()
	k := "lalelu"
	u, _ := url.Parse("http://www.google.com")
	assert.Nil(t, s.Put(k, u))
	v, err := s.Get(k)
	assert.Nil(t, err)
	assert.Equal(t, *u, *v, "Mismatching URLs")
}

func TestGetForUnknownKeyReturnsError(t *testing.T) {
	s := NewInMemoryStore()
	k := "lalelu"

	u, _ := s.Get(k)
	assert.Nil(t, u)
}

func TestPutPersistsURLForKey(t *testing.T) {
	s := NewInMemoryStore()
	k := "lalelu"
	u, _ := url.Parse("http://www.google.com")

	assert.Nil(t, s.Put(k, u))
	v, err := s.Get(k)
	assert.Nil(t, err)
	assert.Equal(t, *u, *v, "Mismatching URLs")
}

func TestPutForSameKeySucceeds(t *testing.T) {
	s := NewInMemoryStore()
	k := "lalelu"
	u, _ := url.Parse("http://www.google.com")

	assert.Nil(t, s.Put(k, u))
	assert.Nil(t, s.Put(k, u))
}
