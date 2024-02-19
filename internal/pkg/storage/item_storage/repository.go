package item_storage

import (
	"context"
	"stock/internal/pkg/model"

	"github.com/jmoiron/sqlx"
)

// ItemRepository for all items methods
type ItemRepository interface {
	ReserveItems(context.Context, []model.ReservedItem, []model.Item) error
	DeleteReservation(context.Context, []model.ReservedItem) error
}

type itemRepository struct {
	db *sqlx.DB
}

// NewItemRepository .
func NewItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{
		db: db,
	}
}
