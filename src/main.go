package main

import (
	//"fmt"
	"icmp"
	"os"
)

func main() {
	msg := icmp.NewPingMessage(1235, 666, []byte{})
	//fmt.Printf("%s\n", msg)
	data := msg.Serialize()
	data = data
	os.Stdout.Write(data)
}
