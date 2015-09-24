package surl

import (
	"fmt"
)

// WriteAheadTicketer implements Ticketer, maintaining a counter value
// in a file, atomically updating it on calls to Next() prior to handing
// out the next key.
type WriteAheadTicketer struct {
	rw      CounterReaderWriter // Persistence mechanism
	counter uint64              // Current counter value.
}

// NewWriteAheadTicketer creates a new WriteAheadTicketer instance,
// reading the last known counter value from the file named fn. If an
// error occurs while trying to read the file, creation is aborted and
// an error is returned.
func NewWriteAheadTicketer(rw CounterReaderWriter) (*WriteAheadTicketer, error) {
	if c, err := rw.Read(0); err != nil {
		return nil, err
	} else {
		return &WriteAheadTicketer{rw, c}, nil
	}
}

// Next() returns a new key, persisting counter state prior to returning.
// Returns the empty string in case of errors.
func (self *WriteAheadTicketer) Next() string {
	if err := self.rw.Write(self.counter + 1); err != nil {
		// TODO(tvoss): Ticketer should allow for reporting errors.
		return ""
	}

	self.counter++
	return fmt.Sprint(self.counter)
}
