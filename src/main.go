package main

import (
	"CoffePotChat/tcp"
	"time"
)

func main() {
	tcp.ReciveMsg()
	time.Sleep(200 * time.Millisecond)
	tcp.Sendmsg()
}