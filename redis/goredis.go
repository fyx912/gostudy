package main

import (
	"fmt"
	"github.com/goredis"
	"strconv"
)

func main() {
	var client goredis.Client
	// 设置端口为redis默认端口
	client.Addr = "127.0.0.1:6379"

	//字符串操作
	client.Set("name", []byte("fyx912"))
	client.Set("age", []byte(strconv.Itoa(26)))
	client.Set("a", []byte("hello"))
	val, _ := client.Get("a")
	//获取多个值
	users, _ := client.Mget("name", "age")
	for i, v := range users {
		fmt.Println(i, ":", string(v))
	}
	fmt.Println(string(val))
	client.Del("a")

	// //hash操作
	// // client.Hmset("user")
	// user := client.Hgetall("user", "name")
	// fmt.Println(" redis   操作hash:", user)
	//list操作
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("l")
}
