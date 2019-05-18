package internal

import (
	"context"
	"net"
	"net/http"

	"github.com/PeppyS/black-api/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func ListenAndServe(grpcAddress string, httpAddress string, businessService proto.BusinessServiceServer) error {
	ctx := context.Background()

	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	// Setup GRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterBusinessServiceServer(grpcServer, businessService)

	// Setup HTTP Gateway proxy
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err = proto.RegisterBusinessServiceHandlerFromEndpoint(ctx, mux, grpcAddress, dialOpts)
	if err != nil {
		return err
	}

	httpGateway := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}
	go httpGateway.ListenAndServe()

	return grpcServer.Serve(listen)
}
