package stock_domain

import (
	"net/http"
	"storage/internal/pkg/domain_model"

	"github.com/rs/zerolog/log"
)

func (d *stockDomain) GetItemsQuantity(r *http.Request, args *domain_model.GetItemsQuantityRequest, response *domain_model.Response) error {
	itemsQuantity, err := d.service.GetItemsQuantityOnStock(r.Context(), args)
	if err != nil {
		log.Error().Err(err).Send()
		return err // TODO человекочитаемая ошибка
	}

	*response = domain_model.BuildResponse(itemsQuantity)

	return nil
}
