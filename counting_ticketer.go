package surl

import "fmt"

// CountingTicketer maintains an incrementing counter.
type CountingTicketer struct {
	counter uint64
}

// Next returns the next available unique key.
func (self *CountingTicketer) Next() string {
	self.counter++
	return fmt.Sprint(self.counter)
}
