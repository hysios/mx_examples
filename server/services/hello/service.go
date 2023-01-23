package hello

import (
	"context"

	pb "github.com/hysios/example_server/gen/proto"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (hello *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	return &pb.HelloResponse{
		Message: "hello " + req.Say,
	}, nil
}
