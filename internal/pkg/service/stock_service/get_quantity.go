package stock_service

import (
	"context"
	"fmt"
	"storage/internal/pkg/domain_model"
	"storage/internal/pkg/service/serializer"

	"github.com/rs/zerolog/log"
)

// GetItemsQuantityOnStock возвращает количество оставшихся товаров на складе.
func (s *stockService) GetItemsQuantityOnStock(ctx context.Context, params *domain_model.GetItemsQuantityRequest) ([]domain_model.GetItemsQuantityResponse, error) {
	isAvailable, err := s.storage.CheckAvailability(ctx, params.StockID)
	if err != nil {
		return nil, fmt.Errorf("storage.CheckAvailability failed: %w", err)
	}
	if !isAvailable {
		return nil, fmt.Errorf("stock is unavailable")
	}

	quantityModel, err := s.storage.GetItemsQuantity(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("storage.GetItemsQuantity failed: %w", err)
	}

	quantity := serializer.ToGetItemsQuantityResponse(quantityModel)

	log.Debug().Str("comp", "stock service: get items quantity").Send()

	return quantity, nil
}
