package tcp

import (
	"bufio"
	"fmt"
	"net"
)

/*
func ReciveMsg() {
	listener, err :=  net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listening error :", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accepting error :", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("CONNETİNG NEW CLİENT:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("ERROR of reading :", err)
			return
		}

		fmt.Printf("recived packages : %s\n", string(buffer[:n]))
	}
}
*/
func ReciveMsg() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("server not started", err)
		return
	}

	defer listener.Close()

	fmt.Println("tcp is started, waiting connet")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error ", err)
			continue
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new client connected", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)


	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading error ", err)

	}

	fmt.Println("client is disconted", conn.RemoteAddr().String())
}