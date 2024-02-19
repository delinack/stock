package item_service

import (
	"context"
	"fmt"
	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/service/serializer"
	"github.com/rs/zerolog/log"
)

// ReserveItems резервирование товара на складе для доставки
// 1. Проверка доступности склада
// 2. Проверка доступности товаров на складе (запрашиваемое для резерва количество товара не должно превышать доступное и не должно быть нулем)
// 3. Резервирование товаров
func (s *itemService) ReserveItems(ctx context.Context, params *domain_model.ReserveItemsOnStockForDeliveryRequest) error {
	isAvailable, err := s.storage.CheckAvailability(ctx, params.StockID)
	if err != nil {
		return fmt.Errorf("storage.CheckAvailability failed: %w", err)
	}
	if !isAvailable {
		return custom_error.ErrUnavailableStock
	}

	for _, item := range params.Items {
		availableQuantity, err := s.storage.GetItemQuantity(ctx, params.StockID, item)
		if err != nil {
			return fmt.Errorf("storage.GetItemsQuantity failed: %w", err)
		}
		if availableQuantity < item.Quantity {
			log.Error().Msgf("item - %s: available quantity %d < request quantity %d", item.ItemID, availableQuantity, item.Quantity)
			return custom_error.ErrExceededValue
		}
		if item.Quantity == 0 {
			log.Error().Msg("quantity for reserve cannot be null")
			return custom_error.ErrNullValue
		}
	}

	reserveItemModel := serializer.ToReserveItemsModelFromReserveRequest(*params)
	itemModel := serializer.ToItemsModelFromReserveRequest(*params)

	err = s.storage.ReserveItems(ctx, reserveItemModel, itemModel)
	if err != nil {
		return fmt.Errorf("storage.ReserveItems failed: %w", err)
	}

	return nil
}

// DeleteItemsReservation освобождение резерва товаров
// 1. Снятие резерва с товаров
func (s *itemService) DeleteItemsReservation(ctx context.Context, params *domain_model.DeleteItemsReserveRequest) error {
	reservedItemModel := serializer.ToReservedItemsModelFromDeleteRequest(*params)

	err := s.storage.DeleteReservation(ctx, reservedItemModel)
	if err != nil {
		return fmt.Errorf("storage.DeleteReservation failed: %w", err)
	}

	return nil
}
