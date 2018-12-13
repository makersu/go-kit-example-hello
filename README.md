# go-kit-example-hello-world

## Install kit
```
go-kit-example-hello> go get github.com/kujtimiihoxha/kit
```

## Create a new service
```
# kit new service hello
go-kit-example-hello> kit n s hello
```

## Define service (string/pkg/service/service.go)
```
type HelloService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Hello(ctx context.Context, s string) (rs string, err error)
}
```

## Implement service
```
// TODO implement the business logic of Hello
func (b *basicHelloService) Hello(ctx context.Context, s string) (rs string, err error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return "Hello " + s, nil
}
```

## Run service
```
go-kit-example-hello> go run hello/cmd/main.go
ts=2018-12-13T05:34:39.06727Z caller=service.go:78 tracer=none
ts=2018-12-13T05:34:39.067779Z caller=service.go:100 transport=HTTP addr=:8081
ts=2018-12-13T05:34:39.068115Z caller=service.go:134 transport=debug/HTTP addr=:8080
ts=2018-12-13T05:34:42.175735Z caller=middleware.go:27 method=Hello s=world rs="Hello, world" err=null
ts=2018-12-13T05:34:42.175783Z caller=middleware.go:33 method=Hello transport_error=null took=56.466µs
```

## Test service
```
go-kit-example-hello> curl -XPOST -d'{"s":"world"}' localhost:8081/hello
{"rs":"Hello, world","err":null}
```

```
.
├── README.md
└── hello
    ├── cmd
    │   ├── main.go
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg
        ├── endpoint
        │   ├── endpoint.go
        │   ├── endpoint_gen.go
        │   └── middleware.go
        ├── http
        │   ├── handler.go
        │   └── handler_gen.go
        └── service
            ├── middleware.go
            └── service.go

7 directories, 11 files
```

## Generate the client library
```
go-kit-example-hello> kit g c hello
```

## Create Client
```
go-kit-example-hello> vi hello/client/main.go
```

```
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
```

## Test Client
```
go-kit-example-hello> go run hello/client/main.go
Result: Hello, Client World
```

```
.
├── README.md
└── hello
    ├── client
    │   ├── http
    │   │   └── http.go
    │   └── main.go
    ├── cmd
    │   ├── main.go
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg
        ├── endpoint
        │   ├── endpoint.go
        │   ├── endpoint_gen.go
        │   └── middleware.go
        ├── http
        │   ├── handler.go
        │   └── handler_gen.go
        └── service
            ├── middleware.go
            └── service.go

9 directories, 13 files
```