package main

import (
	"bufio"
	"fmt"
	"net"
)

//记录所有的client连接
var ConnMap map[string]*net.TCPConn

func main() {
	ConnMap = make(map[string]*net.TCPConn)
	// var tcpAddr *net.TCPAddr
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListen.Close()

	for {
		tcpConn, err := tcpListen.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println(" client connected :", tcpConn.RemoteAddr().String())
		//新client加入map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		go connhandle(tcpConn)
	}

}
func connhandle(tcpConn *net.TCPConn) {
	ipStr := tcpConn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected:", ipStr)
		tcpConn.Close()
	}()
	reader := bufio.NewReader(tcpConn)
	for {
		massage, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(tcpConn.RemoteAddr().String() + ":" + string(massage))
		// 这里返回消息改为了广播
		boradcastMessage(tcpConn.RemoteAddr().String() + ":" + string(massage))
	}
}

func boradcastMessage(messages string) {
	b := []byte(messages)
	//遍历所有的client连接并发送信息
	for _, conn := range ConnMap {
		conn.Write(b)
	}
}
