package icmp

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

const (
	T_ECHO_RESPONSE = byte(0)
	T_ECHO_REQUEST  = byte(8)
)

type ICMPHeader struct {
	messageType byte
	code        byte
	checksum    uint16
}

type ICMPMessage struct {
	ICMPHeader
	payload []byte
}

func NewICMPMessage(messageType byte, code byte, payload []byte) *ICMPMessage {
	msg := new(ICMPMessage)
	msg.messageType = messageType
	msg.code = byte(0)
	msg.checksum = uint16(0)
	msg.payload = payload
	return msg
}


func (msg *ICMPMessage) Serialize() []byte {
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, msg.messageType)
	binary.Write(buff, binary.BigEndian, msg.code)
	binary.Write(buff, binary.BigEndian, msg.checksum)
	buff.Write(msg.payload)
	return buff.Bytes()
}

func CalculateICMPChecksum(data []byte) uint16 {
	fmt.Printf("check\n")
	return uint16(0)
}
