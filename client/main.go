package main

import (
	"context"
	"log"

	"github.com/satooon/grpc-sample/client/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	req := &helloworld.HelloRequest{Name: "Taro"}
	res, err := client.SayHello(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success:", res)
}
