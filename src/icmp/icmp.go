package icmp

import "fmt"

const (
	T_ECHO_RESPONSE = uint8(0)
	T_ECHO_REQUEST  = uint8(8)
)

type ICMPHeader struct {
	messageType uint8
	code        uint8
	checksum    uint16
}

func CalculateICMPChecksum(data []uint8) uint16 {
	fmt.Printf("check\n")
	return uint16(0)
}
