package hello

import (
	"context"

	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/mx/client"
	_ "github.com/hysios/mx/discovery/provider/consul"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer

	echo func() pb.EchoServiceClient
	db   *gorm.DB
}

func (hello *HelloService) Init() error {
	// cfg, err := agent.Config(nil)
	// if err != nil {
	// 	return err
	// }

	// if err = client.LMake("echo.EchoService", &hello.echo); err != nil {
	// 	return err
	// }

	// hello.db, err = common.OpenDatabase(cfg)
	// return err
	return nil
}

func (hello *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	if hello.echo != nil {
		echoresp, err := hello.echo().Echo(ctx, &pb.EchoRequest{Say: req.Say})
		if err != nil {
			return nil, status.Errorf(status.Code(err), "echo failed: %v", err)
		}
		return &pb.HelloResponse{
			Message: "hello " + echoresp.Message,
		}, nil

	} else {
		return &pb.HelloResponse{
			Message: "hello " + req.Say,
		}, nil
	}
}

func init() {
	client.Registry("echo.EchoService", pb.NewEchoServiceClient)
}
