package main

import (
	"context"
	"log"

	pb "github.com/hysios/example_server/gen/proto"
	"github.com/hysios/mx"

	"github.com/hysios/mx/gateway"
	"github.com/hysios/mx/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	var ctx = context.Background()
	conn, err := grpc.DialContext(ctx, ":8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	gw := gateway.New()
	service, err := mx.NewClientService("HelloService", conn, pb.RegisterHelloServiceHandler)
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
