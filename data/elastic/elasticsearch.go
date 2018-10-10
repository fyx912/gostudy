package main

import(
	"fmt"
	"log"
	"os"
	"context"

	"github.com/olivere/elastic"
)

var client *elastic.Client
var err error
var host = "http://192.168.2.10:9200"

func main(){
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}