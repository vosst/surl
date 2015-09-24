package surl

// CounterReaderWriter abstracts reading and writing of a single
// uint64 counter value.
type CounterReaderWriter interface {
	// Read tries to read the counter value from the underlying storage solution.
	// If there exists to previous value, defaultValue is returned.
	Read(defaultValue uint64) (uint64, error)
	// Write stores the value of counter to the underlying storage solution.
	Write(counter uint64) error
}
