package main

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"github.com/vosst/surl"
	"net/http"
	"os"
	"testing"
)

const testFn = "/tmp/test"

func TestRestTicketerCanAccessService(t *testing.T) {
	// Make sure that we have a sensible value available to the
	// FileCounterReaderWriter
	{
		os.Remove(testFn)
		f, _ := os.Create(testFn)
		binary.Write(f, binary.LittleEndian, uint64(41))
		f.Close()
	}

	rst := &RestTicketerService{}
	go rst.Main(":9091", &surl.FileCounterReaderWriter{testFn})

	rt := &surl.RestTicketer{"http://localhost:9091", &http.Client{}}
	assert.Equal(t, "42", rt.Next(), "Ticket mismatch")
}
