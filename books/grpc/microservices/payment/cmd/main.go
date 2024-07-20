package main

import (
	"log"

	"github.com/TimurNiki/go_api_tutorial/books/grpc/microservices/payment/internal/adapters/payment"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err :=
		payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application,
		config.GetApplicationPort())
	grpcAdapter.Run()
}
