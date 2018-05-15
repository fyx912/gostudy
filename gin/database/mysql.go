package database

import(
	"database/sql"
	// "github.com/go-sql-driver/mysql"
	"log"
)


var (
	db          *sql.DB
	databaseUrl = "root:123456@tcp(127.0.0.1:3306)/ding?charset=utf8"
)

func init(){
	db ,err := sql.Open("mysql", databaseUrl)
	defer db.c
	if err !=nil {
		log.Println("Mysql database error becuse :",err.Error())
	}
	//用于设置最大打开的连接数，默认值为0表示不限制
	db.SetMaxIdleConns(200)
	//用于设置闲置的连接数
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(3000)
	err = db.Ping()
	if err!=nil {
		log.Fatalln("Mysql databases error becuse:",err.Error())
	}
}
