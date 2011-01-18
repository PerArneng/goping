package main

import (
	"fmt"
	"icmp"
)

func main() {
	msg := icmp.NewPingMessage(0, 0)
	fmt.Printf("%s\n", msg)
}
