package dplay

import (
	"net"
)


// outMsg is a request for UDP socket to
// send out the message to client.
type outMsg struct {
	msg []byte
	addr *net.UDPAddr
}

func (o *outMsg) AddUint16(d Uint16) {
	o.msg = append(o.msg,byte(d & 0xFF),byte( (d>>8) & 0xFF))
}


func (o *outMsg) AddUint32(d Uint32) {
	o.msg = append(o.msg,
		byte(d & 0xFF),
		byte( (d>>8) & 0xFF),
		byte( (d>>16) & 0xFF),
		byte( (d>>24) & 0xFF),
	)
}

func (o *outMsg) AddArray(d []byte) {
	o.msg = append(o.msg,d...)
}
