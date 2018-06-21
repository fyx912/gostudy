package routers

import(
	"github.com/gin-gonic/gin"
	"gostudy/gin/controllers"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	// "github.com/gin-contrib/sessions/redis"
	"gostudy/gin/common"
	"log"
)

var(
	Router = gin.Default()
)

func init(){
	r:= Router
	//头部中间件
	r.Use(headerMiddleware())
	router(r)
	htmlRouter()
	staticRouter()
}

func router(r *gin.Engine){
	store := memstore.NewStore([]byte("secret"))
	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	Router.Use(gin.Logger())

	r.GET("/test", func (c *gin.Context)  {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		sessionId := session.Get("sessionId")
		if sessionId == nil{
			sessionId = common.UUID()
			session.Set("sessionId", sessionId)
		}
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		mysession ,_:=c.Request.Cookie("mysession")
		sessionValue := session.Get("mysession")
		c.JSON(200, gin.H{"count": session.Get("count"),"session":mysession,"sessionValue:":sessionValue,"flag":count==session.Get("count"),"sessionId":sessionId})
	})

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

//自定义头部中间件
func headerMiddleware() gin.HandlerFunc{
	return func (c *gin.Context) {
		c.SetCookie("sesionId", common.UUID(), 0, "", "", true, true)
		c.Writer.Header().Set("sessionId", common.UUID())
		c.Writer.Header().Set("accept","application/json")
		// c.Writer.Header().Set("status",s)
		log.Println(c.Writer.Status())
	}
}