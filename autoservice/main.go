package main

import (
	_ "github.com/hysios/mx/discovery/provider/consul"
	"github.com/hysios/mx/gateway"
	"github.com/hysios/mx/logger"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/anypb"

	"go.uber.org/zap"
)

func main() {
	gw := gateway.New()
	gw.Serve(":8080")
}

func init() {
	cfg, _ := zap.NewDevelopment()
	logger.SetLogger(cfg)
}
