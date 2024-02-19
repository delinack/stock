package serializer

import (
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/model"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// ToReserveItemsModelFromReserveRequest .
func ToReserveItemsModelFromReserveRequest(params domain_model.ReserveItemsOnStockForDeliveryRequest) []model.ReservedItem {
	reservedItems := make([]model.ReservedItem, 0)

	for _, item := range params.Items {
		reservedItems = append(reservedItems, toReserveItemModelFromReserveRequest(params.StockID, item))
	}

	return reservedItems
}

func toReserveItemModelFromReserveRequest(stockID uuid.UUID, item domain_model.ReserveItem) model.ReservedItem {
	return model.ReservedItem{
		StockID:  stockID,
		ItemID:   item.ItemID,
		Quantity: null.IntFrom(item.Quantity),
	}
}

// ToItemsModelFromReserveRequest .
func ToItemsModelFromReserveRequest(params domain_model.ReserveItemsOnStockForDeliveryRequest) []model.Item {
	items := make([]model.Item, 0)

	for _, item := range params.Items {
		items = append(items, toItemModelFromReserveRequest(item))
	}

	return items
}

func toItemModelFromReserveRequest(item domain_model.ReserveItem) model.Item {
	return model.Item{
		ID:       item.ItemID,
		Quantity: null.IntFrom(item.Quantity),
	}
}

// ToReservedItemsModelFromDeleteRequest .
func ToReservedItemsModelFromDeleteRequest(params domain_model.DeleteItemsReserveRequest) []model.ReservedItem {
	reservedItems := make([]model.ReservedItem, 0)

	for _, item := range params.Items {
		reservedItems = append(reservedItems, toReservedItemModelFromDeleteRequest(params.StockID, item))
	}

	return reservedItems
}

func toReservedItemModelFromDeleteRequest(stockID uuid.UUID, item domain_model.DeleteItem) model.ReservedItem {
	return model.ReservedItem{
		StockID: stockID,
		ItemID:  item.ItemID,
	}
}

// ToItemStockModel .
func ToItemStockModel(quantity null.Int, items model.ReservedItem) model.ItemStock {
	return model.ItemStock{
		ItemID:   items.ItemID,
		StockID:  items.StockID,
		Quantity: quantity,
	}
}

// ToItemModel .
func ToItemModel(quantity null.Int, item model.ReservedItem) model.Item {
	return model.Item{
		ID:       item.ItemID,
		Quantity: quantity,
	}
}
