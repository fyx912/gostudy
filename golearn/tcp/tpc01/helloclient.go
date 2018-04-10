package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":10000")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte("hello"))

	var buf = make([]byte, 50)
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn closed")
			return
		}

		fmt.Println("recv msg byte:%V", buf[0:n])
		fmt.Println("recv msg data:%V", string(buf[0:n]))
	}
}
