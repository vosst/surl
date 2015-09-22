package surl

import "github.com/stretchr/testify/mock"

type MockTicketer struct {
	mock.Mock
}

func (self *MockTicketer) Next() string {
	args := self.Called()
	return args.String(0)
}
