package surl

import (
	"encoding/binary"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testFn = "/tmp/test"

func TestCreationForValidFileSucceeds(t *testing.T) {
	os.Remove(testFn)

	ticketer, err := NewWriteAheadTicketer(testFn)
	assert.Nil(t, err)
	assert.NotNil(t, ticketer)
}

func TestNewlyCreateTicketerReturnsOneAsFirstValue(t *testing.T) {
	os.Remove(testFn)

	ticketer, _ := NewWriteAheadTicketer(testFn)
	assert.Equal(t, "1", ticketer.Next(), "Key mismatch")
}

func TestTicketerPersistsCounterValue(t *testing.T) {
	os.Remove(testFn)

	ticketer, _ := NewWriteAheadTicketer(testFn)
	ticketer.Next()

	counter := uint64(0)
	f, ferr := os.Open(testFn)

	assert.Nil(t, ferr)
	assert.NotNil(t, f)

	binary.Read(f, binary.LittleEndian, &counter)
	assert.Equal(t, uint64(1), counter, "Counter value mismatch")
}
