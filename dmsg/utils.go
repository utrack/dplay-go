package dmsg

func getUint16(b1, b2 byte) uint16 {
	return (uint16(b2) << 8) | uint16(b1)
}

func getUint32(b []byte) uint32 {
	ret := (uint32(b[3]) << 24) |
		(uint32(b[2]) << 16) |
		(uint32(b[1]) << 8) |
		uint32(b[0])
	return ret
}

func setUint16(b []byte, v uint16) {
	b[0] = uint8(v)
	b[1] = uint8(v >> 8)
}

func setUint32(b []byte, v uint32) {
	setUint16(b[0:2], uint16(v))
	setUint16(b[2:4], uint16(v>>16))
}
