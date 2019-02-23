# go-kit-example-hello-world

## Install kit
```
go-kit-example-hello> go get github.com/go-kit/kit
go-kit-example-hello> go get github.com/kujtimiihoxha/kit
```

## Create hello service
```
# kit new service hello
go-kit-example-hello> kit n s hello
```

## Define hello service (hello/pkg/service/service.go)
```
type HelloService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Hello(ctx context.Context, s string) (rs string, err error)
}
```
## Initiate hello service( with default middleware)
```
# kit generate service hello --dmw
go-kit-example-hello> kit g s hello -w
```

## Implement hello service (hello/pkg/service/service.go)

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

go-kit-example-hello> curl -XPOST -d'{"s":""}' localhost:8081/hello
{"err":"empty string"}

go-kit-example-hello> curl localhost:8080/metrics
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

## Generate the HTTP client library
```
go-kit-example-hello> kit g c hello
```

## Create Http Client (hello/client/httpmain.go)

```
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
```

## Test Http Client
```
go-kit-example-hello> go run hello/client/httpmain.go
HTTP Client Result: Hello, Http Client
HTTP Client Error: empty string
```

```
.
├── README.md
└── hello
    ├── client
    │   ├── http
    │   │   └── http.go
    │   └── httpmain.go
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

## Install GRPC
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## Generate GRPC transport
```
kit g s hello --dmw -t grpc
```

## Define proto
```
go-kit-example-hello> vi hello/pkg/grpc/pb/hello.proto
```
```
syntax = "proto3";

package pb;


//The Hello service definition.
service Hello {
 rpc Hello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
 string s = 1;
}

message HelloReply {
 string rs  = 1;
 string err = 2;
}
```

## Generate GRPC server and client stubs(./hello/pkg/grpc/pb/hello.pb.go)
```
go-kit-example-hello> cd hello/pkg/grpc/pb
go-kit-example-hello/hello/pkg/grpc/pb> ./compile.sh
```

## Implement decoder and encoder for GRPC transport(./hello/pkg/grpc/handler.go)
```
func decodeHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HelloRequest)
	return endpoint.HelloRequest{S: req.S}, nil
}

func encodeHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.HelloResponse)
	if resp.Err != nil {
		return &pb.HelloReply{Rs: "", Err: resp.Err.Error()}, nil
	}
	return &pb.HelloReply{Rs: resp.Rs, Err: "null"}, nil
}
```

## Generate the GRPC client library
```
go-kit-example-hello> kit g c hello -t grpc
```

## Implement the GRPC client library
```
func encodeHelloRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(pkgendpoint.HelloRequest)
	return &pb.HelloRequest{S: req.S}, nil
}

func decodeHelloResponse(_ context.Context, reply interface{}) (interface{}, error) {
	hrep := reply.(*pb.HelloReply)

	if hrep.Err != "null" {
		return endpoint1.HelloResponse{Rs: "null", Err: errors.New(hrep.Err)}, nil
	}

	return endpoint1.HelloResponse{Rs: hrep.Rs, Err: nil}, nil
}
```

## Create GRPC Client
```
go-kit-example-hello> vi hello/client/grpcmain.go
```

```
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

```

## Test GRPC Client
```
go-kit-example-hello> go run hello/client/grpcmain.go
GRPC Client Result: Hello, GRPC Client
GRPC Client Error: empty string
```

```
.
├── README.md
└── hello
    ├── client
    │   ├── grpc
    │   │   └── grpc.go
    │   ├── grpcmain.go
    │   ├── http
    │   │   └── http.go
    │   ├── httpmain.go
    │   └── main.go
    ├── cmd
    │   ├── main.go
    │   └── service
    │       ├── service.go
    │       └── service_gen.go
    └── pkg
        ├── endpoint
        │   ├── endpoint.go
        │   ├── endpoint_gen.go
        │   └── middleware.go
        ├── grpc
        │   ├── handler.go
        │   ├── handler_gen.go
        │   └── pb
        │       ├── compile.sh
        │       ├── hello.pb.go
        │       └── hello.proto
        ├── http
        │   ├── handler.go
        │   └── handler_gen.go
        └── service
            ├── middleware.go
            └── service.go

12 directories, 21 files
```
```
wc -l $(find ./ -name "*.go")
     169 .//cmd/service/service.go
      43 .//cmd/service/service_gen.go
       7 .//cmd/main.go
      40 .//client/grpcmain.go
      47 .//client/grpc/grpc.go
      67 .//client/http/http.go
      32 .//client/httpmain.go
      45 .//pkg/grpc/handler.go
     201 .//pkg/grpc/pb/hello.pb.go
      17 .//pkg/grpc/handler_gen.go
      57 .//pkg/http/handler.go
      16 .//pkg/http/handler_gen.go
      24 .//pkg/endpoint/endpoint_gen.go
      53 .//pkg/endpoint/endpoint.go
      39 .//pkg/endpoint/middleware.go
      37 .//pkg/service/service.go
      31 .//pkg/service/middleware.go
     925 total
```

## TODO
- docker integration
- service discovery
- distributed tracing
- ...
