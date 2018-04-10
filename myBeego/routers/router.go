package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myBeego/controllers"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func init() {
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/user", &controllers.LoginController{})

	//控制页面
	beego.Router("/index", &controllers.NavigationController{}, "*:Index")
	beego.Router("/meCenter", &controllers.NavigationController{}, "*:MeCenter")
	beego.Router("/system", &controllers.NavigationController{}, "*:System")
	beego.Router("/forms", &controllers.NavigationController{}, "*:Forms")
	beego.Router("/tables", &controllers.NavigationController{}, "*:Tables")
	beego.Router("/charts", &controllers.NavigationController{}, "*:Charts")
	beego.Router("/typography", &controllers.NavigationController{}, "*:Typography")
	beego.Router("/elements", &controllers.NavigationController{}, "*:Elements")

}
