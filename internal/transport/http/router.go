package http

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	engine  *gin.Engine
	handler *OrderHandler
}

func NewRouter(handler *OrderHandler) *Router {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return &Router{
		engine:  engine,
		handler: handler,
	}
}

func (r *Router) SetupRouter() {
	r.engine.GET("/orders/:id", r.handler.GetOrder)
	r.engine.POST("/orders", r.handler.CreateOrder)
	r.engine.PATCH("/orders/:id", r.handler.UpdateOrder)
	r.engine.DELETE("/orders/:id", r.handler.DeleteOrder)

	pprof.Register(r.engine)
}

func (r *Router) Run(port string) {
	log.Println("HTTP server is starting on port", port)
	r.engine.Run(port)
}
