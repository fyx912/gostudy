package main

import(
	"net/http"
	"gostudy/gin/routers"
	"gostudy/gin/database"
	"runtime"
	// "time"
	"log"
)

func main(){
	runtime.GOMAXPROCS(4)
	defer database.CloseDB()

	// router.Run(":8888")
	// http.Request.Header.Set("Last-Modified",time.Now())
	log.Fatalln(http.ListenAndServe(":8888", routers.Router))
	
}
