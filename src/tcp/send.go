package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)


func Sendmsg() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 3*time.Second)
	if err != nil {
		fmt.Println("connection error", err)
		return
	}
	defer conn.Close()
	// msg
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("enter your message:")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
			return
		}
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("sending msg error:", err)
			return
		}
	}

}