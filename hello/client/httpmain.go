package main

import (
	"context"
	"fmt"

	clienthttp "github.com/makersu/go-kit-example-hello/hello/client/http"

	"github.com/go-kit/kit/transport/http"
)

func main() {
	svc, err := clienthttp.New("http://localhost:8081", map[string][]http.ClientOption{})
	if err != nil {
		panic(err)
	}

	rs, err := svc.Hello(context.Background(), "Http Client")
	if err != nil {
		fmt.Println("HTTP Client Error:", err)
		return
	}
	fmt.Println("HTTP Client Result:", rs)

	// empty string
	rs, err = svc.Hello(context.Background(), "")
	if err != nil {
		fmt.Println("HTTP Client Error:", err)
		return
	}
	fmt.Println("HTTP Client Result:", rs)
}
