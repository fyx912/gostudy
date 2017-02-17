package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	log.Println(" beegin  dial ...")
	//conn, err := net.Dial("tcp", "127.0.0.1:9999")
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9999", 2*time.Second)
	if err != nil {
		fmt.Println(" client dial error :", err)
		return
	}
	defer conn.Close()
	log.Println(" dial ok !")

	fmt.Println(conn)

	senddata := []byte("0 client send hello !")
	conn.Write(senddata)

	var b = make([]byte, 32)
	readData, err := conn.Read(b)
	if err != nil {
		log.Println(" read data error :", err)
	}

	log.Println("read data :", string(b[0:readData]))

}
