package routers

import(
	"github.com/gin-gonic/gin"
	"goStudy/gin/controller"
)

func  Router() *gin.Engine{
	router := gin.Default()
	// router.LoadHTMLGlob("views/**")
	// gin.New().GET("admin",controller.UserLoginHandler)
	router.GET("login",controller.UserLoginHandler)

	return router
}