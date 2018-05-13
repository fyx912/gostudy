package routers

import(
	"github.com/gin-gonic/gin"
	"goStudy/gin/controller"
	// "net/http"
)

func  Router() *gin.Engine{
	router := gin.Default()
	router.Static("static", "gin/static/**")
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// router.StaticFS("static", http.Dir("/home/ding/mygo/src/goStudy/gin/static"))

	router.LoadHTMLGlob("gin/views/*")
	// gin.New().GET("admin",controller.UserLoginHandler)
	router.GET("login",controller.UserLoginHandler)

	router.GET("index",controller.Index)

	return router
}