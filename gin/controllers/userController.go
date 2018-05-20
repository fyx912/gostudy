package controllers

import (
	"time"
	"strings"
	"github.com/gin-gonic/gin"
	"gostudy/gin/service"
	"gostudy/gin/common"
	"log"
	// "encoding/json"
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
		log.Println(reqInfo)
		if  "" != reqInfo.Username  {
			if reqInfo.Password != "" {
				password := common.MD5(reqInfo.Password)
				if service.IsSingin(reqInfo.Username,password) {
					ip := string([]byte(this.Request.RemoteAddr)[:strings.Index(this.Request.RemoteAddr,":")])
					log.Printf("Address : %s" ,ip)
					log.Printf("username: %s ,password: %s", reqInfo.Username,reqInfo.Password)
					service.SaveIp(reqInfo.Username,ip)
					this.JSON(http.StatusOK, gin.H{"code":0,"msg":"success"})
				}else{
					this.JSON(http.StatusOK, gin.H{"code":1,"msg":"Failed! becesu: Incorrect user or password!"})
				}
			}else {
				log.Printf("Failed! becesu: Incorrect user password %s \n", reqInfo.Password)
				this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect password!"})
			}
		}else{
			log.Printf("Failed! becesu: Incorrect user name %s \n", reqInfo.Password)
			this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect user name!"})
		}
	}else{
		log.Printf("Failed! becesu:  %s \n", err.Error())
		this.JSON(200, gin.H{"code":2,"msg":"Failed! becesu: "+err.Error()})
	}
}

func GetUser(this *gin.Context){
	this.Set("Accept-Encoding", "gzip")
	this.Set("Last-Modified/ETag", time.Now())
	jsonMap := make(map[string]interface{})
	jsonMap["code"] = 0
	jsonMap["msg"] = "success"
	user ,err := service.FindUserAll()
	if err != nil {
		log.Println(err,user)
		jsonMap["data"]= nil
	}else{
		jsonMap["data"]= user
	}
	this.JSON(http.StatusOK,jsonMap)
}	

func  GetUserByName(c *gin.Context){
	username := c.Param("username")
	jsonMap := make(map[string]interface{})
	jsonMap["code"] = 0
	jsonMap["msg"] = "success"
	user ,err := service.FindUserByName(username)
	if err != nil {
		log.Println(err,user)
		jsonMap["data"]= nil
	}else{
		jsonMap["data"]= user
	}
	c.JSON(http.StatusOK,jsonMap)
}	