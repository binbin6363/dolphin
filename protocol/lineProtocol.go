package protocol

import (
    "encoding/binary"
    "errors"
    "io"
    "net"
    "github.com/gansidui/gotcp"
)

// string
type LinePacket struct {
    buff []byte
}


func NewLinePacket(buff []byte, hasLengthField bool) *EchoPacket {
    p := &LinePacket{}

    if hasLengthField {
        p.buff = buff

    } else {
        p.buff = make([]byte, 4+len(buff))
        binary.BigEndian.PutUint32(p.buff[0:4], uint32(len(buff)))
        copy(p.buff[4:], buff)
    }

    return p
}

// string
type LineProtocol struct {
    var MaxPacketLen uint32 = (1024 * 100)
}

func (this *LineProtocol) SetMaxPacketLen(maxLen uint32) {
    this.MaxPacketLen = maxLen
}

func (this *LineProtocol) ReadPacket(conn *net.TCPConn) (gotcp.Packet, error) {
    // once read 8 byte, until no data can read
    buff := make([]byte, this.MaxPacketLen)

    return NewLinePacket(buff, true), nil
}
