package stock_domain

import (
	"net/http"
	"storage/internal/pkg/domain_model"
	"storage/internal/pkg/service"
)

// StockDomain for templates domain methods
type StockDomain interface {
	GetItemsQuantity(*http.Request, *domain_model.GetItemsQuantityRequest, *domain_model.Response) error
}

type stockDomain struct {
	service service.Service
}

// NewStockDomain .
func NewStockDomain(service service.Service) StockDomain {
	return &stockDomain{
		service,
	}
}
