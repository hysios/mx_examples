package main

import (
	"github.com/hysios/mx/service"
	pb "play.test/gen/proto"
	"play.test/services/hello"
)

func main() {
	// Services Register
	srv := service.New("HelloService", &hello.HelloService{}, service.WithServiceDesc(&pb.HelloService_ServiceDesc))
	srv.Start()
}
