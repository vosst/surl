package surl

import (
	"encoding/binary"
	"io/ioutil"
	"os"
)

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
