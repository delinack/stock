package stock_service

import (
	"context"
	"storage/internal/pkg/domain_model"
	"storage/internal/pkg/storage"
)

// StockService for all stock methods
type StockService interface {
	GetItemsQuantityOnStock(context.Context, *domain_model.GetItemsQuantityRequest) ([]domain_model.GetItemsQuantityResponse, error)
}

type stockService struct {
	storage storage.Store
}

// NewStockService .
func NewStockService(storage storage.Store) StockService {
	return &stockService{
		storage: storage,
	}
}
