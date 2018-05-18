package service

import(
	"gostudy/gin/dao"
	"log"
)

func IsLogin(username string,password string)  bool{
	flag := dao.IsSingin(username,password)
	log.Printf("UserService ========>>>> IsLogin: %v", flag)
	return flag
}