package controller


import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.Header("Connect-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}