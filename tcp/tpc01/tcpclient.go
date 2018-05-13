package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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
	defer conn.Close()
	if err != nil {
		checkError(err)
	}
	go onReviceMessage(conn)
	b := []byte("time\n")
	conn.Write(b)
	<-quitSemaphore
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
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
