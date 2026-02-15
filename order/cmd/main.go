package main

import (
	"log"

	"github.com/sssseraphim/microservices/order/config"
	"github.com/sssseraphim/microservices/order/internal/adapters/db"
	"github.com/sssseraphim/microservices/order/internal/adapters/grpc"
	"github.com/sssseraphim/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to create db adapter: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
