package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	str := "123\t\t"
	fmt.Println(strings.TrimSpace(str), len(str), len(strings.TrimSpace(str)))
	if len(strings.TrimSpace(str)) > 0 {
		fmt.Println(" dads")
	}
	jsonStr := "{\"Username\":\"admin\",\"Password\":\"123456\"}"
	fmt.Println("json=", jsonStr)
	var jsonUser User
	jsonerr := json.Unmarshal([]byte(jsonStr), &jsonUser)
	if jsonerr != nil {
		fmt.Println("json 不存在:", jsonerr)
	}
	fmt.Println("json 解析: name=", jsonUser.Username, "passwrd=", jsonUser.Password)

	var user User
	user.Username = "ding"
	user.Password = "123456"
	fmt.Println(user)

	lang, err := json.Marshal(user)
	if err != nil {

	}
	fmt.Println(string(lang))
}
