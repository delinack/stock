package item_domain

import (
	"net/http"
	"storage/internal/pkg/domain_model"

	"github.com/rs/zerolog/log"
)

// ReserveItemsForDelivery резервирование товара на складе для доставки
func (d *itemDomain) ReserveItemsForDelivery(r *http.Request, args *domain_model.ReserveItemsOnStockForDeliveryRequest, response *domain_model.Response) error {
	err := d.service.ReserveItems(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		return err // TODO человекочитаемая ошибка
	}

	*response = domain_model.BuildResponse("items have been reserved")

	return nil
}

func (d *itemDomain) DeleteItemsReservation(r *http.Request, args *domain_model.DeleteItemsReserveRequest, response *domain_model.Response) error {
	err := d.service.DeleteItemsReservation(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		return err // TODO человекочитаемая ошибка
	}

	*response = domain_model.BuildResponse("item's reservation have been deleteed")

	return nil
}
