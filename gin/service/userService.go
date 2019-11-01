package service

import(
	"gostudy/gin/models"
	"gostudy/gin/database"
	"log"
	"time"
)

func  IsSingin(username string ,password string )bool{
	db := database.Db
	user := new(models.User)
	data := db.Where("username = ? AND password = ?", username,password).First(&user)
	log.Println("dao  data :",user)
	if data.Error != nil {
		log.Printf("userSerive  error ====> %s \n",data.Error)
		return false
	}else{
		log.Println(" userSerive ====IsSingin===>>> user: ",data.Value)
		return true	
	}	
}
/**保存登陆的Ip*/
func  SaveIp(username string ,ip string ){
	db := database.Db
	user := new(models.User)
	userMap := make(map[string]interface{})
	userMap["login_time"] = time.Now() 
	userMap["login_ip"] = ip
	db.Where("username = ?", username).Select("login_count").First(&user)
	log.Println(" login count :",user.Login_count)
	userMap["login_count"] = user.Login_count + 1
	userMap["update_time"] = time.Now() 
	db.Model(&user).Where("username = ?", username).Updates(userMap)
	log.Println("userSerive --->SaveIp---> data :",user)
}


var userArray []*models.User
//Find User All
func  FindUserAll()(userArray []models.User,err error){
	db := database.Db
	data := db.Find(&userArray)
	if data.Error != nil{
		log.Printf(" userSerive-FindUserAll error: %s \n", data.Error.Error())
		return nil,data.Error
	}else{
		log.Println(" userSerive =====FindUserAll==>>> user:  ", data.Value)
		return userArray,nil
	}
}
//Find user by username
func  FindUserByName(usernmae string)(userArray []models.User,err error){
	db := database.Db
	data := db.Where("username = ?", usernmae).Find(&userArray)
	if data.Error != nil{
		log.Printf(" userSerive---FindUserByName---->error: %s \n", data.Error.Error())
		return nil,data.Error
	}else{
		log.Println(" userSerive =====FindUserByName==>>> user:  ", data.Value)
		return userArray,nil
	}
}

