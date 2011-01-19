package main

import (
	//"fmt" "os"
	"icmp"
)

func main() {
	icmp.Ping("localhost", 1213, 56, []byte{})
}
