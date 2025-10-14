package services

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"github.com/gamegear/shop-service/internal/repositories"
)

// ProductService orchestrates catalogue related use cases.
type ProductService interface {
	List(ctx context.Context, filter models.ProductFilter) ([]models.ProductResponse, int64, error)
	GetByID(ctx context.Context, id uint) (*models.ProductResponse, error)
	Create(ctx context.Context, req models.ProductRequest) (*models.ProductResponse, error)
	Update(ctx context.Context, id uint, req models.ProductRequest) (*models.ProductResponse, error)
	Delete(ctx context.Context, id uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

// NewProductService constructs ProductService.
func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) List(ctx context.Context, filter models.ProductFilter) ([]models.ProductResponse, int64, error) {
	// TODO: query repository, map entities to responses, return with total count
	return nil, 0, nil
}

func (s *productService) GetByID(ctx context.Context, id uint) (*models.ProductResponse, error) {
	// TODO: fetch single product and map to response
	return nil, nil
}

func (s *productService) Create(ctx context.Context, req models.ProductRequest) (*models.ProductResponse, error) {
	// TODO: validate payload, persist product, map to response
	return nil, nil
}

func (s *productService) Update(ctx context.Context, id uint, req models.ProductRequest) (*models.ProductResponse, error) {
	// TODO: update product fields and map to response
	return nil, nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	// TODO: delete product (soft delete/hard delete as per requirements)
	return nil
}
