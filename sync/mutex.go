package main

import(
	"sync"
	"time"
	"fmt"
)


func main(){
	mutex()
	once()
	waitGroups()
}

var waitGroup sync.WaitGroup
func waitGroups(){
	for i := 0; i < 10; i++ {
		waitGroup.Add(1) //添加需要等待goroutine的数量
		go func ()  {
			fmt.Println("hehe")
			time.Sleep(time.Second)
			waitGroup.Done()//减少需要等待goroutine的数量,相当与Add(-1)
		}()
	}
	waitGroup.Wait()//执行阻塞，直到所有的需要等待的goroutine数量变成0
	fmt.Println("over")
}

func once(){
	 var syncOnce   sync.Once
	onecBody := func(){
		fmt.Println(" only once")
	}
	for i := 0; i < 10; i++ {
		go func(a int) {
			syncOnce.Do(onecBody)//只执行一次
		}(i)
	}
	time.Sleep(time.Millisecond * 200)
}


func mutex(){
	num := 0
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func(){
			mu.Lock()
			defer mu.Unlock()
			num +=1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(" sync mutex -------> num : ",num)//如果不加锁这里的num的值会是一个随机数而不是1000
}