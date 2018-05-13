package controller

import (
	"log"
	// "encoding/json"
	"fmt"
	
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponseBody struct {
	Username string `form:"user" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	// code string `form:"code" json:"code" binding:"required"`
}

func GetLoginHandler(this *gin.Context) {
	// this.Ctx.WriteString(" hello world")
	// this.HTML(http.StatusOK, "login.html",gin.H)
	// this.JSON(http.StatusOK, gin.H{
	// 	"message": "login",
	// })
	this.Header("Content-type", "text/html;charset=utf-8")
	this.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
func PostLogin(this * gin.Context){
	var reqInfo LoginResponseBody
	if err := this.BindJSON(&reqInfo); err == nil {
		fmt.Println(reqInfo)
		if reqInfo.Username == "admin"  {
			if reqInfo.Password == "123456" {
				fmt.Printf("username: %s ,password: %s", reqInfo.Username,reqInfo.Password)
				this.JSON(http.StatusOK, gin.H{"code":0,"msg":"success"})
			}else {
				log.Printf("Failed! becesu: Incorrect user password %s \n", reqInfo.Password)
				this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect password"})
			}
		}else{
			log.Printf("Failed! becesu: Incorrect user name %s \n", reqInfo.Password)
			this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect user name"})
		}
	}else{
		log.Printf("Failed! becesu:  %s \n", err.Error())
		this.JSON(200, gin.H{"code":2,"msg":"Failed! becesu : "+err.Error()})
	}
}
