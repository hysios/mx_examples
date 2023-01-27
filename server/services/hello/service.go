package hello

import (
	"context"

	"github.com/hysios/example_server/common"
	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/mx/discovery/agent"
	_ "github.com/hysios/mx/discovery/provider/consul"
	"gorm.io/gorm"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer

	db *gorm.DB
}

func (hello *HelloService) Init() error {
	cfg, err := agent.Config(nil)
	if err != nil {
		return err
	}

	hello.db, err = common.OpenDatabase(cfg)
	return err
}

func (hello *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	return &pb.HelloResponse{
		Message: "hello " + req.Say,
	}, nil
}
