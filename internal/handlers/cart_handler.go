package handlers

import (
	"net/http"

	"github.com/gamegear/shop-service/internal/services"
	"github.com/gin-gonic/gin"
)

// CartHandler exposes member cart endpoints.
type CartHandler struct {
	cartService services.CartService
}

// NewCartHandler constructs CartHandler.
func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// GetActive handles GET /api/cart.
func (h *CartHandler) GetActive(c *gin.Context) {
	// TODO: derive userID from auth context, invoke cartService.GetActiveCart
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement get cart handler"})
}

// AddItem handles POST /api/cart/add.
func (h *CartHandler) AddItem(c *gin.Context) {
	// TODO: bind request, call cartService.AddItem
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement add to cart handler"})
}

// UpdateItem handles PUT /api/cart/update.
func (h *CartHandler) UpdateItem(c *gin.Context) {
	// TODO: bind request, call cartService.UpdateItem
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement update cart handler"})
}

// RemoveItem handles DELETE /api/cart/remove.
func (h *CartHandler) RemoveItem(c *gin.Context) {
	// TODO: bind request, call cartService.RemoveItem
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement remove cart handler"})
}
