package dmsg

// CFrameSack is a SACK CFRAME.
type CFrameSack CFrame

func (f CFrameSack) Flags() SackFlags {
	return SackFlags(f[2])
}

// Retry indicates if last rcvd packet was a retry.
// Must be ignored if SackFlagResponse is not set.
func (f CFrameSack) Retry() bool {
	return f[3] > 0
}

// NumSeq represents the seq number of next DFRAME to send.
func (f CFrameSack) NumSeq() uint8 {
	return uint8(f[4])
}

// NumRcv is an expected sequence number of the next packet.
// if SACK_MASK[1|2] is set then field is supplemented with additional
// dwSACKMask[1|2] bitmask field that selectively acks frames with seq > NumRcv.
func (f CFrameSack) NumRcv() uint8 {
	return uint8(f[5])
}

// Timestamp returns tick count of a sender in ms.
func (f CFrameSack) Timestamp() uint32 {
	return getUint32(f[8:12])
}

// GetAckOrSendMasks returns optional SACK or SEND masks.
func (f CFrameSack) GetAckOrSendMasks(ord int) uint32 {
	return getUint32(f[12+ord : 16+ord])
}

// SackFlags is a SACK flag in SACK CFRAME packet.
type SackFlags byte

const (
	SackFlagResponse SackFlags = 1 << iota
	SackFlagSackMask1
	SackFlagSackMask2
	SackFlagSendMask1
	SackFlagSendMask2
)
