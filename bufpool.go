package dplay

// Returns new []byte slice backed by pooled []byte
// array.
func (s *Server) getBuffer() []byte {
	return s.bufpool.Get().([]byte)[0:0]
}
