package main

import (
	"context"
	pb "grpc-go-practice/example/proto"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TLS 认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.crt", "../../keys/server.key")
	if err != nil {
		log.Fatalf("Failed to gen credentials %v", err)
	}

	// 实例化 grpc server, 并启用 TLS
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册 HelloService
	pb.RegisterHelloServer(s, HelloService)

	log.Println("listen on " + Address + " with TLS.")

	s.Serve(listen)
}
