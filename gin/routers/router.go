package routers

import(
	"github.com/gin-gonic/gin"
	"gostudy/gin/controllers"
	"net/http"
)

func init(){
	htmlRouter()
	staticRouter()
	router()	
}

var(
	Router = gin.Default()
)

func router(){
	// Router.Use(gin.Logger())
	Router.POST("login",controllers.PostLogin)
	Router.GET("loginOut",func(c *gin.Context){c.HTML(200, "login.html", nil)})
	Router.GET("user",controllers.GetUser)
	Router.GET("user/:username",controllers.GetUserByName)
	Router.GET("index",func(c *gin.Context){
		c.HTML(200, "index.html", gin.H{
			"title": "Index",
		})
	})

	Router.GET("sys",controllers.GetSystem)
	Router.GET("map",controllers.GetMap)

}

/**Views HTML Get method*/
func htmlRouter(){
	Router.LoadHTMLGlob("views/*")
	Router.StaticFile("/", "views/login.html")
	Router.StaticFile("/login", "views/login.html")
	Router.StaticFile("/index.html", "views/index.html")
	Router.StaticFile("/charts.html", "views/charts.html")
	Router.StaticFile("/system.html", "views/system.html")
	Router.StaticFile("/elements.html", "views/elements.html")
	Router.StaticFile("/forms.html", "views/forms.html")
	Router.StaticFile("/meCenter.html", "views/meCenter.html")
	Router.StaticFile("/tables.html", "views/meCenter.html")
	Router.StaticFile("/typography.html", "views/meCenter.html")
}
/**Static Router Get method*/
func staticRouter(){
	// 显示当前文件夹下的所有文件/或者指定文件
	Router.StaticFS("show", http.Dir("."))
	Router.Static("static", "./static")
	Router.StaticFile("favicon.ico", "static/img/favicon.ico")
	
		// http.Handle("/views", http.StripPrefix("/views", 
	// 	http.FileServer(http.Dir("/home/ding/mygo/src/goStudy/gin/static/"))))
}