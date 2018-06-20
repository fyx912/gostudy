package main 

import(
	"gopkg.in/redis.v5"
	"fmt"
	"time"
	"sync"
)

func createRedis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:	"127.0.0.1:6379",
		Password:	"",
		DB: 0	,
		PoolSize: 5,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong,err)
	return client
}

func main()  {
	client := createRedis()
	defer client.Close()

	// listOperation(client)
	// stringOperation(client)
	setOperation(client)
	hashOperation(client)
	connectPool(client)
}

func stringOperation(client *redis.Client){
	// 第三个参数是过期时间, 如果是0, 则表示没有过期时间.
	err := client.Set("username", "admin", 0).Err()
	panicErr(err)
	val , err := client.Get("username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("username:",val)

	//设置过期时间
	error := client.Set("age", 21, 1 * time.Second).Err()
	panicErr(error)

	client.Incr("age")//自增
	client.Incr("age")//自增
	client.Decr("age")//自减

	val ,err = client.Get("age").Result()
	panicErr(err)
	fmt.Printf("age : %v\n", val)

	time.Sleep(2 * time.Second)
	val ,err = client.Get("age").Result()
	if err != nil {
		fmt.Printf("error: %v\n",err)
	}
	fmt.Printf("age : %v\n", val)
}

func listOperation(client *redis.Client){
	client.RPush("fruit", "apple")//在key的list尾部添加一个value的元素
	client.RPush("fruit", "mango")
	client.LPush("fruit", "banana")//在名称为 fruit 的list头添加一个值为value的 元素

	length , err := client.LLen("fruit").Result() //返回名称为 fruit 的list的长度
	panicErr(err)
	fmt.Printf(" list length: %v \n", length)

	val , err := client.LRange("fruit", 0,length).Result()//返回名称为key的list中start至end之间的元素
	panicErr(err)
	fmt.Printf(" list values: %v \n", val)

	value , err := client.LPop("fruit").Result()//返回并删除名称为key的list中的首元素
	panicErr(err)
	tail_value , err := client.RPop("fruit").Result()//返回并删除名称为key的list中的尾元素
	panicErr(err)

	fmt.Printf(" list remove  start values : %v  ,remove tail value ：%v  \n",value,tail_value)

}

func setOperation(client *redis.Client){
		client.SAdd("blacklist","Obama") // 向 blacklist 中添加元素
		client.SAdd("blacklist","Hillary","the Elder","James")

		client.SAdd("whitelist", "the Elder","James") // 向 whitelist 添加元素

		blacklistValues , err := client.SMembers("blacklist").Result()//获取blacklist的set所有元素
		panicErr(err)
		fmt.Printf("set operation  blacklist values : %v \n",blacklistValues)

		// 判断元素是否在集合中
		isMember, err := client.SIsMember("blacklist", "James").Result()
		panicErr(err)
		fmt.Println("Is Bush in blacklist: ", isMember)

		// 求交集, 即既在黑名单中, 又在白名单中的元素
		names, err := client.SInter("blacklist", "whitelist").Result()
		if err != nil {
			panic(err)
		}
		// 获取到的元素是 "the Elder"
		fmt.Printf("Inter result: %v \n", names)
}

func hashOperation(client *redis.Client){
	client.HSet("user_tin", "name", "tin")// 向key 的 hash 中添加元素 name
	client.HSet("user_tin","age", 18)// 向key 的 hash 中添加元素 age
	client.HSet("user_tin","address", "ShenZhen") // 向key 的 hash 中添加元素 address

	//获取hash中所有的键（field）及其对应的value
	valAll ,err := client.HGetAll("user_tin").Result()
	panicErr(err)
	fmt.Printf("HSAH operation ---->  user_tin all: %v \n",valAll)


	// 获取名为 user_tin 的 hash 中的字段个数
    length, err := client.HLen("user_tin").Result()
    panicErr(err)
    fmt.Printf("HSAH operation ----> field count in user_tin: %v\n", length) // 字段个数为3


	//批量地向名称为 user_test 的 hash 中添加元素 name 和 age
	client.HMSet("user_test", map[string]string{"name":"admin","age":"20"})//批量添加

	// 批量获取名为 user_test 的 hash 中的指定字段的值.
    fields, err := client.HMGet("user_test", "name", "age").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("HSAH operation    ---fields in user_test: %v\n", fields)



    // 删除名为 user_test 的 age 字段
    client.HDel("user_test", "age")
    age, err := client.HGet("user_test", "age").Result()
    if err != nil {
        fmt.Printf("HSAH operation ----> Get user_test age error: %v\n", err)
    } else {
        fmt.Printf("HSAH operation ----> user_test age is: ", age) // 字段个数为2
	}

	valAll ,err = client.HGetAll("user_test").Result()
	panicErr(err)
	fmt.Printf("HSAH operation ---->  useruser_test_tin all: %v \n",valAll)
}


func connectPool(client *redis.Client) {
    wg := sync.WaitGroup{}
    wg.Add(10)

    for i := 0; i < 10; i++ {
        go func() {
            defer wg.Done()

            for j := 0; j < 100; j++ {
                client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
                client.Get(fmt.Sprintf("name%d", j)).Result()
            }

            fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns);
        }()
    }

    wg.Wait()
}

func panicErr(err error){
	if err != nil {
		panic(err)
	}
}