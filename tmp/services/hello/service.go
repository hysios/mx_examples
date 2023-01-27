package hello

import (
	"context"

	pb "play.test/gen/proto"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	// TODO: implement
	return
}
