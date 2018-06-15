package main 

import(
	"os"
	"log"
	"path"
)

func main(){
	fileName := "logs/sys.log"
	checkFileExist(fileName)
	file , err := os.Create(fileName)
	defer file.Close()
	if err != nil{
		log.Println(err.Error())
	}else{
		log.Println("file is exitst!")
	}

	
}
func checkFileExist(fileName string){
	err := os.Mkdir(path.Dir(fileName),os.ModePerm)
	if err != nil {
		log.Println(err)
	}else{
		log.Println("mkdir ",fileName," success!")
	}
}