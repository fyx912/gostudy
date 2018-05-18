package service

import(
	"gostudy/gin/dao"
	"log"
	"gostudy/gin/model"
)

func IsLogin(username string,password string)  bool{
	flag := dao.IsSingin(username,password)
	log.Printf("UserService ========>>>> IsLogin: %v", flag)
	return flag
}

func FindUserAll() (users []model.User ,err error){
	user ,err := dao.FindUserAll();
	if err!= nil {
		return  nil,err
	}
	return user,err
}