package repositories

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"gorm.io/gorm"
)

// OrderRepository persists orders and order items.
type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	FindHistoryByUserID(ctx context.Context, userID uint) ([]models.Order, error)
	FindByID(ctx context.Context, orderID uint) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository constructs OrderRepository.
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(ctx context.Context, order *models.Order) error {
	// TODO: persist order with items inside transaction
	return nil
}

func (r *orderRepository) FindHistoryByUserID(ctx context.Context, userID uint) ([]models.Order, error) {
	// TODO: return orders filtered by user
	return nil, nil
}

func (r *orderRepository) FindByID(ctx context.Context, orderID uint) (*models.Order, error) {
	// TODO: return order with nested items by ID
	return nil, nil
}
