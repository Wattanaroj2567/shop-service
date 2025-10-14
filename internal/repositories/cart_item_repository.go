package repositories

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"gorm.io/gorm"
)

// CartItemRepository manages individual cart items.
type CartItemRepository interface {
	AddItem(ctx context.Context, item *models.CartItem) error
	UpdateQuantity(ctx context.Context, itemID uint, quantity int) error
	RemoveItem(ctx context.Context, itemID uint) error
}

type cartItemRepository struct {
	db *gorm.DB
}

// NewCartItemRepository constructs CartItemRepository.
func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{db: db}
}

func (r *cartItemRepository) AddItem(ctx context.Context, item *models.CartItem) error {
	// TODO: insert or upsert cart item and recalc totals
	return nil
}

func (r *cartItemRepository) UpdateQuantity(ctx context.Context, itemID uint, quantity int) error {
	// TODO: update cart item quantity then recalc totals
	return nil
}

func (r *cartItemRepository) RemoveItem(ctx context.Context, itemID uint) error {
	// TODO: delete cart item and recalc totals
	return nil
}
