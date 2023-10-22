package main

import (
	"context"
	"flag"
	"fmt"
	"golang-poc/grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := hello.NewHelloClient(conn)

	greeting, err := client.Greet(context.Background(), &hello.Hello{HelloMessage: "This is a book"})
	if err != nil {
		log.Fatalf("Fail to call greet", err)
	}
	fmt.Println(greeting.Round)

}
