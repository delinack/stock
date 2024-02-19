package stock_storage

import (
	"context"

	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// StockRepository for all stock methods
type StockRepository interface {
	GetItemsQuantity(context.Context, *domain_model.GetItemsQuantityRequest) ([]model.ItemStock, error)
	GetItemQuantity(context.Context, uuid.UUID, domain_model.ReserveItem) (int64, error)
	CheckAvailability(context.Context, uuid.UUID) (bool, error)
}

// stockRepository
type stockRepository struct {
	db *sqlx.DB
}

// NewStockRepository .
func NewStockRepository(db *sqlx.DB) StockRepository {
	return &stockRepository{
		db: db,
	}
}
