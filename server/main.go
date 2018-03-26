//go:generate protoc -I helloworld --go_out=plugins=grpc:helloworld helloworld/helloworld.proto
//go:generate protoc -I helloworld --go_out=plugins=grpc:../client/helloworld helloworld/helloworld.proto
package main

import (
	"log"
	"net"

	"github.com/satooon/grpc-sample/server/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println("ctx:", ctx, req)
	return &helloworld.HelloReply{Message: req.GetName() + " HelloWorld"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
