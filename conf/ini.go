package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	iniconf, err := config.NewConfig("ini", "ini.conf")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	username := iniconf.String("user::username")
	password := iniconf.String("user::password")
	fmt.Println("获取init文件中的值:\t name:", username, "\t password:", password)
}
