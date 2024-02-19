package item_service

import (
	"context"
	"stock/internal/pkg/domain_model"
	"stock/internal/pkg/storage"
)

// ItemService for all items methods
type ItemService interface {
	ReserveItems(context.Context, *domain_model.ReserveItemsOnStockForDeliveryRequest) error
	DeleteItemsReservation(context.Context, *domain_model.DeleteItemsReserveRequest) error
}

type itemService struct {
	storage storage.Store
}

// NewItemService .
func NewItemService(storage storage.Store) ItemService {
	return &itemService{
		storage: storage,
	}
}
