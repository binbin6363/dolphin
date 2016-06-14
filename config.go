package dolphin

import (
	"strings"
)

// global config for framework
type Config struct {
	PacketSendChanLimit uint32 // the limit of packet send channel, default set 200000
	PacketRecvChanLimit uint32 // the limit of packet receive channel, default set 200000
	PacketMaxRecvLen    uint32 // the limit of packet len receive, default set (1 * 1024 * 1024) byte
	PacketMaxSendLen    uint32 // the limit of packet len send, default set (1 * 1024 * 1024) byte
}
