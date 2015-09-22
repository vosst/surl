package surl

// Ticketer models generation of ids, or tickets,
// for uniquely identifying entities stored in a datastore.
type Ticketer interface {
	// Next returns the next available ticket.
	// Implementations are expected to provide the shortest
	// possible next ticket
	Next() string
}
