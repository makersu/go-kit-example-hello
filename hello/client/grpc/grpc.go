package grpc

import (
	"context"
	"errors"

	endpoint "github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	pkgendpoint "github.com/makersu/go-kit-example-hello/hello/pkg/endpoint"
	pb "github.com/makersu/go-kit-example-hello/hello/pkg/grpc/pb"
	service "github.com/makersu/go-kit-example-hello/hello/pkg/service"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]kitgrpc.ClientOption) (service.HelloService, error) {
	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = kitgrpc.NewClient(conn, "pb.Hello", "Hello", encodeHelloRequest, decodeHelloResponse, pb.HelloReply{}, options["Hello"]...).Endpoint()
	}

	return pkgendpoint.Endpoints{HelloEndpoint: helloEndpoint}, nil
}

// encodeHelloRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain sum request to a gRPC request.
func encodeHelloRequest(_ context.Context, request interface{}) (interface{}, error) {
	// return nil, errors.New("'client Hello' Encoder is not impelemented")
	req := request.(pkgendpoint.HelloRequest)
	return &pb.HelloRequest{S: req.S}, nil
}

// decodeHelloResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeHelloResponse(_ context.Context, reply interface{}) (interface{}, error) {
	// return nil, errors.New("'client Hello' Decoder is not impelemented")
	hrep := reply.(*pb.HelloReply)

	if hrep.Err != "null" {
		return pkgendpoint.HelloResponse{Rs: "null", Err: errors.New(hrep.Err)}, nil
	}
	// return endpoint1.HelloResponse{Rs: hrep.Err, Err: nil}, nil
	return pkgendpoint.HelloResponse{Rs: hrep.Rs, Err: nil}, nil
}
