package serializer

import (
	"storage/internal/pkg/domain_model"
	"storage/internal/pkg/model"
)

// ToGetItemsQuantityResponse .
func ToGetItemsQuantityResponse(items []model.ItemStock) []domain_model.GetItemsQuantityResponse {
	response := make([]domain_model.GetItemsQuantityResponse, 0)

	for _, item := range items {
		response = append(response, toGetItemsQuantityResponse(item))
	}

	return response
}

func toGetItemsQuantityResponse(item model.ItemStock) domain_model.GetItemsQuantityResponse {
	return domain_model.GetItemsQuantityResponse{
		ItemID:   item.ItemID,
		Quantity: item.Quantity.Int64,
	}
}
