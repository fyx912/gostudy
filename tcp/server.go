
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	buff := make([]byte, 1024)
	// senddata := []byte(" server send hello !")
	// conn.Write(senddata)
	buffData, _ := conn.Read(buff)
	fmt.Println(" msg : ", string(buff[0:buffData]))
	msg := string(buff[0:1])
	switch msg {
	case "0":
		fmt.Println(0)
		conn.Write([]byte(" server 识别为0。"))
	case "1":
		fmt.Println(1)
	case "2":
		fmt.Println(2)
	case "3":
		fmt.Println(3)
	default:
		fmt.Println(" 无法识别!!!")
		conn.Write([]byte(" server 无法识别!!!"))
	}
	log.Println(" connection success !!!")
}

func main() {
	log.Println(" beegin server listen ...")
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("listen error :", err)
		return
	}

	defer listen.Close()
	log.Println("listen ok !")
	for {
		time.Sleep(time.Second * 10)
		conn, err := listen.Accept()
		defer conn.Close()
		log.Println("request address :", conn.RemoteAddr().String(), " connection success!")
		if err != nil {
			fmt.Println(" accept error :", err)
			break
		}
		go handleConn(conn)
	}
}
