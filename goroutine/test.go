package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop(id int) { // id: 该goroutine的标号
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(3) // 最多同时使用3个核

	for i := 0; i < 3; i++ { //开三个goroutine
		go loop(i)
	}

	for i := 0; i < 3; i++ {
		<-quit
	}
}
