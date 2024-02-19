package stock_domain

import (
	"errors"
	"net/http"

	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/rs/zerolog/log"
)

func (d *stockDomain) GetItemsQuantity(r *http.Request, args *domain_model.GetItemsQuantityRequest, response *interface{}) error {
	itemsQuantity, err := d.service.GetItemsQuantityOnStock(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		if errors.Is(err, custom_error.ErrNotFound) {
			return custom_error.ErrNotFound
		} else if errors.Is(err, custom_error.ErrUnavailableStock) {
			return custom_error.ErrUnavailableStock
		} else {
			return custom_error.ErrUnexpected
		}
	}

	*response = domain_model.BuildResponse(itemsQuantity).Result

	return nil
}
