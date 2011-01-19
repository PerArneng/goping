package icmp

import (
	//"fmt"
	"bytes"
	"encoding/binary"
	//"os"
	"net"
)

type PingHeader struct {
	id             uint16
	sequenceNumber uint16
}

type PingMessage struct {
	PingHeader
	payload []byte
}

func NewPingMessage(id uint16, sequenceNr uint16, payload []byte) *PingMessage {
	msg := new(PingMessage)
	msg.id = id
	msg.sequenceNumber = sequenceNr
	msg.payload = payload
	return msg
}

func (msg *PingMessage) Serialize() []byte {
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, msg.id)
	binary.Write(buff, binary.BigEndian, msg.sequenceNumber)
	buff.Write(msg.payload)
	return buff.Bytes()
}

func Ping(hostName string, id uint16, sequence uint16, payload []byte) {

	pingMsg := NewPingMessage(id, sequence, payload)
	pingData := pingMsg.Serialize()

	icmpMsg := NewICMPMessage(T_ECHO_REQUEST, byte(0), pingData)
	icmpData := icmpMsg.Serialize()

	udpAddresss, _ := net.ResolveUDPAddr(hostName)

	udp, _ := net.DialUDP("udp", nil, udpAddresss)
	udp.Write(icmpData)
	udp.Close()
	//os.Stdout.Write(icmpData)
	//fmt.Printf("Hello, World!\n")
}
