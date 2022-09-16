package main

import (
	"database/sql"
	"os"

	"chadgh.com/bank/repository"
	_ "github.com/lib/pq"
)

var service *server

func init() {
	databaseUrl := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	r := repository.NewTransactionRepo(db)

	service = &server{
		repo:    r,
		verbose: false,
	}
}

func LambdaHandler() {
	service.Run()
}
