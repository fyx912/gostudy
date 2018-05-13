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
	sql.Open("mysql", databaseUrl)
	err := db.Ping()
	if err!=nil {
		log.Fatalln("Mysql databases be cuse:",err)
	}
}

func Query()  {
	
}