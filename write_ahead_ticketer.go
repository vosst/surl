package surl

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

// CounterReaderWriter abstracts reading and writing of a single
// uint64 counter value.
type CounterReaderWriter interface {
	// Read tries to read the counter value from the underlying storage solution.
	// If there exists to previous value, defaultValue is returned.
	Read(defaultValue uint64) (uint64, error)
	// Write stores the value of counter to the underlying storage solution.
	Write(counter uint64) error
}

// FileCounterReaderWriter implements CounterReaderWriter, relying on ordinary
// files as the persistence backend.
type FileCounterReaderWriter struct {
	fn string
}

func (self *FileCounterReaderWriter) Read(defaultValue uint64) (uint64, error) {
	value := defaultValue
	if f, err := os.Open(self.fn); err == nil {
		if err = binary.Read(f, binary.LittleEndian, &value); err != nil {
			// There is a previous value that we have issues reading.
			// With that, we would risk consistency and instead bail out
			// here.
			return value, err
		}
	}

	return value, nil
}

func (self *FileCounterReaderWriter) Write(counter uint64) error {
	f, err := ioutil.TempFile("", "SurlTemporaryState")

	if err != nil {
		return err
	}

	tempFileName := f.Name()
	err = binary.Write(f, binary.LittleEndian, counter)
	f.Close()

	if err != nil {
		return err
	}

	if err = os.Rename(tempFileName, self.fn); err != nil {
		return err
	}

	return nil
}

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
