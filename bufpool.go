package dplay

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

// Returns new []byte slice backed by pooled []byte
// array.
func (s *Server) getBuffer() *data {
	d := s.bufpool.Get().(*data)
	d.Reset()
	return d
}

func NewDataPackage() *data {
	return &data{Buffer: bytes.NewBuffer(make([]byte, 0, defaultBufLen))}
}

type data struct {
	*bytes.Buffer
}

func (o *data) AddUint16(d Uint16) {
	o.WriteByte(byte(d & 0xFF))
	o.WriteByte(byte((d >> 8) & 0xFF))
}

func (o *data) AddUint32(d Uint32) {
	o.WriteByte(byte(d & 0xFF))
	o.WriteByte(byte((d >> 8) & 0xFF))
	o.WriteByte(byte((d >> 16) & 0xFF))
	o.WriteByte(byte((d >> 24) & 0xFF))
}
func (o *data) AddByte(d byte) {
	o.WriteByte(d)
}

func (o *data) AddArray(d []byte) {
	binary.Write(o, binary.BigEndian, d)
}

func (o *data) AddString(s string) {
	o.WriteString(s)
}

func (o *data) AddStringUnicodeTerm(s string) {
	binary.Write(o, binary.LittleEndian, utf16.Encode([]rune(s)))
	o.WriteByte('\x00')
	o.WriteByte('\x00')
}
