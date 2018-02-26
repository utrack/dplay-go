package dmsg

// Base is a base type for any rcvd message.
type Base []byte

// Command returns a command byte for this message.
func (b Base) Command() Command {
	return Command(b[0])
}

// Command describes a type of rcvd message.
type Command byte

func (c Command) IsCFrame() bool {
	return (c & CmdCFrame) == CmdCFrame
}

const (
	// CmdCFrame is a PACKET_COMMAND_CFRAME.
	CmdCFrame Command = 0x80
	// CmdCFramePoll is a PACKET_COMMAND_POLL - CFRAME that should be ACKd
	// immediately.
	CmdCFramePoll = 0x08
)

const (
	// CmdData is a PACKET_COMMAND_DATA. It contains user data.
	CmdData Command = 1 << iota
	// CmdDReliable is a DFRAME that should be delivered reliably.
	// See PACKET_COMMAND_RELIABLE.
	CmdDReliable
	// CmdDSequential indicates that this frame is a part of a sequential payload.
	// See PACKET_COMMAND_SEQUENTIAL.
	CmdDSequential
	// CmdDPoll requests the partner to ack this packet immediately.
	// See PACKET_COMMAND_POLL.
	CmdDPoll
	// CmdDNewMsg marks this pkt as a start of a seq message.
	// See PACKET_COMMAND_NEW_MSG.
	CmdDNewMsg
	// CmdDEndMsg marks this pkt as an end of a seq message.
	// See PACKET_COMMAND_END_MSG.
	CmdDEndMsg
	CmdDUser1
	CmdDUser2
)

// CFrame is a CFRAME message.
// CFRAME does not contain any app payload data.
// See MC-DPL8R 2.2.1.
type CFrame Base

func (f CFrame) ExtOpCode() CFrameOpCode {
	return CFrameOpCode(f[1])
}

// CFrameOpCode indicates what type of a request is being made with this CFRAME.
// See MC-DPL8R 2.2.1.[1-5] for details.
type CFrameOpCode byte

const (
	// CfOpCodeConnect is a CONNECT CFRAME message. It is used to request a connection.
	// Response is CONNECTED or CONNECTED_SIGNED.
	CfOpCodeConnect CFrameOpCode = 0x01
	// CfOpCodeConnected is a CONNECTED CFRAME message.
	CfOpCodeConnected = 0x02
	// TODO Signed 0x03

	// CfOpCodeHardDiscon is a HARD_DISCONNECT message.
	CfOpCodeHardDiscon = 0x04
	// CfOpCodeSack is a SACK message.
	CfOpCodeSack = 0x06
)
