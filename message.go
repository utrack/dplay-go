package dplay

import (
	"net"
)

// message is a request for UDP socket to
// send out the message to client.
type message struct {
	wr   *data
	addr *net.UDPAddr
}

func newMessage(b *data, addr *net.UDPAddr) message {
	return message{
		wr:   b,
		addr: addr,
	}
}
