package icmp

import "fmt"

const (
	T_ECHO_RESPONSE = byte(0)
	T_ECHO_REQUEST  = byte(8)
)

type ICMPHeader struct {
	messageType byte
	code        byte
	checksum    uint16
}

func CalculateICMPChecksum(data []byte) uint16 {
	fmt.Printf("check\n")
	return uint16(0)
}
