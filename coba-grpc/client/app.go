package main

import (
	"context"
	"log"

	"github.com/baitpuisi/go-example-project/coba-grpc/hello"
	"google.golang.org/grpc"
)

const (
	address = "localhost:1010"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := hello.NewHelloClient(conn)

	// Make payment to the payment gateway and print out its response.
	request := &hello.HelloRequest{
		UserID:  1,
		TokenID: 12323,
	}

	r, err := c.HelloWorld(context.Background(), request)
	if err != nil {
		log.Fatalf("could not crack Token: %v", err)
	}

	log.Printf("crack token response: %s", r.String())
}
