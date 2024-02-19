package domain_model

import "github.com/google/uuid"

// GetItemsQuantityRequest response model
type GetItemsQuantityRequest struct {
	StockID uuid.UUID `json:"stock_id"`
}

// GetItemQuantityRequest response model
type GetItemQuantityRequest struct {
	StockID uuid.UUID `json:"stock_id"`
	ItemID  uuid.UUID `json:"item_id"`
}

// GetItemsQuantityResponse response model
type GetItemsQuantityResponse struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity int64     `json:"quantity"`
}
