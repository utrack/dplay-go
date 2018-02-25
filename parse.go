package dplay

import (
	"bytes"
	"net"

	"github.com/Sirupsen/logrus"
)

// preprocessPacket determines the DP packet type and sends it further.
func (s *Server) preprocessPacket(msg []byte, ep *net.UDPAddr) {
	p := Packet{msg, 0}

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
		queryType := p.GetUint8() // 1 if there's GUID in the request
		if queryType == 0x01 {
			guid := p.GetBytes(16)
			logrus.Debug("got enum request guid %v", guid)
			if bytes.Compare(guid, s.opts.ApplicationGUID) != 0 {
				return
				// Do not respond to unknown GUID enum
				// requests as per dplay spec
			}
		}
		s.sendEnumResponse(ep, enumPayload)

	default:
		logrus.WithField("msg", msg).Debug("Unknown cmd type!")
	}
}
