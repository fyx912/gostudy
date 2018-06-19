package main 

import(
	"gopkg.in/redis.v5"
	"fmt"
)

func createRedis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:	"127.0.0.1:6379",
		Password:	"",
		DB: 0	,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong,err)
	return client
}

func main()  {
	client := createRedis()
	client.Set("ding", "tintie", 0)

	fmt.Println(client.Get("ding").Result())
}