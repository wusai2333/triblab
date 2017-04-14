package triblab

import (
	"trib"
	"fmt"
	"net"
	"net/rpc"
	"net/http"
)

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	return &client{Addr:addr}
}

// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	server := NewServer()
	e := server.RegisterName("Storage", b.Store)
	rpc.HandleHTTP()

	if e != nil {
		fmt.Println(e)
		if b.Ready != nil {
			b.Ready <- false
		}
		return e
	}

	l, e := net.Listen("tcp", b.Addr)

	if e != nil {
		fmt.Println(e)
		if b.Ready != nil {
			b.Ready <- false
		}
		return e
	}

	if b.Ready != nil {
		b.Ready <- true
	}

	return http.Serve(l, server)
}
