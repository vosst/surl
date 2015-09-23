package surl

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

// WriteAheadTicketer implements Ticketer, maintaining a counter value
// in a file, atomically updating it on calls to Next() prior to handing
// out the next key.
type WriteAheadTicketer struct {
	fn      string // File that contains the last counter value.
	counter uint64 // Current counter value.
}

// NewWriteAheadTicketer creates a new WriteAheadTicketer instance,
// reading the last known counter value from the file named fn. If an
// error occurs while trying to read the file, creation is aborted and
// an error is returned.
func NewWriteAheadTicketer(fn string) (*WriteAheadTicketer, error) {
	var initialCounterValue uint64
	if f, err := os.Open(fn); err == nil {
		if err = binary.Read(f, binary.LittleEndian, &initialCounterValue); err != nil {
			// There is a previous value that we have issues reading.
			// With that, we would risk consistency and instead bail out
			// here.
			return nil, err
		}
	}

	return &WriteAheadTicketer{fn, initialCounterValue}, nil
}

// Next() returns a new key, persisting counter state prior to returning.
// Returns the empty string in case of errors.
func (self *WriteAheadTicketer) Next() string {
	f, err := ioutil.TempFile("", "SurlTemporaryState")

	if err != nil {
		// TODO(tvoss): Add error handling capabilities to
		// Ticketer interface.
		return ""
	}

	tempFileName := f.Name()
	err = binary.Write(f, binary.LittleEndian, self.counter+1)
	f.Close()

	if err != nil {
		// TODO(tvoss): Add error handling capabilities to
		// Ticketer interface.
		return ""
	}

	if err = os.Rename(tempFileName, self.fn); err != nil {
		// TODO(tvoss): Add error handling capabilities to
		// Ticketer interface.
		return ""
	}

	self.counter++
	return fmt.Sprint(self.counter)
}
