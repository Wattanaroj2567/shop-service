package repositories

import (
	"context"

	"github.com/gamegear/shop-service/internal/models"
	"gorm.io/gorm"
)

// CartRepository manages cart aggregate.
type CartRepository interface {
	FindActiveByUserID(ctx context.Context, userID uint) (*models.Cart, error)
	Create(ctx context.Context, cart *models.Cart) error
	UpdateTotals(ctx context.Context, cart *models.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

// NewCartRepository constructs CartRepository.
func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) FindActiveByUserID(ctx context.Context, userID uint) (*models.Cart, error) {
	// TODO: implement find cart by userID with eager-loaded items
	return nil, nil
}

func (r *cartRepository) Create(ctx context.Context, cart *models.Cart) error {
	// TODO: implement create cart
	return nil
}

func (r *cartRepository) UpdateTotals(ctx context.Context, cart *models.Cart) error {
	// TODO: implement update cart totals after mutation
	return nil
}
