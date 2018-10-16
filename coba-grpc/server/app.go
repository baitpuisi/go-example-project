package main

import (
	"context"
	"log"
	"net"

	"github.com/baitpuisi/go-example-project/coba-grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Printf("REQUEST  %+v", *req)

	return &hello.HelloResponse{Success: true, Message: "success"}, nil
}

var port = ":1010"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("fail to start server ", err.Error())
	}

	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &server{})
	reflection.Register(s)

	log.Println("server is listening at port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	// protoc3 --go_out=plugins=grpc:. hello.proto

}
