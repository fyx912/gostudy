package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id       int
	Username sql.NullString
	Password sql.NullString
}

var (
	db          *sql.DB
	databaseUrl = "root:123456@tcp(127.0.0.1:3306)/ding?charset=utf8"
)

// func init() {
// 	db, err := OpenDatabase()
// 	CheckError(err)
// 	// db.SetConnMaxLifetime(3000)
// 	//用于设置最大打开的连接数，默认值为0表示不限制
// 	db.SetMaxIdleConns(300)
// 	//用于设置闲置的连接数
// 	db.SetMaxOpenConns(50)
// 	// error := db.Ping()
// 	// CheckError(error)
// 	// defer db.Close()
// }

//打开数据库
func OpenDatabase() (db *sql.DB, err error) {
	return sql.Open("mysql", databaseUrl)
}

//检测errort
//主要go外部文件引用函数,函数名称头字母要大写func C
func CheckError(err error) {
	if err != nil {
		log.Println("err==", err)
	}
}

func main() {
	db, err := OpenDatabase()
	CheckError(err)
	Query(db)
}

func Query(db *sql.DB) {
	sql := "SELECT * FROM user "
	rows, err := db.Query(sql)
	defer rows.Close()
	CheckError(err)
	columns, err := rows.Columns()
	CheckError(err)

	log.Println("%s Clumens= %v", columns)

	//字典类型
	//构造scanArgs,values两个数组,scanArgs的每一个值指向valuses的值的地址
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	record := make(map[interface{}]string)
	for rows.Next() {
		//将数据保存record字典
		err := rows.Scan(scanArgs...)
		CheckError(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}
