package main

import (
	// "strings"
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/fyx912/mahonia"
)

const(
	address = "http://www.szse.cn/"
	sz_url = address+"main/marketdata/jypz/colist/"
)

func main(){
	//Requst the HTML page
	res ,err := http.Get(sz_url)
	CheckError(err)
	defer res.Body.Close()
	if res.StatusCode !=200  {
		log.Fatalf("status code error: %d %s",res.StatusCode,res.Status)
	}
	//Load the HTML document
	doc ,err := goquery.NewDocumentFromReader(res.Body)
	CheckError(err)

	table := doc.Find(".cls-title-table-common.cls-title-table").Find("tbody")
	table.Find("tr").Find("td").First().Find("div").Each(func (i int,s *goquery.Selection){
		titleValue := s.Text()
		onclick, _:= s.Attr("onclick")
		//defalut  
		tabKey := "tab1"
		if  onclick != "" {
			tabKey = onclick[22:len(onclick)-2]
		}
		fmt.Printf("Review Title %d: %s , onclick: %s \n", i,Decoder(titleValue),tabKey)
		if Decoder(titleValue)=="上市公司列表" {
			table := doc.Find(".cls-data-table-common.cls-data-table").Find("tbody").Find("tr")
			table.Find("th").Each(func(i int ,s *goquery.Selection){
				fmt.Printf("name %d: %s \n",i,Decoder(s.Text()))
			})
			table.Find("td").Each(func(i int ,s *goquery.Selection){
				brand :=  s.Text()
				fmt.Printf("data %d: %s \n",i,Decoder(brand))
			})
		}
	})	
	
	//Get HTML A  label 
	val,tag:=table.Find("td").End().Find("a").Attr("href")
		if tag {
			fmt.Println(" get  <a> the address : ",(address+val))
		}


	//HTML calss if exists
	menu := doc.Find("div").HasClass("x_menu")
	fmt.Println(" HTML calss exists : ",menu)
	doc.Find("div").Find("ul").Find("li").Find("b").Each(func (i int,s *goquery.Selection)  {
		fl := s.Text()
		fmt.Printf("Review %d: %s \n",i,Decoder(fl))
	})
}

func getData(){
	
}

func CheckError(e error){
	if e != nil {
		log.Fatal(e)
	}
}

/*GBK to UTF-8*/
func Decoder(src string) string {
	dec := mahonia.NewDecoder("GBK")
	return 	dec.ConvertString(src)
}