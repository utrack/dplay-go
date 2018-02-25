package dplay

type Packet struct {
	buf []byte
	pos uint
}

type Uint8 byte
type Uint16 uint
type Uint32 uint

func (p *Packet) GetUint8() (ret Uint8) {
	p.pos++
	return Uint8(p.buf[p.pos-1])
}

func (p *Packet) GetUint16() (ret Uint16) {
	ret = Uint16((uint(p.buf[p.pos+1]) << 8) | uint(p.buf[p.pos]))
	p.pos += 2
	return
}

func (p *Packet) GetBytes(n uint) []byte {
	ret := p.buf[p.pos : p.pos+n]
	p.pos += n
	return ret
}
