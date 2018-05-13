package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		log.Println("Error : ", err.Error())
		os.Exit(1)
	}
}

func connHandle(conn *net.TCPConn) {
	ipAddress := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" client disconnectde :", ipAddress)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("byte data:", reader)
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(" server connection close ...")
			return
		}
		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	checkError(err)
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	defer tcpListen.Close()
	fmt.Println(" tcp server begin ..... ")
	checkError(err)
	fmt.Println(" tcp server wait connection ..... ")
	for {
		conn, err := tcpListen.AcceptTCP()
		fmt.Println(" client connection  success : ", conn.RemoteAddr().String())
		if err != nil {
			continue
		}
		go connHandle(conn)

		// conn.SetReadDeadline(time.Now().Add(time.Duration(10)) * time.Second)
	}

}
