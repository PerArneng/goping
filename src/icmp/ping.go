package icmp

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type PingHeader struct {
	id             uint16
	sequenceNumber uint16
}

type PingMessage struct {
	ICMPHeader
	PingHeader
	data []byte
}

func NewPingMessage(id uint16, sequenceNr uint16, data []byte) *PingMessage {
	msg := new(PingMessage)
	msg.messageType = T_ECHO_REQUEST
	msg.code = byte(0)
	msg.checksum = uint16(0)
	msg.id = id
	msg.sequenceNumber = sequenceNr
	msg.data = data
	return msg
}

func (msg *PingMessage) Serialize() []byte {
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, msg.messageType)
	binary.Write(buff, binary.BigEndian, msg.code)
	binary.Write(buff, binary.BigEndian, msg.checksum)
	binary.Write(buff, binary.BigEndian, msg.id)
	binary.Write(buff, binary.BigEndian, msg.sequenceNumber)
	buff.Write(msg.data)
	return buff.Bytes()
}

func Ping(hostName string, port int) {
	fmt.Printf("Hello, World!\n")
}
