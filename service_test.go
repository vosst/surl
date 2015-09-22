package surl

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceCallsIntoTicketerAndStoreOnPut(t *testing.T) {
	u, _ := url.Parse("http://www.google.com")

	mt := &MockTicketer{}
	mt.On("Next").Return("lalelu")

	ms := &MockStore{}
	ms.On("Get", "lalelu").Return(u, nil)
	ms.On("Put", "lalelu", u).Return(nil)

	mr := &MockServiceReporter{}
	mr.On("PutStart", u).Return(u)
	mr.On("PutEnd", "lalelu", nil).Return("lalelu", nil)
	mr.On("GetStart", "lalelu").Return("lalelu")
	mr.On("GetEnd", u, nil).Return(u, nil)

	service := NewService(mt, ms, mr)

	service.Put(u)
	v, _ := service.Get("lalelu")

	assert.Equal(t, *u, *v, "Mismatching URLs")
	mt.AssertExpectations(t)
}
