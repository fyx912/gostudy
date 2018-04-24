package main 

import(
	"github.com/json-iterator/go"
	"fmt"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	Id       int `json:"-"`
	Username string `json:"username"`
	Password string	`json:"password"`
}

func main()  {
	jsonStr := "{\"Username\":\"admin\",\"Password\":\"123456\"}"
	var jsonUser  *User
	err := json.Unmarshal([]byte(jsonStr),&jsonUser)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println("解析json.  username:",jsonUser.Username,"\t psssword:",jsonUser.Password)

	user := new(User)
	user.Username = "tinTin"
	user.Password = "123456"
	jsonObject,err := json.Marshal(user);
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("转换JSON，json=",string(jsonObject))

	//获取JSON对象中的值
	str := json.Get(jsonObject, "username").ToString()
	fmt.Println(str)
}