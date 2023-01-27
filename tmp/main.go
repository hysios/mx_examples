package main

import (
	"github.com/hysios/mx/service"
	pb "play.test/gen/proto"
	"play.test/services/hello"
)

func main() {
	// Services Register
	srv := service.NewServiceFileDescriptor(&pb.HelloService_ServiceDesc,
		&hello.HelloService{},
		pb.File_proto_hello_proto,
	)
	srv.Start()
}
