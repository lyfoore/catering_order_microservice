package main

import (
	database "github.com/lyfoore/order-service/internal/db"
	"github.com/lyfoore/order-service/internal/service"
	"github.com/lyfoore/order-service/internal/transport/gRPC"
	"github.com/lyfoore/order-service/internal/transport/http"
	"log"
)

func main() {
	//mockDB := mock.NewMockDB()
	postgresDB, err := database.NewPostgresDB("host=db port=5432 user=postgres password=postgres dbname=orders sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	orderService := service.NewOrderService(postgresDB)
	orderHandler := http.NewOrderHandler(orderService)

	grpcServer := gRPC.NewOrderServer(orderService)
	go grpcServer.RunServer()

	engine := http.NewRouter(orderHandler)
	engine.SetupRouter()
	engine.Run(":8080")
}
