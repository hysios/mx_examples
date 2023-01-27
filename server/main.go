package main

import (
	"flag"

	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/example_server/services/echo"
	"github.com/hysios/example_server/services/hello"
	"github.com/hysios/mx/discovery/agent"
	_ "github.com/hysios/mx/discovery/provider/consul"
	"github.com/hysios/mx/logger"
	"github.com/hysios/mx/server"
	"go.uber.org/zap"
)

var (
	addr         = flag.String("addr", ":0", "server address")
	filedescript = flag.Bool("desc", false, "with file descriptor")
)

func main() {
	flag.Parse()

	srv := server.New("ServiceCluster")
	if *filedescript {
		srv.RegisterService(
			&pb.HelloService_ServiceDesc,
			&hello.HelloService{},
			server.WithFileDescriptor(pb.File_proto_hello_proto),
		)

		srv.RegisterService(
			&pb.EchoService_ServiceDesc,
			&echo.EchoService{},
			server.WithFileDescriptor(pb.File_proto_echo_proto),
		)
	} else {
		srv.RegisterService(
			&pb.HelloService_ServiceDesc,
			&hello.HelloService{},
		)

		srv.RegisterService(
			&pb.EchoService_ServiceDesc,
			&echo.EchoService{},
		)
	}

	if err := agent.RegisterServer(srv); err != nil {
		logger.Logger.Error("register server", zap.Error(err))
	}

	srv.ServeOn(*addr)
}

func init() {
	cfg, _ := zap.NewDevelopment()
	logger.SetLogger(cfg)
}
