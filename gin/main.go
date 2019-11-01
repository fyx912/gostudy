package main

import(
	"net/http"
	"gostudy/gin/routers"
	"gostudy/gin/database"
	"runtime"
	// "time"
	// "log"
)

func main(){
	runtime.GOMAXPROCS(4)
	defer database.CloseDB()
	router := routers.Router

	// router.Run(":8888")
	// http.Request.Header..Set("Last-Modified",time.Now())
	http.ListenAndServe(":8888", router)

	// server := &http.Serve{
	// 	Addr:           ":8888",
    //     Handler:        router,
    //     ReadTimeout:    10 * time.Second,
    //     WriteTimeout:   10 * time.Second,
    //     MaxHeaderBytes: 1 << 20,
	// }
	// log.Fatalln(server.ListenAndServe())
	
}
