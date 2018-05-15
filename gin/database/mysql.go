package database

import(
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
)


var (
	db *gorm.DB
	databaseUrl = "root:123456@tcp(127.0.0.1:3306)/ding?charset=utf8"
)

func init(){
	OpenDB()
}

func OpenDB(){
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
		log.Println("Ping mysql connect success......")
	}
	// //用于设置最大打开的连接数，默认值为0表示不限制
	// db.DB().SetMaxIdleConns(200)
	// //用于设置闲置的连接数
	// db.DB().SetMaxOpenConns(20)
	// db.DB().SetMaxIdleConns(5)
	// db.DB().SetConnMaxLifetime(3000)
	// 全局禁用表名复数
	db.SingularTable(true)// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
}

func Close(){
	defer db.Close()
	log.Fatalln("Mysql connect close .....")
}