package stock_storage

import (
	"fmt"
	"stock/internal/pkg/domain_model"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func buildGetItemsQuantityQuery(params *domain_model.GetItemsQuantityRequest) (string, []interface{}, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("item_id", "quantity").
		From("items_stocks").
		Where(sq.Eq{"stock_id": params.StockID})

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql() failed: %w", err)
	}

	return q, args, nil
}

func buildGetItemQuantityQuery(stockID uuid.UUID, params domain_model.ReserveItem) (string, []interface{}, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("quantity").
		From("items_stocks").
		Where(sq.Eq{"stock_id": stockID}).
		Where(sq.Eq{"item_id": params.ItemID})

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql() failed: %w", err)
	}

	return q, args, nil
}

func buildCheckStockAvailability(stockID uuid.UUID) (string, []interface{}, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("is_available").
		From("stocks").
		Where(sq.Eq{"id": stockID})

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql() failed: %w", err)
	}

	return q, args, nil
}
