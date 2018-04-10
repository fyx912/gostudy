package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var hosts = []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}

func main() {
	zookeeperConn, connChan, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		panic(err)
	}
	defer zookeeperConn.Close()
	for {
		isConnected := false
		select {
		case connEvent := <-connChan:
			if connEvent.State == zk.StateConnected {
				fmt.Println("connect to zookeeper server success!")
				getNode(zookeeperConn)
				// go watchCreateNode(connEvent)
			}
		case _ = <-time.After(time.Second * 5): //5 second not connect sucees be connect failed
			fmt.Println("connect to zookeeper server timeout!")
		}
		if isConnected {
			break
		}
	}
}
func getNode(conn *zk.Conn) {
	children, _, err := conn.Children("/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(children)
}
func watchCreateNode(ech <-chan zk.Event) {
	envent := <-ech
	fmt.Println("path:", envent.Path)
	fmt.Println("path:", envent.Server)
	fmt.Println("path:", envent.State)
	//fmt.Println("path:", envent.Type)
}
