package main

import (
	"log"

	"github.com/MHS-20/Order-Microservice/order/config"
	"github.com/MHS-20/Order-Microservice/order/internal/adapters/db"
	"github.com/MHS-20/Order-Microservice/order/internal/adapters/grpc"
	"github.com/MHS-20/Order-Microservice/order/internal/adapters/payment"
	"github.com/MHS-20/Order-Microservice/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
