package kcp

import "net"

const pingLen = 8
const reqKey = "ping"
const ackKey = "pong"
const reqKeyLen = len(reqKey)

func pingCore(input []byte) []byte {
	response := make([]byte, pingLen)
	copy(response, ackKey)
	copy(response[reqKeyLen:], input[reqKeyLen:])
	return response
}

func PluginPing(conn net.PacketConn, buf []byte, n int, addr net.Addr) (holdUp bool) {
	if n == 8 && buf[0] == reqKey[0] && buf[1] == reqKey[1] && buf[2] == reqKey[2] && buf[3] == reqKey[3] {
		// 直接操作 []byte，避免类型转换
		response := pingCore(buf)
		_, _ = conn.WriteTo(response, addr)
		return true
	}
	return false
}
