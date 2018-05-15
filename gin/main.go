package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"gostudy/gin/routers"
	"gostudy/gin/database"
)

func main(){
	defer database.Close()
	router := routers.Router()
	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"It works")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("user:name",func (c *gin.Context)  {
		name := c.Param("name")
		message := "hello " +name +"!"
		c.String(http.StatusOK, message) 
	})
	// router.Run(":8888")
	http.ListenAndServe(":8888", router)

	
}
