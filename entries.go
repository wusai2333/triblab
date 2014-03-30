package triblab

import (
	"trib"
)

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	panic("todo")
}

// Makes a front end that talks to a list of backends
func NewFront(backs []string) trib.Server {
	panic("todo")
}

// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	panic("todo")
}
