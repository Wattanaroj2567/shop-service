package handlers

import (
	"net/http"

	"github.com/gamegear/shop-service/internal/models"
	"github.com/gamegear/shop-service/internal/services"
	"github.com/gin-gonic/gin"
)

// ProductHandler exposes catalogue endpoints.
type ProductHandler struct {
	productService services.ProductService
}

// NewProductHandler constructs ProductHandler.
func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// List handles GET /api/products.
func (h *ProductHandler) List(c *gin.Context) {
	var filter models.ProductFilter
	_ = filter // TODO: bind query into models.ProductFilter, call service.List, return paginated response
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement list products handler"})
}

// Get handles GET /api/products/:id.
func (h *ProductHandler) Get(c *gin.Context) {
	// TODO: parse path param, call service.GetByID
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement get product handler"})
}
