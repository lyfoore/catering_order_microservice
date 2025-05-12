package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyfoore/order-service/internal/genproto/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func GetOrderGRPC(ctx *gin.Context) {
	conn, _ := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := orders.NewOrderServiceClient(conn)
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	order, err := client.GetOrder(ctx, &orders.GetOrderRequest{Id: id})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, order)
}

func main() {
	restServer := gin.Default()
	restServer.GET("/order/:id", GetOrderGRPC)
	restServer.Run(":8081")
}
