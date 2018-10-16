package main

import (
	"context"
	"log"
	"net"

	"github.com/tokopedia/coba-grpc/hello"
	"github.com/tokopedia/iris/src/go/hub/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	log.Printf("REQUEST  %+v", *req)
	val, _ := data.GetTracker(ctx, 1)
	log.Println("Data : ", val)

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
