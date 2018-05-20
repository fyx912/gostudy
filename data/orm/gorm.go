package main

import(
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
	"time"
	"fmt"
)
//默认表名是`users`
type User struct{
  Id int64   `gorm:"primary_key,AUTO_INCREMENT"`// 字段`ID`为默认主键
  Username string 
  Password string 
  Name string 
  Age int 
  Phone string 
  Date time.Time `gorm: not null`
}
func (User) TableName() string{
	return "user"
}

//对应表名t_lottery_list
type LotteryList struct{
	Id int `gorm:"primary_key,AUTO_INCREMENT`
	LotteryId string 
	LotteryType string
	LotteryName string
	DateTime time.Time  `gorm: not null`
	Remarks string
}

func (LotteryList) TableName() string{
	return "t_lottery_list"
}

func init(){
	openDB()
	//defer closeDB()
}

var db *gorm.DB
func closeDB(){
	error := db.Close()
	if error != nil {
		panic(error.Error())
	}else{
		log.Println("Mysql Connect stop success......")
	}
}
func openDB(){
	var err error
	db,err = gorm.Open("mysql", "root:123456@/ding?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Databases error becuse:"+err.Error())
	}else{
		log.Println("Mysql Databases start......")
	}
	error := db.DB().Ping()
	if error != nil{
		panic("Ping mysql error becuse:"+error.Error())
	}else{
		fmt.Println("Ping mysql connect success......")
	}
}
func main()  {
    defer closeDB()
	//启用Logger，显示详细日志
	db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true)// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	
	var user User
	data := db.First(&user)
	if data.Error != nil{
		log.Printf("data error %s \n",data.Error.Error())
	}else{
		log.Printf("data %s \n", data.Value)
		fmt.Println(user)
	}

	db.Where("username = ? AND password = ?", "admin","123456").First(&user)
	
	fmt.Println("条件查询：",user)

	//修改表名
	gorm.DefaultTableNameHandler(db,"t_lottery_list")
	var lottery []LotteryList
	 lotteryData := db.Find(&lottery)
	 fmt.Println("lottery : ", lotteryData.Value)
	
	 var users []User
	 db.Find(&users)

	 fmt.Println(users)

	//  fmt.Println(" query userAll :",query)
}

//有问题，待解决
func QueryAll(object []interface{})([]interface{}){
	db.Find(&object)
	return object
}