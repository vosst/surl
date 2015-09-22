package surl

import "net/url"

// ServiceReporter abstracts reporting of important events happening
// in the service implementation.
type ServiceReporter interface {
	// ReportPutStart marks the begin of a put operation on the service.
	// Returns u.
	PutStart(u *url.URL) *url.URL
	// ReportPutEnd marks the end of a put operation on the service.
	// Returns key and err
	PutEnd(key string, err error) (string, error)

	// ReportGetStarted marks the begin of a get operation on the service.
	// Returns key.
	GetStart(key string) string
	// ReportGetEnd marks the end of a get operation on the service.
	// Returns key.
	GetEnd(u *url.URL, err error) (*url.URL, error)
}
