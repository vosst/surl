package surl

import (
	"encoding/binary"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testFn = "/tmp/test"

func TestDefaultValueIsReturnedIfFileDoesNotExist(t *testing.T) {
	os.Remove(testFn)

	rw := &FileCounterReaderWriter{testFn}

	c, err := rw.Read(42)
	assert.Nil(t, err)
	assert.Equal(t, uint64(42), c, "Counter value mismatch")
}

func TestCorrectValueIsPersisted(t *testing.T) {
	os.Remove(testFn)

	rw := &FileCounterReaderWriter{testFn}
	assert.Nil(t, rw.Write(42))
	c, err := rw.Read(0)
	assert.Nil(t, err)
	assert.Equal(t, uint64(42), c, "Counter value mismatch")
}

func TestErrorIsReturnedWhenReadingCorruptFile(t *testing.T) {
	refValue := int8(-1)
	// Set up the precondition
	os.Remove(testFn)
	f, err := os.Create(testFn)
	assert.Nil(t, err)
	err = binary.Write(f, binary.LittleEndian, &refValue)
	assert.Nil(t, err)
	f.Close()

	rw := &FileCounterReaderWriter{testFn}
	_, rerr := rw.Read(0)
	t.Log(rerr)
	assert.NotNil(t, rerr)
}
