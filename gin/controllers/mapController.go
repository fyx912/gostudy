package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)
//http://restapi.amap.com/v3/place/text?key=您的key&keywords=北京大学&types=高等院校&city=北京&children=1&offset=20&page=1&extensions=all
func GetMap(c *gin.Context){
	url := "http://restapi.amap.com/v3/place/text?"
	url += "key=351ccff01cdfb5997ed7ac26e547d53b"
	url += "&keywords=中国体育彩票投注站"
	url += "&city=北京&children=1&offset=20&page=1&extensions=all"
	log.Printf("map url: %s",url)
	res ,_:= http.Get(url)
	body ,err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("GetMap:",err.Error())
	}
	// fmt.Print(string(body))
	var data map[string]interface{}
	err = json.Unmarshal(body,&data)
	if err != nil {
		log.Printf(err.Error())
	}
	c.JSON(http.StatusOK,data)
}
