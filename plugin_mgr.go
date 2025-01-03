package kcp

import "net"

type PLUGIN_TYPE int32

const PLUGIN_PING PLUGIN_TYPE = 1

var PluginMgr = map[PLUGIN_TYPE]func(conn net.PacketConn, data []byte, addr net.Addr) bool{
	PLUGIN_PING: PluginPing,
}
