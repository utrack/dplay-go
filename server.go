// Package dplay provides DirectPlay server implementation.
package dplay

import (
	"github.com/Sirupsen/logrus"
	"net"
	"sync"
)

// Server is a DirectPlay server.
// It controls the primary DP UDP socket and
// DirectPlay sessions.
type Server struct {
	// listening conn (UDP)
	listenConn *net.UDPConn
	// r/w buffer pool
	bufpool sync.Pool
}

const (
	// default Server's buffer size
	defaultBufLen = 0x100
)

// NewServer returns new DirectPlay server.
func NewServer(addr string) (ret *Server,err error) {
	udpAddr,err := net.ResolveUDPAddr("udp",addr)
	if err != nil {
		return
	}

	c,err := net.ListenUDP("udp",udpAddr)
	if err != nil {
		return
	}

	ret = &Server{listenConn:c}
	ret.bufpool.New = func() interface{} {
		return make([]byte,0,defaultBufLen)
	}
	return
}

// Listen starts the socket synchronously.
func (s *Server) Listen() {
	for {
		buf := s.bufpool.Get().([]byte)
		// TODO start parsing
		_,_,err := s.listenConn.ReadFromUDP(buf)

		if err != nil {
			logrus.Warn(err)
		}
	}
}
