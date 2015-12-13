package dplay

import (
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
)

func (s *Server) sendEnumResponse(ep *net.UDPAddr, enumPayload Uint16) {
	logrus.WithField("ep", ep).Debug("Sending enum response...")

	// adata is application-specific enum response's data.
	adata := s.getBuffer()
	defer s.bufpool.Put(adata)
	// TODO: this works for FL, need to adapt it for everything else
	adata = append(adata,[]byte(fmt.Sprintf(`1:1:%v:-1910309061:%v:%v`,s.opts.Version,s.opts.ServerID,s.opts.Description))...)

	ret := s.getBuffer()
	defer s.bufpool.Put(ret)
	ret = append(ret,0x00,0x03) // Enum response
	p := outMsg{ret,ep}

	p.AddUint16(enumPayload) // EnumPayload
	p.AddUint32(Uint32(0x58 + uint(len(s.opts.Name)) * 2)) // ReplyOffset
	p.AddUint32(Uint32(len(adata))) // ReplySize + ResponseSize
	p.AddUint32(Uint32(0x50)) // ApplicationDescSize

	// ApplicationDescFlags
	{
		n1 := uint(0x00)
		if s.opts.IsPassworded {
			n1 = uint(0x80)
		}
		// TODO 0x00 instead of 0x40 if DPNSVR is used; however it's dead
		p.AddUint32(Uint32(n1 | uint(0x40))) // ApplicationDescFlags
	}

	p.AddUint32(Uint32(s.opts.MaxPlayers)+1)
	// TODO send current player count
	p.AddUint32(Uint32(0))

	p.AddUint32(Uint32( 0x58)); // SessionNameOffset
	p.AddUint32(Uint32( len(s.opts.Name) * 2)) // SessionNameSize
	p.AddUint32(Uint32( 0)) // PasswordOffset
	p.AddUint32(Uint32( 0)) // PasswordSize
	p.AddUint32(Uint32( 0)) // ReservedDataOffset
	p.AddUint32(Uint32( 0)) // ReservedDataSize
	p.AddUint32(Uint32( 0)) // ApplicationReservedDataOffset
	p.AddUint32(Uint32( 0)) // ApplicationReservedDataSize

	p.AddArray(s.opts.ApplicationInstanceGUID)
	p.AddArray(s.opts.ApplicationGUID)
	// TODO: unicode?
	p.AddArray([]byte(s.opts.Name))
	p.AddArray(adata)

	s.outQueue <- p
}
