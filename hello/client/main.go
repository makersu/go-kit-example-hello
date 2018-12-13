package main

import (
	"context"
	"fmt"

	client "github.com/makersu/go-kit-example-hello/hello/client/http"

	"github.com/go-kit/kit/transport/http"
)

func main() {
	svc, err := client.New("http://localhost:8081", map[string][]http.ClientOption{})
	if err != nil {
		panic(err)
	}

	rs, err := svc.Hello(context.Background(), "Client World")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", rs)
}
