package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("  redis 连接失败!")
	} else {
		fmt.Println("redis 连接成功!")
	}
	//操作字符串(string)
	conn.Do("SET", "user:user0", 123)
	conn.Do("SET", "user:user1", 456)
	conn.Do("APPEND", "user:user0", 87)
	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	user1, err := redis.Int(conn.Do("GET", "user:user1"))

	fmt.Printf("user0 is %d , user1 is %d \n", user0, user1)

	//操作哈希表(hash)
	var p1, p2 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "golang  redisGo"
	p1.Author = "fyx912"
	p1.Body = "Hello"
	if _, err := conn.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		fmt.Println(" hmset error:", err)
	}

	id1, err := redis.Values(conn.Do("hgetall", "id1"))
	checkErr(err)

	err = redis.ScanStruct(id1, &p2)
	checkErr(err)

	fmt.Println("go redisGo hash操作:")
	fmt.Printf("%+v\n", p2)
	hlenSize, err := conn.Do("Hlen", "id1")
	checkErr(err)
	fmt.Println("  id1的hash sezi=", hlenSize)

	//操作列表(list)
	vals := []string{"hello", "word", "fyx912", "ding", "redisgo-list"}
	for _, v := range vals {
		_, err = conn.Do("Rpush", "redislist", v)
		checkErr(err)
	}
	//读取list
	redislist, err := redis.Values(conn.Do("Lrange", "redislist", 0, 100))
	checkErr(err)
	list, err := redis.Strings(redislist, err)
	checkErr(err)
	fmt.Println(" list : ", list)
	//获取list大小
	listzise, err := conn.Do("LLEN", "redislist")
	checkErr(err)
	fmt.Println("list  size: ", listzise)

	//操作集合(SET)
	setVals := []string{"hello", "word", "fyx912", "ding", "redisgo-list"}
	for _, v := range setVals {
		_, err = conn.Do("SADD", "sets", v) //添加一个或多个
		checkErr(err)
	}
	//读取set
	sets, err := redis.Values(conn.Do("smembers", "sets"))
	checkErr(err)
	getSET, err := redis.Strings(sets, err)
	fmt.Println(" 获取set :", getSET)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(" err==", err)
	}
}
