package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.40.116/100mshCloud/")
	if err != nil {
		fmt.Println(" error=", err)
	}
	defer resp.Body.Close()
	fmt.Println("  body=", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" error=", err)
	}
	fmt.Println(body)

}
