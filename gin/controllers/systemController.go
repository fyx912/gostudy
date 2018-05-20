package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
)

func GetSystem(c *gin.Context){
	sysMap := make(map[string]interface{})
	sysMap["code"]=0
	sysMap["msg"]="success"
	sysMap["os"]=runtime.GOOS
	sysMap["arch"]=runtime.GOARCH
	hostname, _ := os.Hostname()
	sysMap["hostName"]=hostname
	sysMap["goroot"]=runtime.GOROOT()
	sysMap["version"]=runtime.Version()
	sysMap["numCPU"]=runtime.NumCPU()
	sysMap["NumGoroutine"]=runtime.NumGoroutine()
	c.JSON(http.StatusOK, sysMap)
}