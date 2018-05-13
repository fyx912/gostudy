package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var quitSemaphore chan bool

func checkError(err error) {
	if err != nil {
		fmt.Println("error :", err.Error())
		os.Exit(1)
	}
}

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer func() {
		fmt.Println("client close!")
		conn.Close()
	}()
	if err != nil {
		checkError(err)
	}
	go onReviceMessage(conn)
	// 控制台聊天功能输入
	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" || msg == "exit" {
			break
		}

		b := []byte(msg + "\n")
		conn.Write(b)

		<-quitSemaphore
	}

}

func onReviceMessage(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}
