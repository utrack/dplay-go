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
	return getUint16(f[6], f[7]), getUint16(f[4], f[5])
}

// SessionID is a random number, unique per session.
func (f CFrameConnect) SessionID() uint32 {
	return getUint32(f[8:12])
}

// Timestamp returns tick count of a sender in ms.
func (f CFrameConnect) Timestamp() uint32 {
	return getUint32(f[12:16])
}

// CFrameConnectWriter creates new CONNECT pkt.
type CFrameConnectWriter CFrameConnect

// CFrameConnectedWriter creates new CONNECTED pkt.
type CFrameConnectedWriter = CFrameConnectWriter

// CFrameHardDisconnectWriter creates new HARD_DISCONNECT pkt.
type CFrameHardDisconnectWriter = CFrameConnectWriter

// NewCFrameConnect returns CFrameConnect constructor.
func NewCFrameConnect(buf []byte) CFrameConnectWriter {
	ret := CFrameConnectWriter(buf[:16])
	ret[0] = byte(CmdCFrame | CmdCFramePoll)
	ret[1] = byte(CfOpCodeConnect)
	return ret
}

// NewCFrameConnected returns new CFrameConnected constructor.
// isAckFromClient is true if this package is a response from client to server's
// CONNECTED pkt.
func NewCFrameConnected(buf []byte, isAckFromClient bool) CFrameConnectedWriter {
	ret := CFrameConnectWriter(buf[:16])

	if isAckFromClient {
		ret[0] = byte(CmdCFrame)
	} else {
		ret[0] = byte(CmdCFrame | CmdCFramePoll)
	}
	ret[1] = byte(CfOpCodeConnected)
	return ret
}

// MessageID sets message ID for this pkt.
func (w CFrameConnectWriter) MessageID(id uint8) {
	w[2] = id
}

// ProtoVersion sets lower part of a DPlay version.
// Upper part is always 0x01.
func (w CFrameConnectWriter) ProtoVersion(lower uint16) {
	setUint16(w[4:6], lower)
	setUint16(w[6:8], 0x01)
}

// SessionID sets Session ID for this pkt.
func (w CFrameConnectWriter) SessionID(id uint32) {
	setUint32(w[8:12], id)
}

// Timestamp sets Timestamp for this pkt.
func (w CFrameConnectWriter) Timestamp(ts uint32) {
	setUint32(w[12:16], ts)
}
