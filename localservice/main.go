package main

import (
	"log"

	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/example_server/services/hello"
	"github.com/hysios/mx"
	"github.com/hysios/mx/gateway"
	"github.com/hysios/mx/logger"
	"go.uber.org/zap"
)

func main() {
	gw := gateway.New()
	service, err := mx.NewLocalService("HelloService", &hello.HelloService{}, pb.RegisterHelloServiceHandlerServer)
	if err != nil {
		log.Fatalf("new client service failed: %v", err)
	}
	gw.RegisterService(service)
	gw.Serve(":8080")
}

func init() {
	cfg, _ := zap.NewDevelopment()
	logger.SetLogger(cfg)
}
