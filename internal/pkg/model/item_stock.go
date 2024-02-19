package model

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// ItemStock db model
type ItemStock struct {
	ID        uuid.UUID `db:"id"`
	StockID   uuid.UUID `db:"stock_id"`
	ItemID    uuid.UUID `db:"item_id"`
	Quantity  null.Int  `db:"quantity"`
	CreatedAt null.Time `db:"created_at"`
	UpdatedAt null.Time `db:"updated_at"`
	DeletedAt null.Time `db:"deleted_at"`
}

// ReservedItem db model
type ReservedItem struct {
	ID        uuid.UUID `db:"id"`
	ItemID    uuid.UUID `db:"item_id"`
	StockID   uuid.UUID `db:"stock_id"`
	Quantity  null.Int  `db:"quantity"`
	CreatedAt null.Time `db:"created_at"`
	UpdatedAt null.Time `db:"updated_at"`
}
