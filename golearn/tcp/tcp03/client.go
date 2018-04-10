//客户端发送封包
package main

import (
	"./protocol"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckErr(err)

	defer func() {
		fmt.Println(" client close")
		conn.Close()
	}()

	fmt.Println("connect success")

	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}

func sender(conn *net.TCPConn) {
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		//客户端封包发送
		conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println("send over")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s", err.Error())
		os.Exit(1)
	}
}
