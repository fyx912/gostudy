package main

// grpc server.go
import (
	"fmt"
	dp "golearn/rpc/grpc/hello"
	"log"
	"net"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	port = ":41005"
)

// server is used to implement helloworld.GreeterServer.
type Data struct{}

func main() {
	fmt.Println("grpc")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	server := grpc.NewServer()
	dp.RegisterDataServer(server, &Data{})
	server.Serve(listen)

	log.Printf("grpc server in: %s", port)
}

// Mothed definition
func (data *Data) SayHello(cxt context.Context, request *dp.HelloRequest) (response *dp.HelloReply, err error) {
	response = &dp.HelloReply{
		Name: "Hello " + request.Name,
	}
	return response, err
}
func (data *Data) Hi(cxt context.Context, requset *dp.HelloRequest) (response *dp.HelloReply, err error) {
	return &dp.HelloReply{Name: "Hi Hello world"}, err
}
