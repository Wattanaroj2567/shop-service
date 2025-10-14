package repositories

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"gorm.io/gorm"
)

// ProductRepository handles persistence for catalogue data.
type ProductRepository interface {
	List(ctx context.Context, filter models.ProductFilter) ([]models.Product, error)
	Count(ctx context.Context, filter models.ProductFilter) (int64, error)
	FindByID(ctx context.Context, id uint) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id uint) error
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository wires a GORM-backed repository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) List(ctx context.Context, filter models.ProductFilter) ([]models.Product, error) {
	// TODO: implement catalogue query filtered by pagination/category/search
	return nil, nil
}

func (r *productRepository) Count(ctx context.Context, filter models.ProductFilter) (int64, error) {
	// TODO: implement count for pagination metadata
	return 0, nil
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*models.Product, error) {
	// TODO: implement find by ID
	return nil, nil
}

func (r *productRepository) Create(ctx context.Context, product *models.Product) error {
	// TODO: implement create product
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *models.Product) error {
	// TODO: implement update product
	return nil
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	// TODO: implement delete product
	return nil
}
