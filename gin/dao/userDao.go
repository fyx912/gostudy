package dao

import(
	"github.com/jinzhu/gorm"
	"gostudy/gin/model"
	"log"
)
var (
	db *gorm.DB
	user *model.User
)
func  IsSingin(username string ,password string )bool{
	log.Printf(" UserDao =======>>> username: %s , password: %s  \n", username,password)
	data := db.Where("username = ? AND password = ?", username,password).First(&user)
	log.Fatalln("data:",user)
	if data.Error != nil {
		log.Printf("UserDao  error %s \n",data.Error.Error())
		return false
	}else{
		log.Println(" UserDao =======>>> user:  ", data.Value)
		return true	
	}	
}
