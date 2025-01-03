package kcp

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"testing"
)

func int32ToBytesBigEndian(n int32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(n))
	return bytes

}

func TestPingCoreFailed(t *testing.T) {
	{
		addr := "127.0.0.1:123"
		_, _ = ListenWithOptions(addr, nil, 0, 0)
		udpPint(addr, t)
	}
}

func PluginPing(conn net.PacketConn, buf []byte, addr net.Addr) (holdUp bool) {
	const pingLen = 8
	const reqKey = "ping"
	const ackKey = "pong"
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
func TestPingCoreSuccess(t *testing.T) {
	{
		addr := "127.0.0.1:456"
		_, _ = ListenWithOptions(addr, nil, 0, 0, PluginPing)
		udpPint(addr, t)
	}
}

func udpPint(addr string, t *testing.T) {
	message := "ping"

	// 解析UDP地址
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// 建立UDP连接
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送消息
	_, err = conn.Write(append([]byte(message), int32ToBytesBigEndian(12313213)...))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent:", message)
	readMsg := make([]byte, 100)
	n, _, err := conn.ReadFromUDP(readMsg)
	t.Logf("[addr:%v]Message received: %s [number:%v][n:%v]", addr, string(readMsg[:4]), int32(binary.BigEndian.Uint32(readMsg[4:])), n)
}
