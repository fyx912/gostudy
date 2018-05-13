package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, _ := net.Listen("tcp", ":10000")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			checkError(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		fmt.Println("byte data:", reader)
		if err != nil {
			fmt.Println(" server connection close ...")
			return
		}
		fmt.Println(string(message))
		msg := "wolrd"
		b := []byte(msg)
		conn.Write(b)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
