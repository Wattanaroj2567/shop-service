package handlers

import (
	"net/http"

	"github.com/gamegear/shop-service/internal/services"
	"github.com/gin-gonic/gin"
)

// OrderHandler exposes order endpoints.
type OrderHandler struct {
	orderService services.OrderService
}

// NewOrderHandler constructs OrderHandler.
func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// Create handles POST /api/orders.
func (h *OrderHandler) Create(c *gin.Context) {
	// TODO: bind request, derive userID, invoke orderService.Create
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement create order handler"})
}

// History handles GET /api/orders/history.
func (h *OrderHandler) History(c *gin.Context) {
	// TODO: derive userID, invoke orderService.History
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement order history handler"})
}

// GetByID handles GET /api/orders/:id.
func (h *OrderHandler) GetByID(c *gin.Context) {
	// TODO: parse orderID, derive userID, invoke orderService.GetByID
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement get order detail handler"})
}
