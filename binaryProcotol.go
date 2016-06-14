package dolphin

import (
	"encoding/binary"
	"errors"
	"io"
	"net"

	"github.com/gansidui/gotcp"
)

// len + binarydata
type BinaryPacket struct {
	buff []byte
}


func NewBinaryPacket(buff []byte, hasLengthField bool) *EchoPacket {
    p := &BinaryPacket{}

    if hasLengthField {
        p.buff = buff

    } else {
        p.buff = make([]byte, 4+len(buff))
        binary.BigEndian.PutUint32(p.buff[0:4], uint32(len(buff)))
        copy(p.buff[4:], buff)
    }

    return p
}

// len + binarydata
type BinaryProtocol struct {
    var MaxPacketLen uint32 = (1024 * 100)
}

func (this *BinaryProtocol) SetMaxPacketLen(maxLen uint32) {
    this.MaxPacketLen = maxLen
}

func (this *BinaryProtocol) ReadPacket(conn *net.TCPConn) (dolphin.Packet, error) {
	var (
		lengthBytes []byte = make([]byte, 4)
		length      uint32
	)

	// read length, block
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		return nil, err
	}
	if length = binary.BigEndian.Uint32(lengthBytes); length > this.MaxPacketLen {
		return nil, errors.New("the size of packet is larger than the limit")
	}

	buff := make([]byte, 4+length)
	copy(buff[0:4], lengthBytes)

	// read body ( buff = lengthBytes + body ), block
	if _, err := io.ReadFull(conn, buff[4:]); err != nil {
		return nil, err
	}

    return NewBinaryPacket(buff, true), nil
}
