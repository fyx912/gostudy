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
}