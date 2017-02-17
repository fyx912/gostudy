package main

import (
	"fmt"
	//"math/rand"
	"runtime"
	"time"
)

func producer(c chan int, max int) {
	// defer close(c)

	for i := 0; i < max; i++ {
		fmt.Println("producer:", i)
		c <- i
	}
}

func consumer(c chan int, max int) {
	for i := 0; i < max; i++ {
		v := <-c
		fmt.Println("consumer:", v)
	}
}

func autoProducer(c chan int) {
	var i int = 0
	go func() {
		for { //死循环，一直不停生产数据
			// rand_int := rand.Int()
			i += 1
			fmt.Println("auto producer:", i)
			c <- i
		}
	}()
}

func autoConsumer(c <-chan int) {
	for {
		select {
		case ws := <-c:
			fmt.Println("\t auto consumer:", ws)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(" \t\t chanel is full ! ")
		}
	}
}
func main() {
	c := make(chan int, 10)
	runtime.GOMAXPROCS(runtime.NumCPU())
	size := 10
	go producer(c, size)
	go consumer(c, size)

	// time.Sleep(1 * time.Second)
	fmt.Println("------------------------------------------------------")

	//持续生产与消费随机数, 一直跑起来
	go autoProducer(c)
	go autoConsumer(c)

	//为了保持主进程程不死掉，在main函数最下面加上这句。
	for {
		time.Sleep(1 * time.Second)
	}
}
