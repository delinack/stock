package item_domain

import (
	"net/http"
	"storage/internal/pkg/domain_model"
	"storage/internal/pkg/service"
)

// ItemDomain for templates domain methods
type ItemDomain interface {
	ReserveItemsForDelivery(*http.Request, *domain_model.ReserveItemsOnStockForDeliveryRequest, *domain_model.Response) error
	DeleteItemsReservation(*http.Request, *domain_model.DeleteItemsReserveRequest, *domain_model.Response) error
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
