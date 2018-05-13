package model

import (
	"time"
	"encoding/json"
	"github.com/go-xorm/xorm"
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