package main

import (
	"flag"

	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/example_server/services/hello"
	"github.com/hysios/mx/discovery/agent"
	_ "github.com/hysios/mx/discovery/provider/consul"
	"github.com/hysios/mx/service"
)

var (
	addr         = flag.String("addr", ":0", "server address")
	filedescript = flag.Bool("desc", false, "with file descriptor")
)

func main() {
	flag.Parse()

	var srv *service.Server
	if *filedescript {
		srv = service.NewServiceFileDescriptor(&pb.HelloService_ServiceDesc,
			&hello.HelloService{},
			pb.File_proto_hello_proto,
		)
	} else {
		srv = service.NewServiceDesc(&pb.HelloService_ServiceDesc,
			&hello.HelloService{},
		)
	}

	agent.RegisterServer(srv)
	srv.ServeOn(*addr)
}
