package main

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/hysios/mx"
	"github.com/hysios/mx/provisioning"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// provisioning hello service
func init() {
	provisioning.Provision(func(gw *mx.Gateway) {

		// _ = gw.RegisterService("HelloService", mx.WithServiceClient(pb.NewHelloServiceClient, pb.RegisterHelloServiceHandlerClient))
	})
}
