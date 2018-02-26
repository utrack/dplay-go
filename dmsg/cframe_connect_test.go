package dmsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCFRAME_Connect_Getters(t *testing.T) {
	so := assert.New(t)

	rawPkt := []byte{
		0x80 | 0x08, // bCommand
		0x01,        // bExtOpCode
		0x05,        // bMsgID
		0x06,        // bRspID
		0x01, 0x00,  // proto major
		0x05, 0x00, // proto minor

		0x10, 0x05, // dwSessID
		0xFF, 0x20,

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
	so.Equal(uint32(0x20ff0510), conn.SessionID())
	so.Equal(uint32(0x13371337), conn.Timestamp())
}
