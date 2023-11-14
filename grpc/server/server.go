package main

import (
	"context"
	"flag"
	"fmt"
	"golang-poc/grpc/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type helloServer struct {
	hello.UnimplementedHelloServer
}

var round int64 = 0

func (h helloServer) Greet(ctx context.Context, req *hello.Hello) (*hello.Greeting, error) {
	round++
	fmt.Println("Greet from socket_client: ", req.HelloMessage)
	return &hello.Greeting{Round: round}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	hello.RegisterHelloServer(grpcServer, helloServer{})
	fmt.Println("Grpc scoket_server serve at port: 8080")
	grpcServer.Serve(lis)
}
