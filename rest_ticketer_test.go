package surl

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestTicketerCallsIntoRemoteEnd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/next" {
			json.NewEncoder(w).Encode(RestTicketerDocument{"42"})
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	}))
	defer server.Close()

	rt := &RestTicketer{server.URL, &http.Client{}}
	assert.Equal(t, "42", rt.Next())
}
