package dplay

import (
	"net"

	"github.com/Sirupsen/logrus"
)

// preprocessPacket determines the DP packet type and sends it further.
func (s *Server) preprocessPacket(msg []byte, ep *net.UDPAddr) {
	defer s.bufpool.Put(msg)
	p := Packet{msg,0}

	cmd := p.GetUint8()

	switch cmd {
		// enum server status
		// MS-DPDX 2.2.5
	case 0x00:
		opcode := p.GetUint8()
		if opcode != 0x02 {
			return
		}
		enumPayload := p.GetUint16()
		s.sendEnumResponse(ep,enumPayload)

	default:
		logrus.WithField("msg", msg).Debug("Unknown cmd type!")
	}
}
