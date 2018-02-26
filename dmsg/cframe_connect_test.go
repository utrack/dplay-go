package dmsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBuf() []byte {
	buf := make([]byte, 0, 128)
	return buf
}

func TestCFRAME_Connect_Getters(t *testing.T) {
	so := assert.New(t)

	rawPkt := []byte{
		0x80 | 0x08, // bCommand
		0x01,        // bExtOpCode
		0x05,        // bMsgID
		0x06,        // bRspID
		0x05, 0x00,  // proto minor
		0x01, 0x00, // proto major

		0xC6, 0xAE, // dwSessID
		0xC9, 0x79,

		0x37, 0x13, // tTimestamp 13371337
		0x37, 0x13,
	}

	base := Base(rawPkt)
	cf := CFrame(base)
	conn := CFrameConnect(cf)

	so.Equal(Command(0x88), base.Command())
	so.Equal(CfOpCodeConnect, cf.ExtOpCode())
	so.Equal(uint8(0x05), conn.MessageID())

	vMajor, vMinor := conn.ProtoVersion()
	so.Equal(uint16(1), vMajor)
	so.Equal(uint16(5), vMinor)
	so.Equal(uint32(0x79C9AEC6), conn.SessionID())
	so.Equal(uint32(0x13371337), conn.Timestamp())
}

func TestCFRAME_Connect_Setter(t *testing.T) {
	so := assert.New(t)

	conn := NewCFrameConnect(newBuf())
	so.Equal(len(conn), 16)
	conn.Timestamp(0x13371337)
	conn.SessionID(0x13C9C7AE)
	conn.ProtoVersion(0x05)
	conn.MessageID(0xF2)

	cfc := CFrameConnect(conn)

	so.Equal(uint8(0xF2), cfc.MessageID())

	vMj, vMn := cfc.ProtoVersion()
	so.Equal(uint16(1), vMj)
	so.Equal(uint16(5), vMn)

	so.Equal(uint32(0x13C9C7AE), cfc.SessionID())
	so.Equal(uint32(0x13371337), cfc.Timestamp())

	cf := CFrame(cfc)
	base := Base(cfc)
	so.True(base.Command().IsCFrame())
	so.Equal(CmdCFramePoll, base.Command()&CmdCFramePoll)
	so.Equal(CfOpCodeConnect, cf.ExtOpCode())
}

func TestCFRAME_Connected_Setter(t *testing.T) {
	so := assert.New(t)

	// See 2.2.1.2 of MC-DPL9R on rules about COMMAND_POLL for client and server.
	conn := NewCFrameConnected(newBuf(), false)
	base := Base(conn)
	so.True(base.Command().IsCFrame())
	so.Equal(CmdCFramePoll, base.Command()&CmdCFramePoll)

	conn = NewCFrameConnected(newBuf(), true)
	base = Base(conn)
	so.True(base.Command().IsCFrame())
	so.Equal(Command(0x0), base.Command()&CmdCFramePoll)
}
