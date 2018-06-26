package main

import (
	"context"
	pb "grpc-go-practice/example/proto"
	"log"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	// TLS 连接
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.crt", "HY")
	if err != nil {
		log.Fatalf("Failed to create credentials %v", err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"

	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(r.Message)
}
