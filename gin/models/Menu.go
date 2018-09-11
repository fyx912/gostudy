package models

import()

type Menu struct{
	Id int64 //ID	`gorm:"primary_key,AUTO_INCREMENT"`// 字段`ID`为默认主键,自增
	Mid int64	//菜单id `gorm:"not null,unique"`	//自动不能为null,唯一索引
	ParentId int64 //父级Id `gorm:"not null"
	OrderBy int	//排序 `gorm:"not null"`
	Name string	//菜单名称 `gorm:"not null"`
	menu string	//导航栏分类 `gorm:"not null"`
	status int	//状态(0可用,1不可用)
	Descrition string	//描述
}

func (m Menu) TableName() string{
	return "sys_menu"
}