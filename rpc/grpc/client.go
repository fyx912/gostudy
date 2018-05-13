package main

import (
	"log"

	dp "golearn/rpc/grpc/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:41005"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(" not connection: %s", err)
	}
	defer conn.Close()
	connDp := dp.NewDataClient(conn)
	name := defaultName
	read, err := connDp.SayHello(context.Background(), &dp.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln("SayHello cloud not greet : %s", err)
	}
	log.Printf("Greeting: %s", read.Name)
	readHi, err := connDp.Hi(context.Background(), &dp.HelloRequest{Name: "22"})
	if err != nil {
		log.Fatalln("Hi cloud not greet : %s", err)
	}
	log.Printf("Greeting: %s", readHi)
}
