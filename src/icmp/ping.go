package icmp

import "fmt"

type PingHeader struct {
	id             uint16
	sequenceNumber uint16
}

type PingMessage struct {
	ICMPHeader
	PingHeader
	data []byte
}

func NewPingMessage(id uint16, sequenceNr uint16) *PingMessage {
	msg := new(PingMessage)
	msg.messageType = T_ECHO_REQUEST
	msg.code = uint8(0)
	msg.checksum = uint16(0)
	msg.id = id
	msg.sequenceNumber = sequenceNr
	return msg
}

func Ping(hostName string, port int) {
	fmt.Printf("Hello, World!\n")
}
