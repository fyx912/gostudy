package dao

import(
"github.com/jinzhu/gorm"
	"gostudy/gin/model"
	"gostudy/gin/database"
	"log"
)

func  IsSingin(username string ,password string )bool{
	db := database.Db
	user := new(model.User)
	data := db.Where("username = ? AND password = ?", username,password).First(&user)
	log.Println("dao  data :",user)
	log.Println("data error:",data.Error)
	if data.Error != nil {
		log.Printf("UserDao  error ====> %s \n",data.Error)
		return false
	}else{
		log.Println(" UserDao =======>>> user:  ", data.Value)
		return true	
	}	
}

// func  FindUserAll()([]model.User, error){
// 	db := database.Db
// 	var user []*model.User
// 	data := db.Find(&user)

// 	if data.Error != nil{
// 		log.Printf(" userDao-FindUserAll error: %s \n", data.Error.Error())
// 		return nil,data.Error
// 	}else{
// 		return data.Value,nil
// 	}
// }

