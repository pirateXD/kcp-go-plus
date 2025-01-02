package kcp

import (
	"encoding/binary"
	"testing"
)

func int32ToBytesBigEndian(n int32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(n))
	return bytes
}
func TestPingCore(t *testing.T) {
	pongBytes := pingCore(append([]byte("ping"), int32ToBytesBigEndian(123333)...))
	t.Logf("Message received: %s [number:%v][n:%v]", string(pongBytes[:4]), int32(binary.BigEndian.Uint32(pongBytes[4:])), len(pongBytes))
}
