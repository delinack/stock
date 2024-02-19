package item_domain

import (
	"errors"
	"github.com/delinack/stock/internal/pkg/custom_error"
	"net/http"

	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/rs/zerolog/log"
)

func (d *itemDomain) ReserveItemsForDelivery(r *http.Request, args *domain_model.ReserveItemsOnStockForDeliveryRequest, response *interface{}) error {
	err := d.service.ReserveItems(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		if errors.Is(err, custom_error.ErrNotFound) {
			return custom_error.ErrNotFound
		} else if errors.Is(err, custom_error.ErrUnavailableStock) {
			return custom_error.ErrUnavailableStock
		} else if errors.Is(err, custom_error.ErrNullValue) {
			return custom_error.ErrNullValue
		} else if errors.Is(err, custom_error.ErrExceededValue) {
			return custom_error.ErrExceededValue
		} else {
			return custom_error.ErrUnexpected
		}
	}

	*response = domain_model.BuildResponse("items have been reserved").Result

	return nil
}
