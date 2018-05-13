package main

import (
	"log"
	"os"
)

func main(){
	fileName := "my_debug.log"
	logFile,err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln(" Open file error! ",err)
	}
	debug := log.New(logFile, "[Debug]", log.LstdFlags)
	debug.Println(" The debug message  here")
	debug.SetPrefix("[Info]")
	debug.Println("The Info message here ")
	debug.SetFlags(debug.Flags()|log.LstdFlags)
	debug.Println("The different prefix")
}