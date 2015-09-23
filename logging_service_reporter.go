package surl

import (
	"log"
	"net/url"
)

// LoggingServiceReporter outputs Service events via a log.Logger instance
type LoggingServiceReporter struct {
	Logger *log.Logger
}

func (self LoggingServiceReporter) PutStart(u *url.URL) *url.URL {
	self.Logger.Printf("Service.Put start: %s", u.String())
	return u
}

func (self LoggingServiceReporter) PutEnd(key string, err error) (string, error) {
	if err != nil {
		self.Logger.Printf("Service.Put end: %s with error: %s", key, err)
	} else {
		self.Logger.Printf("Service.Put end: %s", key)
	}

	return key, err
}

func (self LoggingServiceReporter) GetStart(key string) string {
	self.Logger.Printf("Service.Get start: %s", key)
	return key
}

func (self LoggingServiceReporter) GetEnd(u *url.URL, err error) (*url.URL, error) {
	if err != nil {
		self.Logger.Printf("Service.Get end: %s with error: %s", u, err)
	} else {
		self.Logger.Printf("Service.Get end: %s", u.String())
	}

	return u, err
}
