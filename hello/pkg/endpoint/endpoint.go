package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/makersu/go-kit-example-hello/hello/pkg/service"
)

// HelloRequest collects the request parameters for the Hello method.
type HelloRequest struct {
	S string `json:"s"`
}

// HelloResponse collects the response parameters for the Hello method.
type HelloResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeHelloEndpoint returns an endpoint that invokes Hello on the service.
func MakeHelloEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		rs, err := s.Hello(ctx, req.S)
		return HelloResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r HelloResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Hello implements Service. Primarily useful in a client.
func (e Endpoints) Hello(ctx context.Context, s string) (rs string, err error) {
	request := HelloRequest{S: s}
	response, err := e.HelloEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HelloResponse).Rs, response.(HelloResponse).Err
}
