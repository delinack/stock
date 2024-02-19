package stock_domain

import (
	"net/http"

	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/service"
)

// StockDomain for templates domain methods
type StockDomain interface {
	GetItemsQuantity(*http.Request, *domain_model.GetItemsQuantityRequest, *interface{}) error
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
