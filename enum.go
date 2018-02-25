package dplay

import (
	"net"
)

func (s *Server) sendEnumResponse(ep *net.UDPAddr, enumPayload Uint16) {

	p := s.getBuffer()
	p.AddByte(0x00)
	p.AddByte(0x03) // enum response

	// ResponseSize-related
	// This has aux data specific for our application.
	auxData := s.opts.EnumResponseData()

	p.AddUint16(enumPayload)                                       // EnumPayload
	p.AddUint32(Uint32(0x58 + uint(len([]byte(s.opts.Name))*2+2))) // ReplyOffset
	p.AddUint32(Uint32(len(auxData)))                              // ReplySize + ResponseSize
	p.AddUint32(Uint32(0x50))                                      // ApplicationDescSize

	// ApplicationDescFlags
	{
		n1 := uint(0x00)
		// TODO password
		// if s.opts.IsPassworded {
		// 	n1 = uint(0x80)
		// }
		// TODO 0x00 instead of 0x40 if DPNSVR is used; however it's dead
		p.AddUint32(Uint32(n1 | uint(0x40))) // ApplicationDescFlags
	}

	p.AddUint32(Uint32(s.opts.MaxPlayers) + 1)
	// TODO send current player count
	p.AddUint32(Uint32(0 + 1))

	p.AddUint32(Uint32(0x58))                           // SessionNameOffset
	p.AddUint32(Uint32(len([]byte(s.opts.Name))*2 + 2)) // SessionNameSize
	// TODO password
	p.AddUint32(Uint32(0)) // PasswordOffset
	p.AddUint32(Uint32(0)) // PasswordSize
	p.AddUint32(Uint32(0)) // ReservedDataOffset
	p.AddUint32(Uint32(0)) // ReservedDataSize
	p.AddUint32(Uint32(0)) // ApplicationReservedDataOffset
	p.AddUint32(Uint32(0)) // ApplicationReservedDataSize

	p.AddArray(s.opts.ApplicationInstanceGUID)
	p.AddArray(s.opts.ApplicationGUID)
	p.AddStringUnicodeTerm(s.opts.Name)
	p.AddArray(auxData)

	s.outQueue <- newMessage(p, ep)
}
