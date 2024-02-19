package stock_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/model"
	"github.com/google/uuid"
)

// GetItemsQuantity возвращает количество всех товаров на складе.
func (r *stockRepository) GetItemsQuantity(ctx context.Context, params *domain_model.GetItemsQuantityRequest) ([]model.ItemStock, error) {
	var quantity []model.ItemStock

	q, args, err := buildGetItemsQuantityQuery(params)
	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &quantity, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, custom_error.ErrNotFound
		} else {
			return nil, fmt.Errorf("db.SelectContext failed: %w: get items quantity query", err)
		}
	}

	return quantity, nil
}

// GetItemQuantity возвращает количество определенного товара на складе. Проверки доступности товара для резервации.
func (r *stockRepository) GetItemQuantity(ctx context.Context, stockID uuid.UUID, item domain_model.ReserveItem) (int64, error) {
	var quantity int64

	q, args, err := buildGetItemQuantityQuery(stockID, item)
	if err != nil {
		return 0, err
	}

	err = r.db.GetContext(ctx, &quantity, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, custom_error.ErrNotFound
		} else {
			return 0, fmt.Errorf("db.GetContext failed: %w: get item quantity query", err)
		}
	}

	return quantity, nil
}

// CheckAvailability проверка доступности склада.
func (r *stockRepository) CheckAvailability(ctx context.Context, stockID uuid.UUID) (bool, error) {
	var isAvailable bool

	q, args, err := buildCheckStockAvailability(stockID)
	if err != nil {
		return false, err
	}

	err = r.db.GetContext(ctx, &isAvailable, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, custom_error.ErrNotFound
		} else {
			return false, fmt.Errorf("db.GetContext failed: %w: get availability query", err)
		}
	}

	return isAvailable, nil
}
