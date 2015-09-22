package surl

import (
	"log"
	"net/url"
	"os"
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

	service := NewService(mt, ms, log.New(os.Stdout, "surl ", log.LstdFlags))

	service.Put(u)
	v, _ := service.Get("lalelu")

	assert.Equal(t, *u, *v, "Mismatching URLs")
	mt.AssertExpectations(t)
}
