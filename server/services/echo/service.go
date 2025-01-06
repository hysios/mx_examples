package echo

import (
	"context"

	pb "github.com/hysios/example_server/gen/proto"
)

type EchoService struct {
	pb.UnimplementedEchoServiceServer
}

func (echo *EchoService) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoResponse, err error) {
	return &pb.EchoResponse{
		Message: req.Say,
	}, nil
}
