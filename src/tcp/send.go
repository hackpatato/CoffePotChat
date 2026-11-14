package tcp

import (
	"fmt"
	"net"
	"time"
)


func Sendmsg() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 3*time.Second)
	if err != nil {
		fmt.Println("connection error", err)
		return
	}
	defer conn.Close()
	//msg
	var msg string
	
	fmt.Println("enter your message:")
	_, err = fmt.Scanf("%s", &msg)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("sending msg error :", err)
		return
	}
    

}