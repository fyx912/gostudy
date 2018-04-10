package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id       int
	Username string
	Password string
}

//打开数据库
func OpenDatabase() (db *sql.DB, err error) {
	return sql.Open("mysql", "root:123456@tcp(192.168.0.103.1:3306)/ding?charset=utf8")
}

//数据库关闭连接
func CloseDatebase() {
	db, err := OpenDatabase()
	CheckError(err)
	db.Close()
}

//检测errort

//主要go外部文件引用函数,函数名称头字母要大写func C
func CheckError(err error) {
	if err != nil {
		log.Println("err==", err)
		fmt.Println("err==", err)
	}
}

func main() {
	db, err := OpenDatabase()
	CheckError(err)
	rows, err := db.Query(" SELECT user,password FROM User ")
	CheckError(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	fmt.Println("columns=", columns)

	for rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		CheckError(err)
		// User.Username = username
		// User.Password = password
		fmt.Printf("%s is %v \n", username, password)
	}
}
