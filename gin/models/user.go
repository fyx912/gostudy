package models

import (
	"time"
)
/**用户*/
type User struct{
	Uid int64	//用户ID	`gorm:"primary_key,AUTO_INCREMENT"`// 字段`ID`为默认主键,自增
  	Username string	//账号	`gorm:"not null,unique"`	//自动不能为null,唯一索引
  	Password string	//密码	`gorm:"not null"`
  	Name string	//姓名	`gorm:"not null"`
  	Mobile string	//电话	
	Email string	//邮箱	
	Sex int8	//性别	
	Login_time time.Time //登陆时间	`gorm:"sql.null"`
	Last_login_time time.Time //上次登陆时间	`gorm:"sql.null"`
	Login_count int64 //登陆总次数	`gorm:"not null"`
	Login_ip *string //登陆ip
	Last_login_ip *string //上次登陆ip
	Create_time time.Time //创建时间	`gorm:"not null","-"`
	UpdateTime time.Time //修改时间	`gorm:"column:update_time,not null","-"`
}
	
func (u User) TableName() string{
	return "sys_user"
}