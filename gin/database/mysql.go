package database

import(
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
)
var (
	Db *gorm.DB
	databaseUrl = "root:123456@tcp(127.0.0.1:3306)/ding?charset=utf8&parseTime=True&loc=Local"
)

func init(){
	openDB()
}

func openDB(){
	var err error
	Db,err = gorm.Open("mysql",databaseUrl)
	if err != nil {
		panic("Databases error becuse:"+err.Error())
	}else{
		log.Println("Mysql Databases start......")
	}
	error := Db.DB().Ping()
	if error != nil{
		panic("Ping mysql error becuse:"+error.Error())
	}else{
		log.Println("Ping mysql connect success......")
	}
	//用于设置最大打开的连接数，默认值为0表示不限制
	Db.DB().SetMaxIdleConns(200)
	//用于设置闲置的连接数
	Db.DB().SetMaxOpenConns(20)
	Db.DB().SetMaxIdleConns(5)
	Db.DB().SetConnMaxLifetime(3000)
	// 全局禁用表名复数
	Db.SingularTable(true)// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	Db.LogMode(true)
}

func CloseDB(){
	defer Db.Close()
	log.Fatalln("Mysql connect close .....")
}