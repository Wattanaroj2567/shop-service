package handlers

import "github.com/gin-gonic/gin"

// RegisterRoutes wires shop-service routes to provided handlers.
func RegisterRoutes(
	router *gin.Engine,
	productHandler *ProductHandler,
	cartHandler *CartHandler,
	orderHandler *OrderHandler,
) {
	api := router.Group("/api")

	api.GET("/products", productHandler.List)
	api.GET("/products/:id", productHandler.Get)

	api.GET("/cart", cartHandler.GetCart)
	api.POST("/cart/add", cartHandler.AddItem)
	api.PUT("/cart/update", cartHandler.UpdateItem)
	api.DELETE("/cart/remove", cartHandler.RemoveItem)

	api.POST("/orders", orderHandler.Create)
	api.GET("/orders/history", orderHandler.History)
	api.GET("/orders/:id", orderHandler.GetByID)
}
