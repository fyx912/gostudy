package routers

import(
	"github.com/gin-gonic/gin"
	"gostudy/gin/controller"
	"net/http"
)

func  Router() *gin.Engine{
	router := gin.Default()
	router.Static("static", "static")
	router.StaticFile("/favicon.ico", "static/img/favicon.ico")
	router.StaticFS("mystatic", http.Dir("/home/ding/mygo/src/goStudy/gin/static/"))

	// http.Handle("/views", http.StripPrefix("/views", 
	// 	http.FileServer(http.Dir("/home/ding/mygo/src/goStudy/gin/static/"))))
	router.StaticFile("/system.html", "./views/ystem.html")

	router.LoadHTMLGlob("views/*")
	// gin.New().GET("admin",controller.UserLoginHandler)s
	router.GET("login",controller.GetLoginHandler)
	router.POST("login",controller.PostLogin)

	router.GET("index",controller.Index)
	router.GET("index.html",func (c * gin.Context)  {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// router.GET("charts.html",controller.Charts)
	// router.GET("elements.html",controller.Elements)
	// router.GET("forms.html",controller.Forms)
	// router.GET("meCenter.html",controller.MeCenter)
	// router.GET("system.html",func (c * gin.Context)  {
	// 	c.HTML(http.StatusOK, "system.html", nil)
	// })
	// router.GET("tables.html",func (c * gin.Context)  {
	// 	c.HTML(http.StatusOK, "tables.html", nil)
	// })
	// router.GET("typography.html",func (c * gin.Context)  {
	// 	c.HTML(http.StatusOK, "typography.html", nil)
	// })
	return router
}