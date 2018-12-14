package grpc

import (
	"context"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/makersu/go-kit-example-hello/hello/pkg/endpoint"
	pb "github.com/makersu/go-kit-example-hello/hello/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeHelloHandler creates the handler logic
func makeHelloHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HelloEndpoint, decodeHelloRequest, encodeHelloResponse, options...)
}

// decodeHelloResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	// return nil, errors.New("'pkg Hello' Decoder is not impelemented")
	req := r.(*pb.HelloRequest)
	return endpoint.HelloRequest{S: req.S}, nil
}

// encodeHelloResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	// return nil, errors.New("'pkg Hello' Encoder is not impelemented")
	resp := r.(endpoint.HelloResponse)
	if resp.Err != nil {
		return &pb.HelloReply{Rs: "", Err: resp.Err.Error()}, nil
	}
	return &pb.HelloReply{Rs: resp.Rs, Err: "null"}, nil

}

func (g *grpcServer) Hello(ctx context1.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	_, rep, err := g.hello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloReply), nil
}
