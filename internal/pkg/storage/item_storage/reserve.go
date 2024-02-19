package item_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/rs/zerolog/log"

	"github.com/delinack/stock/internal/pkg/model"
	"github.com/jmoiron/sqlx"
)

// ReserveItems резервирование товаров на определённых складах
// 1. Создание/обновление записи о резерве определенного количества товара на складе
// 2. Обновление уникальной записи, хранящую связь товара со складом, - уменьшение количества товара, доступного для резервирования
func (r *itemRepository) ReserveItems(ctx context.Context, itemsForReserve []model.ReservedItem, itemsForUpdate []model.Item) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("db.BeginTxx failed")
		return custom_error.ErrWithTransaction
	}
	defer tx.Rollback()

	for _, item := range itemsForReserve {
		err = r.reserveItem(ctx, tx, item)
		if err != nil {
			return fmt.Errorf("cannot reserve item: %w", err)
		}
	}

	for _, item := range itemsForUpdate {
		err = r.reduceItemQuantity(ctx, tx, item)
		if err != nil {
			return fmt.Errorf("cannot reduce item quantity: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("tx.Commit failed")
		return custom_error.ErrWithTransaction
	}

	return nil
}

func (r *itemRepository) reduceItemQuantity(ctx context.Context, tx *sqlx.Tx, item model.Item) error {
	item.Quantity.Int64 *= -1
	q, args, err := buildUpdateItemQuery(item)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_error.ErrNotFound
		} else {
			return fmt.Errorf("tx.ExecContext failed: %w: cannot execute reduce item quantity query", err)
		}
	}

	return nil
}

func (r *itemRepository) reserveItem(ctx context.Context, tx *sqlx.Tx, item model.ReservedItem) error {
	// создаём/обновляем запись о резевре
	q, args, err := buildReserveItemQuery(item)
	if err != nil {
		return fmt.Errorf("buildReserveItemQuery: %w", err)
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_error.ErrNotFound
		} else {
			return fmt.Errorf("tx.ExecContext failed: %w: cannot execute reserve item query", err)
		}
	}

	// уменьшаем количество доступного для резера товара
	q, args, err = buildReduceItemQuantityOnStockQuery(item)
	if err != nil {
		return fmt.Errorf("buildReduceItemQuantityOnStockQuery: %w", err)
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_error.ErrNotFound
		} else {
			return fmt.Errorf("tx.ExecContext failed: %w: cannot execute reduce item quantity on stock query", err)
		}
	}

	return nil
}
