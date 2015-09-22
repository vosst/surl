package surl

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

type MockServiceReporter struct {
	mock.Mock
}

func (self MockServiceReporter) PutStart(u *url.URL) *url.URL {
	args := self.Called(u)
	return args.Get(0).(*url.URL)
}

func (self MockServiceReporter) PutEnd(key string, err error) (string, error) {
	args := self.Called(key, err)
	return args.String(0), args.Error(1)
}

func (self MockServiceReporter) GetStart(key string) string {
	args := self.Called(key)
	return args.String(0)
}

func (self MockServiceReporter) GetEnd(u *url.URL, err error) (*url.URL, error) {
	args := self.Called(u, err)
	return args.Get(0).(*url.URL), args.Error(1)
}
