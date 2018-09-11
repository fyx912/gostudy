package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"gostudy/gin/service"
	"gostudy/gin/common"
)

func GetMenu(c *gin.Context)  {
	menu := service.FindMenu()
	c.JSONP(http.StatusOK,common.Json(menu))
}
