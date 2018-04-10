package main

import (
	"fmt"
	"time"
)

const (
	dateFrom = "2006-01-02"
)

func main() {
	fmt.Println("time : ", time.Now().Format("2006-01-02"))
	t, _ := time.Parse(dateFrom, time.Now().String())
	fmt.Println("time 2 : ", t)
	c := make(chan bool)
	go func() {
		Go()
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func Go() {
	fmt.Println(" Go Go Go !!!")
}
