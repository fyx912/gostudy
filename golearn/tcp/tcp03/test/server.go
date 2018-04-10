//粘包问题演示服务端
package main

import (
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
	tmpBuffer := make([]byte, 1024)
	for {
		n, err := conn.Read(tmpBuffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connected error:", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive data:", tmpBuffer[0:n])
		Log(conn.RemoteAddr().String(), "receive data string :", string(tmpBuffer[:n]))
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
