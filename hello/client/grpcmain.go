package main

import (
	"context"
	"fmt"
	"log"

	grpckit "github.com/go-kit/kit/transport/grpc"
	client "github.com/makersu/go-kit-example-hello/hello/client/grpc"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	svc, err := client.New(conn, map[string][]grpckit.ClientOption{})
	if err != nil {
		panic(err)
	}

	rs, err := svc.Hello(context.Background(), "GRPC Client")
	if err != nil {
		fmt.Println("GRPC Client Error:", err)
		return
	}
	fmt.Println("GRPC Client Result:", rs)

	// empty string
	rs, err = svc.Hello(context.Background(), "")
	if err != nil {
		fmt.Println("GRPC Client Error:", err)
		return
	}
	fmt.Println("GRPC Client Result:", rs)

}
