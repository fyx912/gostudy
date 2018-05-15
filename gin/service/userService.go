package service

import(
	// "database/sql"
	"gostudy/gin/dao"
	"log"
)

func IsLogin(username string,password string)  bool{
	flag := dao.IsSingin(username,password)
	log.Printf("UserService ========>>>> IsLogin: %v", flag)
	return flag
}