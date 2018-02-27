package dconn

import (
	"time"

	"github.com/utrack/dplay-go/dcommon"
	"github.com/utrack/dplay-go/dmsg"
)

const maxDPVer = uint16(0x4)

func stateConnect(c *conn) stateFn {
	for {
		var msg dmsg.Base
		var ok bool
		select {
		case msg, ok = <-c.inRaw:
			if !ok {
				return nil
			}
		case <-c.ctxClose.Done():
			return nil
		}

		pkt, ok := getPkgConnect(msg)
		if !ok {
			// TODO put msg back to buf
			continue
		}

		verMaj, verMin := pkt.ProtoVersion()
		if verMaj != 0x1 {
			// TODO put msg back to buf
			continue
		}

		minVer := maxDPVer
		if verMin < minVer {
			minVer = verMin
		}

		c.protoVer = minVer
		c.sessID = pkt.SessionID()

		rsp := dmsg.NewCFrameConnected([]byte{}, false) // TODO grab an existing buffer
		rsp.ResponseID(pkt.MessageID())
		rsp.SessionID(c.sessID)
		rsp.ProtoVersion(c.protoVer)

		timeout := 200 * time.Millisecond
		retries := uint8(0)
		shouldSend := true

		// TODO put msg back to buf

		for {
			if shouldSend {
				rsp.MessageID(retries)
				rsp.Timestamp(dcommon.Now())

				c.sendPkt(rsp)
			}

			select {
			case got := <-c.inRaw:
				_, ok := getPkgConnected(got, c.sessID, retries, false)
				// TODO put got back to buf
				if !ok {
					shouldSend = false
					continue
				}

				// TODO switch to state CONNECTED if ok
				return nil

				// TODO timeout should not reset with each incoming pkt.
				// Need to write aux func/type for retries with fixed real time.
			case <-time.After(timeout):
				if retries > 14 {
					c.closeConn() // no ACK rcvd
					return nil
				}
				timeout = timeout * 2
				if timeout > 5*time.Second {
					timeout = 5 * time.Second
				}
				retries++
				shouldSend = true
				continue
			}
		}
	}
}

func getPkgConnect(msg dmsg.Base) (dmsg.CFrameConnect, bool) {
	command := msg.Command()
	if !command.IsCFrame() {
		return nil, false
	}
	if !(command.IsEq(dmsg.CmdCFrame) && command.IsEq(dmsg.CmdCFramePoll)) {
		return nil, false
	}

	cframe := dmsg.CFrame(msg)
	if !cframe.ExtOpCode().IsEq(dmsg.CfOpCodeConnect) {
		return nil, false
	}
	pkt := dmsg.CFrameConnect(msg)
	return pkt, true
}

func getPkgConnected(msg dmsg.Base, sessID uint32, msgID uint8, shouldHaveAckNow bool) (dmsg.CFrameConnected, bool) {
	command := msg.Command()
	if !command.IsCFrame() {
		return nil, false
	}
	if !command.IsEq(dmsg.CmdCFrame) {
		return nil, false
	}
	if shouldHaveAckNow && !command.IsEq(dmsg.CmdCFramePoll) {
		return nil, false
	}
	cframe := dmsg.CFrame(msg)
	if !cframe.ExtOpCode().IsEq(dmsg.CfOpCodeConnected) {
		return nil, false
	}
	ret := dmsg.CFrameConnected(msg)
	if ret.ResponseID() != msgID {
		return nil, false
	}
	if ret.SessionID() != sessID {
		return nil, false
	}
	return ret, true
}
