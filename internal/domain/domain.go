package domain

import (
	"storage/internal/domain/item_domain"
	"storage/internal/domain/stock_domain"
	"storage/internal/pkg/service"
)

//// Domain for all domain interfaces
//type Domain interface {
//	stock_domain.StockDomain
//	item_domain.ItemDomain
//}

// Domain is a concrete structure that handles a bunch of requests from http
type Domain struct {
	StockDomain stock_domain.StockDomain
	ItemDomain  item_domain.ItemDomain
}

// NewDomain is a constructor that build domain instance
func NewDomain(service service.Service) *Domain {
	return &Domain{
		StockDomain: stock_domain.NewStockDomain(service),
		ItemDomain:  item_domain.NewItemDomain(service),
	}
}
