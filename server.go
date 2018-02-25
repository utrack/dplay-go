// Package dplay provides DirectPlay server implementation.
package dplay

import (
	"bytes"
	"net"
	"sync"

	"github.com/Sirupsen/logrus"
)

// Server is a DirectPlay server.
// It controls the primary DP UDP socket and
// DirectPlay sessions.
type Server struct {
	// IsPassworded bool

	opts ServerOptions

	// listening conn (UDP)
	listenConn *net.UDPConn
	// r/w buffer pool
	bufpool sync.Pool

	// output queue
	outQueue chan message
}

const (
	// default Server's message buffer size
	defaultBufLen = 0x100
)

// NewServer returns new DirectPlay server.
func NewServer(addr string, opts ServerOptions) (ret *Server, err error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return
	}

	c, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return
	}

	ret = &Server{listenConn: c, opts: opts}
	ret.outQueue = make(chan message, 50)
	ret.bufpool.New = func() interface{} {
		return &data{
			Buffer: bytes.NewBuffer(make([]byte, 0, defaultBufLen)),
		}
	}
	return
}

// Listen starts the socket processes synchronously.
// Also initiates the output messages' pump.
func (s *Server) Listen() {
	logrus.Info("Running")
	go s.outPump()
	for {
		wr := s.getBuffer()
		buf := wr.Bytes()
		buf = buf[:defaultBufLen]
		n, ep, err := s.listenConn.ReadFromUDP(buf)

		if err != nil {
			logrus.Warn(err)
		}
		if n < 2 {
			logrus.Debug("Small packet")
			continue
		}

		// Start parsing
		s.preprocessPacket(buf[:n], ep)
		s.bufpool.Put(wr)
	}
}

// outPump pumps out outgoing messages.
func (s *Server) outPump() {
	logrus.Debug("Dplay output pump started")
	for {
		out := <-s.outQueue
		msg := out.wr.Bytes()
		logrus.Debugf("put % #x", msg)
		s.listenConn.WriteToUDP(msg, out.addr)
		s.bufpool.Put(out.wr)
	}
}
