package main


import (
  _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "log"
    "encoding/json"
    "time"
    "fmt"
)

type User struct{
  Id int64   `xorm:"pk autoincr 'id'" json:"id"`
  Username string `xorm:"varchar(16)" json:"username"`
  Password string `xorm:"varchar(16)" json:"passwrod"`
  Name string `xorm:"varchar(16)" json:"name"`
  Age int `json:"age"`
  Phone string `xorm:"varchar(16)" json:"pone"`
  Date time.Time `xorm:"date" json:"date"  time.format:"2018-01-02"`
}

func main()  {
  engin,err := xorm.NewEngine("mysql", "root:123456@/ding?charset=utf8")
  if err!=nil {
    log.Println(err)
  }
  engin.Ping()

  engin.ShowSQL(true)
  //如果需要设置连接池的空闲数大小
  engin.SetMaxIdleConns(20)
  //如果需要设置最大打开连接数
  engin.SetMaxOpenConns(200)

  user := new(User)
  //获取一条记录
  lines,err := engin.Id(11).Get(user)
  //删除一条记录
  engin.Id(11).Unscoped().Delete(user)
  
  if err!=nil {
    log.Println(err)
  }
  fmt.Println("获取一行数据:",lines,*user)
  user.Password="12345666"
  user.Date=time.Now()
  id ,_:= engin.Id(5).Update(user)
  fmt.Println(" 更新",id)

  userTotal := new(User)
  // total ,_:=engin.Where("id>?",0).Count(userTotal)
  total ,_:=engin.Count(userTotal)
  fmt.Println("count :",total)

  //数组接收多行值
  userArray := make([]User,0)
  engin.Find(&userArray)
  fmt.Println("slice "," size :",len(userArray),"\n 获取所有的数据:",userArray)
  jsonArray ,_:= json.Marshal(userArray)
  fmt.Println("JSON----Array:",string(jsonArray))
  //map 接收多行值
  userMap := make(map[int64]User)
  engin.Find(&userMap)
  fmt.Println("size :",len(userMap),"\n 获取所有的数据:",userMap)
   
  jsonMap ,_:= json.Marshal(userMap)
  fmt.Println("map 转JSON :",string(jsonMap))

  //使用sql语句
  result,_:= engin.Query("select * from user")
  jsonSql,_:=json.Marshal(result)
  fmt.Println("sql查询", string(jsonSql))
  // resultJson :=sql(result)
  // fmt.Println(resultJson)
}

//[]map[string][]byte 转JSON
func sql(result []map[string][]byte ) string{
   length := len(result)
  paresArray := make([]map[string]string,0)
   for i := 0; i < length; i++ {
     paresMap :=result[i]
     resultMap := make(map[string]string)
     for keys := range paresMap {
      //  fmt.Println("key:",keys,"\t valuse",string(paresMap[keys]))
       resultMap[keys]=string(paresMap[keys])
     }
     paresArray=append(paresArray,resultMap)
   }
   jsonData,_:=json.Marshal(paresArray)
  return string(jsonData)
}