package item_domain

import (
	"net/http"

	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/service"
)

// ItemDomain for templates domain methods
type ItemDomain interface {
	ReserveItemsForDelivery(*http.Request, *domain_model.ReserveItemsOnStockForDeliveryRequest, *interface{}) error
	DeleteItemsReservation(*http.Request, *domain_model.DeleteItemsReserveRequest, *interface{}) error
}

type itemDomain struct {
	service service.Service
}

// NewItemDomain .
func NewItemDomain(service service.Service) ItemDomain {
	return &itemDomain{
		service,
	}
}
