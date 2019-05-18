package server

import (
	"net"

	"github.com/PeppyS/black-api/internal/middleware"
	"github.com/PeppyS/black-api/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func ListenAndServeGRPCServer(address string, businessService proto.BusinessServiceServer) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc.UnaryServerInterceptor(middleware.Logging),
		)),
	)
	proto.RegisterBusinessServiceServer(server, businessService)

	return server.Serve(listen)
}
