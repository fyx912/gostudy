package model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type User struct{
	Id int64   `gorm:"primary_key,AUTO_INCREMENT"`// 字段`ID`为默认主键
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Name string 
	Age int 
	Phone string 
	Date time.Time 
	}
	
func (u User) TableName() string{
	return "user"
}