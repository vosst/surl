package surl

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (self *MockStore) Get(key string) (*url.URL, error) {
	args := self.Called(key)
	return args.Get(0).(*url.URL), args.Error(1)
}

func (self *MockStore) Put(key string, url *url.URL) error {
	args := self.Called(key, url)
	return args.Error(0)
}
