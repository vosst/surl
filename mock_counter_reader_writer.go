package surl

import (
	"github.com/stretchr/testify/mock"
)

type MockCounterReaderWriter struct {
	mock.Mock
}

func (self *MockCounterReaderWriter) Read(defaultValue uint64) (uint64, error) {
	args := self.Called(defaultValue)
	return args.Get(0).(uint64), args.Error(1)
}

func (self *MockCounterReaderWriter) Write(value uint64) error {
	args := self.Called(value)
	return args.Error(0)
}
