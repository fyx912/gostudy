package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

/**
 * [loop description]允许Go使用多核(runtime.GOMAXPROCS)
 * @return {[type]} [description]
 */
func loop() {
	for i := 0; i < 1000; i++ { //为了观察，跑多些
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(2) // 最多使用2个核

	fmt.Println("  返回当前系统的cpu核数:", runtime.NumCPU())
	go loop()
	go loop()
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
}

/**
 * 手动显式调动(runtime.Gosched)
 */
// func loop() {
// 	for i := 0; i < 10; i++ {
// 		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
// 		fmt.Printf("%d ", i)
// 	}
// 	quit <- 0
// }

// func main() {

// 	go loop()
// 	go loop()

// 	for i := 0; i < 2; i++ {
// 		<-quit
// 	}
// }
