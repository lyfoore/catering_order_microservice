package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lyfoore/order-service/internal/domain"
	"github.com/lyfoore/order-service/internal/sagas"
	"github.com/lyfoore/order-service/internal/service"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	service *service.OrderService
	saga    *sagas.OrderSaga
}

func NewOrderHandler(service *service.OrderService, saga *sagas.OrderSaga) *OrderHandler {
	return &OrderHandler{
		service: service,
		saga:    saga,
	}
}

func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.service.GetOrderService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, *order)
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var request domain.Order
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(request.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "items are required",
		})
		return
	}

	newOrder := &domain.Order{
		UserID:  request.UserID,
		Items:   request.Items,
		Message: request.Message,
	}

	order, err := h.saga.CreateOrder(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, *order)
}

func (h *OrderHandler) UpdateOrder(ctx *gin.Context) {
	var request domain.Order
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error in request": err.Error(),
		})
		return
	}

	request.ID, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)

	ctx.JSON(http.StatusOK, request)

	updatedOrder := &domain.Order{
		ID:     request.ID,
		UserID: request.UserID,
		Items:  request.Items,
		Status: request.Status,
	}

	ctx.JSON(http.StatusOK, *updatedOrder)

	order, err := h.service.UpdateOrderService(updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, *order)
}

func (h *OrderHandler) DeleteOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.DeleteOrderService(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "order deleted",
	})
}
