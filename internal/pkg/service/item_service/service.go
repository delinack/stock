package item_service

import (
	"context"

	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/storage"
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
