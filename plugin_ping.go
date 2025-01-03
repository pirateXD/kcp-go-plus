package kcp

import "net"

const pingLen = 8
const reqKey = "ping"
const ackKey = "pong"

func PluginPing(conn net.PacketConn, buf []byte, addr net.Addr) (holdUp bool) {
	if len(buf) != pingLen {
		return false
	}

	if buf[0] == reqKey[0] && buf[1] == reqKey[1] && buf[2] == reqKey[2] && buf[3] == reqKey[3] {
		buf[1] = ackKey[1]
		_, _ = conn.WriteTo(buf, addr)
		return true
	}
	return false
}
