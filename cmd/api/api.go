package main

import (
	"log"
	"os"

	businessDelivery "github.com/PeppyS/black-api/internal/business/delivery"
	"github.com/PeppyS/black-api/internal/server"
)

func main() {
	grpcAddress := ":" + os.Getenv("GRPC_PORT")
	httpAddress := ":" + os.Getenv("HTTP_PORT")

	bd := businessDelivery.NewBusinessDelivery()

	go server.ListenAndServeHTTPGateway(httpAddress, grpcAddress)
	log.Fatal(server.ListenAndServeGRPCServer(grpcAddress, bd))
}
