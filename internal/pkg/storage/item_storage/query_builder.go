package item_storage

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/delinack/stock/internal/pkg/model"
)

const (
	nowSQL = "now()"
)

func buildReserveItemQuery(item model.ReservedItem) (string, []interface{}, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("reserved_items").
		Columns("item_id", "stock_id", "quantity", "created_at").
		Values(item.ItemID, item.StockID, item.Quantity.Int64, nowSQL).
		Suffix("ON CONFLICT (stock_id, item_id) DO UPDATE SET quantity = reserved_items.quantity + EXCLUDED.quantity, updated_at = now()")

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql() failed: %w", err)
	}

	return q, args, nil
}

func buildReduceItemQuantityOnStockQuery(item model.ReservedItem) (string, []interface{}, error) {
	updateItemMap := make(map[string]interface{})

	updateItemMap["quantity"] = sq.Expr("quantity - ?", item.Quantity.Int64)
	updateItemMap["updated_at"] = nowSQL

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("items_stocks").
		SetMap(updateItemMap).
		Where(sq.Eq{"stock_id": item.StockID}).
		Where(sq.Eq{"item_id": item.ItemID})

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql() failed: %w", err)
	}

	return q, args, nil
}

func buildUpdateItemQuery(item model.Item) (string, []interface{}, error) {
	updateItemMap := make(map[string]interface{})

	if item.Name.Valid {
		updateItemMap["name"] = item.Name.String
	}
	if item.Size.Valid {
		updateItemMap["size"] = item.Size.String
	}
	if item.Quantity.Valid {
		updateItemMap["quantity"] = sq.Expr("quantity + ?", item.Quantity.Int64)
	}
	if item.UniqueCode.Valid {
		updateItemMap["unique_code"] = item.UniqueCode.String
	}

	updateItemMap["updated_at"] = nowSQL

	q, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("items").
		SetMap(updateItemMap).
		Where(sq.Eq{"id": item.ID}).
		ToSql()

	if err != nil {
		return "", nil, fmt.Errorf("query builder error - ToSql(): %w", err)
	}

	return q, args, nil
}

func buildUpdateItemOnStockQuery(item model.ItemStock) (string, []interface{}, error) {
	updateItemMap := make(map[string]interface{})

	if item.Quantity.Valid {
		updateItemMap["quantity"] = sq.Expr("quantity + ?", item.Quantity.Int64)
	}

	updateItemMap["updated_at"] = nowSQL

	q, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("items_stocks").
		SetMap(updateItemMap).
		Where(sq.Eq{"stock_id": item.StockID}).
		Where(sq.Eq{"item_id": item.ItemID}).
		ToSql()

	if err != nil {
		return "", nil, fmt.Errorf("ToSql(): %w", err)
	}

	return q, args, nil
}

func buildDeleteItemReservationQuery(item model.ReservedItem) (string, []interface{}, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete("reserved_items").
		Where(sq.Eq{"item_id": item.ItemID, "stock_id": item.StockID}).
		Suffix("RETURNING quantity")

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("ToSql(): %w", err)
	}

	return q, args, nil
}
