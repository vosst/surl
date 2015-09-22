package surl

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStartPrintsCorrectLine(t *testing.T) {
	buffer := bytes.Buffer{}

	logger := log.New(&buffer, "surl ", 0)
	reporter := LoggingServiceReporter{logger}

	reporter.GetStart("lalelu")
	assert.Equal(t, "surl Service.Get start: lalelu", buffer.String(), "Log records mismatch")
	fmt.Print(buffer.String())
}
