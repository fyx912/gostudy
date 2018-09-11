package service

import(
	"gostudy/gin/database"
	"gostudy/gin/models"
)


func FindMenu() (menu []models.Menu){
	db :=database.Db
	db.Where("status=?",0).Find(&menu)
	return menu
}