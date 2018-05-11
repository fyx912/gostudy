package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginHandler(this *gin.Context) {
	// this.Ctx.WriteString(" hello world")
	// this.HTML(http.StatusOK, "login.html",gin.H)
	this.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
	// this.Header("Content-type", "text/html;charset=utf-8")
	// this.HTML(http.StatusOK, "views/login.html", gin.H{
	// 	"title": "Login",
	// })
}
