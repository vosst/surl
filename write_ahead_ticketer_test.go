package surl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const tmpTest = "/tmp/test"

func TestNewReadsInitialValueFromCounterReaderWriter(t *testing.T) {
	mc := &MockCounterReaderWriter{}
	mc.On("Read", uint64(0)).Return(uint64(0), nil)

	_, err := NewWriteAheadTicketer(mc)
	assert.Nil(t, err)

	mc.AssertExpectations(t)
}

func TestNextWritesValue(t *testing.T) {
	mc := &MockCounterReaderWriter{}
	mc.On("Read", uint64(0)).Return(uint64(0), nil)
	mc.On("Write", uint64(1)).Return(nil)

	ticketer, _ := NewWriteAheadTicketer(mc)

	assert.Equal(t, "1", ticketer.Next(), "Ticket mismatch")
	mc.AssertExpectations(t)
}
