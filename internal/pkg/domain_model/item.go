package domain_model

import "github.com/google/uuid"

// ReserveItemsOnStockForDeliveryRequest request model
type ReserveItemsOnStockForDeliveryRequest struct {
	StockID uuid.UUID     `json:"stock_id" validate:"uuid4,required"`
	Items   []ReserveItem `json:"items"`
}

// ReserveItem request model
type ReserveItem struct {
	ItemID   uuid.UUID `json:"id" validate:"uuid4,required"`
	Quantity int64     `json:"quantity" validate:"int64,required"`
}

// DeleteItemsReserveRequest request model
type DeleteItemsReserveRequest struct {
	StockID uuid.UUID    `json:"stock_id" validate:"uuid4,required"`
	Items   []DeleteItem `json:"items"`
}

// DeleteItem request model
type DeleteItem struct {
	ItemID uuid.UUID `json:"id" validate:"uuid4,required"`
}
