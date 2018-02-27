package dconn

import (
	"context"

	"github.com/utrack/dplay-go/dmsg"
)

type conn struct {
	closeConn func()
	ctxClose  context.Context // context is closed when the conn is closed
	inRaw     <-chan dmsg.Base
	sessID    uint32
	protoVer  uint16 // has minor mart of a DP version

	sendPkt func([]byte)
}
