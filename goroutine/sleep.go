package main

import (
	"fmt"
	"time"
)

func sleep(i int) {
	for ; ; i += 2 {
		fmt.Println(i, "sleep")
	}
}

func main() {
	go sleep(1)
	go sleep(2)
	time.Sleep(time.Millisecond)
}
