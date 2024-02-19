package item_domain

import (
	"errors"
	"net/http"

	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/rs/zerolog/log"
)

func (d *itemDomain) DeleteItemsReservation(r *http.Request, args *domain_model.DeleteItemsReserveRequest, response *interface{}) error {
	err := d.service.DeleteItemsReservation(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		if errors.Is(err, custom_error.ErrNotFound) {
			return custom_error.ErrNotFound
		} else {
			return custom_error.ErrUnexpected
		}
	}

	*response = domain_model.BuildResponse("item's reservation have been deleted").Result

	return nil
}
