package main

import (
	"time"
	"fmt"
)

func main()  {
	times := time.Now()
	fmt.Println(times)
	fmt.Println(times.Day())
	// t,_ :=time.Parse("2018-01-01", times)
	fmt.Println(times.Format("2006-01-02 15:04:05"))

	timeChan := make(chan int , 10)
	ticker := time.NewTicker(time.Second * 1)
	go tick(timeChan,ticker)
}

func tick(timeChan chan int,ticker *time.Ticker)  {
	for{
		select{
		case  <- ticker.C:
			fmt.Print(" ticker at %v\n",time.Now())
		}
	}
	 <-timeChan
}