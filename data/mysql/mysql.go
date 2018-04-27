package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/json"
)

type User struct {
	Id       int
	Username sql.NullString
	Password sql.NullString
}

var (
	db          *sql.DB
	databaseUrl = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
)
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
	//用于设置最大打开的连接数，默认值为0表示不限制
	db.SetMaxIdleConns(300)
	//用于设置闲置的连接数
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(3000)
	error := db.Ping()
	CheckError(error)
	defer db.Close()

	CheckError(err)
	mapJson := Query(db)

	fmt.Println(mapJson)
	jsonData,err := json.Marshal(&mapJson)
	CheckError(err)
	fmt.Println("json :",string(jsonData))

	// db.Exec("query", args)
}

func UserBy(db *sql.DB)  {
	// sql = " select * from user where  id=?"
	// db.Query(query, args)
}

func Query(db *sql.DB) []map[string]string {
	sql := "SELECT * FROM user "
	rows, err := db.Query(sql)
	defer rows.Close()
	CheckError(err)
	columns, err := rows.Columns()
	CheckError(err)

	log.Println("Clumens=", columns)

	//字典类型
	//构造scanArgs,values两个数组,scanArgs的每一个值指向valuses的值的地址
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var numbers []map[string]string

	record := make(map[string]string)
	for rows.Next() {
		//将数据保存record字典
		err := rows.Scan(scanArgs...)
		CheckError(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		numbers = append(numbers,record)
	}
	return numbers
}
