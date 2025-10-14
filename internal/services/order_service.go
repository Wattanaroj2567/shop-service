package services

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"github.com/gamegear/shop-service/internal/repositories"
)

// OrderService coordinates checkout lifecycle.
type OrderService interface {
	Create(ctx context.Context, userID uint, req models.CreateOrderRequest) (*models.OrderDetail, error)
	History(ctx context.Context, userID uint) ([]models.OrderSummary, error)
	GetByID(ctx context.Context, userID uint, orderID uint) (*models.OrderDetail, error)
}

type orderService struct {
	orderRepo repositories.OrderRepository
	cartRepo  repositories.CartRepository
}

// NewOrderService constructs OrderService.
func NewOrderService(orderRepo repositories.OrderRepository, cartRepo repositories.CartRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (s *orderService) Create(ctx context.Context, userID uint, req models.CreateOrderRequest) (*models.OrderDetail, error) {
	// TODO: load cart, calculate totals, persist order & items, clear cart
	return nil, nil
}

func (s *orderService) History(ctx context.Context, userID uint) ([]models.OrderSummary, error) {
	// TODO: fetch order list and map to summaries
	return nil, nil
}

func (s *orderService) GetByID(ctx context.Context, userID uint, orderID uint) (*models.OrderDetail, error) {
	// TODO: fetch single order with items ensuring it belongs to user
	return nil, nil
}
