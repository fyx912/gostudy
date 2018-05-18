package model

import (
	// "time"
	// "github.com/jinzhu/gorm"
)

type User struct{
	Id int64   `gorm:"primary_key,AUTO_INCREMENT"`// 字段`ID`为默认主键
  	Username string 
  	Password string 
  	Name string 
  	Age int 
  	Phone string 
  	Date string `gorm:"not null"`
}
	
func (u User) TableName() string{
	return "user"
}