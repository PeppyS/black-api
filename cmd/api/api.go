package main

import (
	"log"
	"os"

	"github.com/PeppyS/black-api/internal"
	businessDelivery "github.com/PeppyS/black-api/internal/business/delivery"
)

func main() {
	grpcAddress := ":" + os.Getenv("GRPC_PORT")
	httpAddress := ":" + os.Getenv("HTTP_PORT")

	bd := businessDelivery.NewBusinessDelivery()

	log.Fatal(internal.ListenAndServe(grpcAddress, httpAddress, bd))
}
