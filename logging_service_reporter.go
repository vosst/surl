package surl

import (
	"log"
	"net/url"
)

type LoggingServiceReporter struct {
	Logger *log.Logger
}

func (self LoggingServiceReporter) PutStart(u *url.URL) *url.URL {
	self.Logger.Printf("Service.Put start: %s", u)
	return u
}

func (self LoggingServiceReporter) PutEnd(key string, err error) (string, error) {
	self.Logger.Printf("Service.Put end: %s with error: %s", key, err)
	return key, err
}

func (self LoggingServiceReporter) GetStart(key string) string {
	self.Logger.Printf("Service.Get start: %s", key)
	return key
}

func (self LoggingServiceReporter) GetEnd(u *url.URL, err error) (*url.URL, error) {
	self.Logger.Printf("Service.Get end: %s with error: %s", u, err)
	return u, err
}
