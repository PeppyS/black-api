package main

import (
	"log"
	"os"

	"github.com/PeppyS/black-api/internal/business/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	postgresDSN := os.Getenv("POSTGRES_DSN")

	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
}

func getBusinesses() []model.Business {
	return []model.Business{}
}
