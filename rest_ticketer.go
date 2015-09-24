package surl

import (
	"encoding/json"
	"net/http"
)

const (
	// NextEndpoint is the name of the endpoint for sampling
	// the next key
	RestTicketerNextEndpoint = "/next"
)

type RestTicketerDocument struct {
	Next string
}

// RestTicketer implements Ticketer, taling to a web-service for generating tickets.
type RestTicketer struct {
	Host   string       // stem of the URL
	Client *http.Client // client instance for connecting to the remote end
}

// Next returns the next ticket, or an error communication with the other side fails.
func (self *RestTicketer) Next() string {
	if resp, err := self.Client.Get(self.Host + RestTicketerNextEndpoint); err != nil {
		return ""
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return ""
		}

		document := RestTicketerDocument{}
		if err = json.NewDecoder(resp.Body).Decode(&document); err != nil {
			return ""
		}

		return document.Next
	}

}
