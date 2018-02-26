package dmsg

// CFrameConnect is a CONNECT CFRAME message.
type CFrameConnect CFrame

// CFrameConnected is a CONNECTED CFRAME message, response to CFrameConnect.
type CFrameConnected = CFrameConnect

// CFrameHardDisconnect is a HARD_DISCONNECT CFRAME message.
// It is the same as CONNECT unless there's signing present
// (TODO signing), so atm it's aliased.
type CFrameHardDisconnect = CFrameConnect

// TODO CONNECTED_SIGNED

// MessageID is an identifier of a message.
// It SHOULD start at 0 and SHOULD increment each time one pkt is retried.
func (f CFrameConnect) MessageID() uint8 {
	return f[2]
}

// ProtoVersion returns current protocol version.
// First digit MUST be 0x0001.
// Second digit controls features:
// >= 0x0005 adds coalescence support
// >= 0x0006 adds signing.
func (f CFrameConnect) ProtoVersion() (uint16, uint16) {
	return getUint16(f[4], f[5]), getUint16(f[6], f[7])
}

// SessionID is a random number, unique per session.
func (f CFrameConnect) SessionID() uint32 {
	return getUint32(f[8:12])
}

// Timestamp returns tick count of a sender in ms.
func (f CFrameConnect) Timestamp() uint32 {
	return getUint32(f[12:16])
}
