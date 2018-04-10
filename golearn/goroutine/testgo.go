package main

import (
	"fmt"
	// "time"
	"sync"
)

var lock sync.Mutex
var count int

func add() int {
	count++
	return count
	// fmt.Println(count)
}

func main() {

	for i := 0; i < 100; i++ {
		lock.Lock()
		go add()
		lock.Unlock()

	}

	// time.Sleep(time.Second * 1)
	fmt.Println("sum = ", count)
}
