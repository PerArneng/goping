package icmp

import (
	"fmt"
	"encoding/binary"
	"bytes"
	"net"
	"os"
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

type ICMPClient struct {
	localAddr  *net.IPAddr
	remoteAddr *net.IPAddr
	conn       *net.IPConn
}

func NewICMPClient(localAddr, remoteAddr *net.IPAddr) (*ICMPClient, os.Error) {
	client := new(ICMPClient)
	client.localAddr = localAddr
	client.remoteAddr = remoteAddr
	conn, e := net.ListenIP("ip4:icmp", localAddr)
	if e != nil {
		fmt.Printf("%s\n", e)
		return nil, e
	}
	client.conn = conn
	return client, nil
}

func (client *ICMPClient) Send(message *ICMPMessage) {
	cnt, e := client.conn.WriteToIP(message.Serialize(), client.remoteAddr)
	if e != nil {
		fmt.Printf("%s\n", e)
	}
	fmt.Printf("%s\n", cnt)
}

func (client *ICMPClient) Close() {
	client.conn.Close()
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
