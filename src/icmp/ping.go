package icmp

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"os"
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

type PingClient struct {
	icmpClient *ICMPClient
}

func NewPingClient(localAddr, remoteAddr *net.IPAddr) (*PingClient, os.Error) {
	client := new(PingClient)
	icmpClient, e := NewICMPClient(localAddr, remoteAddr)
	if e != nil {
		fmt.Printf("%s\n", e)
		return nil, e
	}
	client.icmpClient = icmpClient
	return client, nil
}

func (client *PingClient) Send(message *PingMessage) {
	pingData := message.Serialize()
	icmpMessage := NewICMPMessage(T_ECHO_REQUEST, byte(0), pingData)
	e := client.icmpClient.Send(icmpMessage)
	if e != nil {
		fmt.Printf("%s\n", e)
	}
}

func (client *PingClient) Close() {
	client.icmpClient.Close()
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

func Ping(hostName string, id uint16, sequence uint16, payload []byte) os.Error {

	pingMsg := NewPingMessage(id, sequence, payload)
	pingData := pingMsg.Serialize()

	icmpMsg := NewICMPMessage(T_ECHO_REQUEST, byte(0), pingData)

	localAddr, e := net.ResolveIPAddr("0.0.0.0")
	if e != nil {
		fmt.Printf("%s\n", e)
		return e
	}

	remoteAddr, e := net.ResolveIPAddr(hostName)
	if e != nil {
		fmt.Printf("%s\n", e)
		return e
	}

	cl, e := NewICMPClient(localAddr, remoteAddr)
	if e != nil {
		fmt.Printf("%s\n", e)
		return e
	}

	cl.Send(icmpMsg)
	cl.Close()

	return nil
}
