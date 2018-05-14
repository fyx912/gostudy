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
  Id int64   `gorm:"primary_key,AUTO_INCREMENT`// 字段`ID`为默认主键
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
	defer closeDB()
}

var db *gorm.DB
func closeDB(){
	db.Close()
}
func openDB(){
	var err error
	db,err = gorm.Open("mysql", "root:123456@/ding?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Println("Databases error becuse:",err.Error())
	}
	error := db.DB().Ping()
	if error !=nil{
		log.Fatalln("Ping mysql error becuse:" ,error.Error())
	}
}
func main()  {
 	openDB()
    defer closeDB()
	//启用Logger，显示详细日志
	db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true)// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	
	var user User
	data := db.First(&user)
	fmt.Printf("data %s \n", data)
	fmt.Println(user)
	
	//修改表名
	gorm.DefaultTableNameHandler(db,"t_lottery_list")
	var lottery []LotteryList
	 lotteryData := db.Find(&lottery)
	 fmt.Println("lottery : ", lotteryData.Value)
	
	//  var users []User
	//  query := QueryAll(users)

	 fmt.Println(" query userAll :",query)
}

//有问题，待解决
func QueryAll(object []interface{})([]interface{}){
	db.Find(&object)
	return object
}