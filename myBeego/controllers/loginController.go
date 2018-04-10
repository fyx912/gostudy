package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/session"
	"log"
	"myBeego/models"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	jsonCode := "{\"code\":1,\"msg\":\"failed\"}"
	var user models.User
	fmt.Println("json=", string(this.Ctx.Input.RequestBody))
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Println("err:", err)
	}
	username := strings.TrimSpace(user.Username)
	password := strings.TrimSpace(user.Password)
	fmt.Println("user=",user)
	log.Println("user:", user.Username, user.Password)
	if len(username) > 0 || len(password) > 0 {
		if username == "admin"|| username == "ding" {
			if password == "123456" {
				//只有通过登录才能保存session
				this.SetSession(username, this.Ctx.Input.Cookie("beegosessionID"))
				this.SetSession("sessionID", this.Ctx.Input.Cookie("beegosessionID"))
				fmt.Println("sessionId=", this.GetSession(username))
				jsonCode = "{\"code\":0,\"msg\":\"success\"}"
			} else {
				jsonCode = "{\"code\":1,\"msg\":\"password error\"}"
			}
		} else {
			jsonCode = "{\"code\":1,\"msg\":\"username error\"}"
		}
	} else {
		jsonCode = "{\"code\":1,\"msg\":\"username or password not nil\"}"
	}
	this.Ctx.WriteString(jsonCode)
}

func (this *LoginController) Get() {
	// this.Ctx.WriteString(" hello world")
	this.TplName = "login.html"
}
