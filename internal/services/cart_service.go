package services

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"github.com/gamegear/shop-service/internal/repositories"
)

// CartService manages cart operations for members.
type CartService interface {
	GetActiveCart(ctx context.Context, userID uint) (*models.CartResponse, error)
	AddItem(ctx context.Context, userID uint, req models.AddToCartRequest) (*models.CartResponse, error)
	UpdateItem(ctx context.Context, userID uint, req models.UpdateCartItemRequest) (*models.CartResponse, error)
	RemoveItem(ctx context.Context, userID uint, req models.RemoveCartItemRequest) (*models.CartResponse, error)
}

type cartService struct {
	cartRepo     repositories.CartRepository
	cartItemRepo repositories.CartItemRepository
	productRepo  repositories.ProductRepository
}

// NewCartService constructs CartService.
func NewCartService(
	cartRepo repositories.CartRepository,
	cartItemRepo repositories.CartItemRepository,
	productRepo repositories.ProductRepository,
) CartService {
	return &cartService{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo,
		productRepo:  productRepo,
	}
}

func (s *cartService) GetActiveCart(ctx context.Context, userID uint) (*models.CartResponse, error) {
	// TODO: load active cart and map to response structure
	return nil, nil
}

func (s *cartService) AddItem(ctx context.Context, userID uint, req models.AddToCartRequest) (*models.CartResponse, error) {
	// TODO: ensure cart exists, validate stock, add item, recalc totals
	return nil, nil
}

func (s *cartService) UpdateItem(ctx context.Context, userID uint, req models.UpdateCartItemRequest) (*models.CartResponse, error) {
	// TODO: adjust quantity and recalc totals
	return nil, nil
}

func (s *cartService) RemoveItem(ctx context.Context, userID uint, req models.RemoveCartItemRequest) (*models.CartResponse, error) {
	// TODO: remove item and re-evaluate cart totals
	return nil, nil
}
