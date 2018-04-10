package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	var i = "hello word"
	go func(a string) {
		fmt.Println("1", a)
	}(i)
	fmt.Println("2")
	go fmt.Println(" go 1 你好")

	ch <- " channel"
	go func() {
		fmt.Println(" channel :", <-ch)
	}()
	time.Sleep(1 * time.Second)

}
