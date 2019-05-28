package main

import (
	"log"
	"os"

	businessDelivery "github.com/PeppyS/black-api/internal/business/delivery"
	businessService "github.com/PeppyS/black-api/internal/business/service"
	"github.com/PeppyS/black-api/internal/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	grpcAddress := ":" + os.Getenv("GRPC_PORT")
	httpAddress := ":" + os.Getenv("HTTP_PORT")
	postgresDSN := os.Getenv("POSTGRES_DSN")

	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	bs := businessService.NewBusinessService(db)
	bd := businessDelivery.NewBusinessDelivery(bs)

	go server.ListenAndServeHTTPGateway(httpAddress, grpcAddress)
	log.Fatal(server.ListenAndServeGRPCServer(grpcAddress, bd))
}
