package surl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipleInvocationsReturnUniqueValues(t *testing.T) {
	var empty struct{}

	ct := &CountingTicketer{}
	store := make(map[string]*struct{})

	for i := 0; i < 10; i++ {
		k := ct.Next()
		assert.Nil(t, store[k])
		store[k] = &empty
	}
}
