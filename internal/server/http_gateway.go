package server

import (
	"context"
	"net/http"

	"github.com/PeppyS/black-api/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func ListenAndServeHTTPGateway(httpAddress string, grpcAddress string) error {
	ctx := context.Background()

	mux := runtime.NewServeMux(
		// By default proto3 marshals json with omitempty set to true
		// EmitDefaults sets json marshal option to always emit default values
		// OrigName sets option to present same way as defined in proto files
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true, OrigName: true}),
	)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterBusinessServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}

	return server.ListenAndServe()
}
