package service

import (
	"storage/internal/pkg/service/item_service"
	"storage/internal/pkg/service/stock_service"
	"storage/internal/pkg/storage"
)

// Service for all service interfaces
type Service interface {
	stock_service.StockService
	item_service.ItemService
}

type service struct {
	stock_service.StockService
	item_service.ItemService
}

// NewService constructor for all services
func NewService(store storage.Store) Service {
	return &service{
		stock_service.NewStockService(store),
		item_service.NewItemService(store),
	}
}
