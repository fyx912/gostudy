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

func Charts(c *gin.Context)  {
	c.Header("Connect-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}

func Elements(c *gin.Context)  {
	c.Header("Connect-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "elements.html", gin.H{
		"title": "elements",
	})
}

func Forms(c *gin.Context)  {
	c.Header("Connect-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "forms.html", gin.H{
		"title": "forms",
	})
}

func MeCenter(c *gin.Context)  {
	c.Header("Connect-type", "text/html;charset=utf-8")
	c.HTML(http.StatusOK, "meCenter.html", gin.H{
		"title": "meCenter",
	})
}