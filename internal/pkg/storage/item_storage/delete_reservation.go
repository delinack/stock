package item_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/rs/zerolog/log"

	"github.com/delinack/stock/internal/pkg/model"
	"github.com/delinack/stock/internal/pkg/service/serializer"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

// DeleteReservation освобождение резерва товаров
// 1. Обновление записи о резерве товара на складе - удаление записи о резервации товара с возвратом количества освобожденого товара N
// 2. Обновление уникальной записи, хранящую связь товара со складом, - увеличение количества товара, доступного для резервирования, на N
// 3. Обновление записи в таблице с товарами - увеличение количества товара, доступного для резервирования, на N
func (r *itemRepository) DeleteReservation(ctx context.Context, items []model.ReservedItem) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("db.BeginTxx failed")
		return custom_error.ErrWithTransaction
	}
	defer tx.Rollback()

	for _, reservedItem := range items {
		quantity, err := r.deleteItemReservation(ctx, tx, reservedItem)
		if err != nil {
			return fmt.Errorf("cannot delete reservedItem: %w", err)
		}

		itemStock := serializer.ToItemStockModel(quantity, reservedItem)
		err = r.addItemQuantityOnStock(ctx, tx, itemStock)
		if err != nil {
			return fmt.Errorf("cannot add reservedItem quantity on stock: %w", err)
		}

		item := serializer.ToItemModel(quantity, reservedItem)
		err = r.addItemQuantity(ctx, tx, item)
		if err != nil {
			return fmt.Errorf("cannot add reservedItem quantity: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("tx.Commit failed")
		return custom_error.ErrWithTransaction
	}

	return nil
}

// удаление записи о резерве товаров на складе из таблицы reserved_items
func (r *itemRepository) deleteItemReservation(ctx context.Context, tx *sqlx.Tx, item model.ReservedItem) (null.Int, error) {
	q, args, err := buildDeleteItemReservationQuery(item)
	if err != nil {
		return null.Int{}, fmt.Errorf("query builder error: %w", err)
	}

	row := tx.QueryRowContext(ctx, q, args...)
	if row.Err() != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return null.Int{}, custom_error.ErrNotFound
		} else {
			return null.Int{}, fmt.Errorf("tx.QueryRowContext failed: %w: cannot execute delete item reservation query", row.Err())
		}
	}

	err = row.Scan(&item.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return null.Int{}, custom_error.ErrNotFound
		} else {
			return null.Int{}, fmt.Errorf("row.Scan failed: %w", err)
		}
	}

	return item.Quantity, nil
}

// увеличение количества свободного для резерва в таблице items
func (r *itemRepository) addItemQuantity(ctx context.Context, tx *sqlx.Tx, item model.Item) error {
	q, args, err := buildUpdateItemQuery(item)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_error.ErrNotFound
		} else {
			return fmt.Errorf("tx.ExecContext failed: %w: cannot execute add item quantity query", err)
		}
	}

	return nil
}

// увеличение количества свободного для резерва в таблице items_stocks
func (r *itemRepository) addItemQuantityOnStock(ctx context.Context, tx *sqlx.Tx, item model.ItemStock) error {
	q, args, err := buildUpdateItemOnStockQuery(item)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_error.ErrNotFound
		} else {
			return fmt.Errorf("tx.ExecContext failed: %w: cannot execute add item quantity on stock query", err)
		}
	}

	return nil
}
