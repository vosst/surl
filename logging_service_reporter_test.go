package surl

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLoggerWithBuffer() (*bytes.Buffer, *log.Logger) {
	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "surl ", 0)
	return buffer, logger
}

func TestGetStartPrintsCorrectLine(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}

	reporter.GetStart("lalelu")
	assert.Equal(t, "surl Service.Get start: lalelu\n", b.String(), "Log records mismatch")
}

func TestGetEndPrintsCorrectLineForNilError(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}
	u, _ := url.Parse("http://www.google.com")

	reporter.GetEnd(u, nil)
	assert.Equal(t, fmt.Sprintf("surl Service.Get end: %s\n", u.String()), b.String(), "Log records mismatch")
}

func TestGetEndPrintsCorrectLineForError(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}
	u, _ := url.Parse("http://www.google.com")

	reporter.GetEnd(u, errors.New("test"))
	assert.Equal(t, fmt.Sprintf("surl Service.Get end: %s with error: %s\n", u.String(), "test"), b.String(), "Log records mismatch")
}

func TestPutStartPrintsCorrectLine(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}
	u, _ := url.Parse("http://www.google.com")

	reporter.PutStart(u)
	assert.Equal(t, fmt.Sprintf("surl Service.Put start: %s\n", u.String()), b.String(), "Log records mismatch")
}

func TestPutEndPrintsCorrectLineForNilError(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}

	reporter.PutEnd("lalelu", nil)
	assert.Equal(t, fmt.Sprintf("surl Service.Put end: %s\n", "lalelu"), b.String(), "Log records mismatch")
}

func TestPutEndPrintsCorrectLineForError(t *testing.T) {
	b, l := setupLoggerWithBuffer()
	reporter := LoggingServiceReporter{l}

	reporter.PutEnd("lalelu", errors.New("test"))
	assert.Equal(t, fmt.Sprintf("surl Service.Put end: %s with error: %s\n", "lalelu", "test"), b.String(), "Log records mismatch")
}
