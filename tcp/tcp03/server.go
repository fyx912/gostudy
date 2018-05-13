//服务端解包过程
package main

import (
	"./protocol"
	"fmt"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", "127.0.0.1:9999")
	CheckErr(err)
	defer netListen.Close()
	Log("wait for clients....")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " TCP connected success")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connected error:", err)
			return
		}
		//服务器进行解包处理
		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readerChnnel chan []byte) {
	for {
		select {
		case data := <-readerChnnel:
			Log(string(data))
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s", err.Error())
		os.Exit(1)
	}
}
