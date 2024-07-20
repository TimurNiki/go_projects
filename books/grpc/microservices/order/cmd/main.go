package main

import (
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/adapters/db"
	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := gprc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
