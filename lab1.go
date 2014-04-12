package triblab

import (
	"trib"
)

import "tribsol"

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	return tribsol.NewClient(addr)
}

// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	return tribsol.ServeBack(b)
}
