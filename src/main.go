package main

import (
	//"fmt"
	"icmp"
	//"net"
)

func main() {
	icmp.Ping("127.0.0.1", 1213, 1, []byte{})
}
