package main

import (
	"fmt"
	"math/rand"
)

func rand_int() int {
	return rand.Int()
}

func rand_generator() chan int {
	out := make(chan int)
	go func() {
		for {
			out <- rand_int()
		}
	}()
	return out
}

//多路复用
func rand_generator2() chan int {
	// 创建两个随机数生成器服务
	rand_generator_1 := rand_generator()
	rand_generator_2 := rand_generator()

	out := make(chan int)

	go func() {
		for {
			out <- <-rand_generator_1
		}
	}()
	go func() {
		for {
			out <- <-rand_generator_2
		}
	}()
	return out
}

func main() {
	fmt.Println("rand :", rand.Int())
	rand_service_handler := <-rand_generator()
	fmt.Println("%d ", rand_service_handler)
	rand_service_handler2 := <-rand_generator2()
	fmt.Println("%dn", rand_service_handler2)
}
