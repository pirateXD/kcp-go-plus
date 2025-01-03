package kcp

import "net"

type PluginFunc func(conn net.PacketConn, data []byte, addr net.Addr) bool
