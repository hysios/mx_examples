package hello

import (
	"context"

	pb "play.test/gen/proto"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) GetUser(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {

	// TODO: implement
	return &pb.UserResponse{
		Name: "hello",
	}, nil

}
