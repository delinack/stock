package storage

import (
	"github.com/delinack/stock/internal/pkg/storage/item_storage"
	"github.com/delinack/stock/internal/pkg/storage/stock_storage"
)

// Store for all stock interfaces.
type Store interface {
	stock_storage.StockRepository
	item_storage.ItemRepository
}

type store struct {
	stock_storage.StockRepository
	item_storage.ItemRepository
}

// NewStorage .
func NewStorage(stock stock_storage.StockRepository, item item_storage.ItemRepository) Store {
	return &store{
		stock,
		item,
	}
}
