package main

import (
	"fmt"
)

var quit chan bool = make(chan bool)

func main() {
	go testGoroutine()
	<-quit
}

func testGoroutine() {
	for i := 0; i < 10; i++ {
		fmt.Println(" hi....", i)
	}
	quit <- true
}
