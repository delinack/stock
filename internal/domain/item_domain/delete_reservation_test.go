package item_domain

import (
	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/mock/storage_mock"
	"github.com/delinack/stock/internal/pkg/service"
	"github.com/delinack/stock/internal/pkg/service/serializer"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestDeleteItemsReservation_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.DeleteItemsReserveRequest{
		StockID: uuid.New(),
		Items: []domain_model.DeleteItem{
			{
				ItemID: uuid.New(),
			},
		},
	}

	reservedItemModel := serializer.ToReservedItemsModelFromDeleteRequest(*params)

	storages.EXPECT().DeleteReservation(gomock.Any(), reservedItemModel).Return(nil)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.DeleteItemsReservation(&http.Request{}, params, &resp.Result)

	require.NoError(t, err)
	require.Equal(t, "item's reservation have been deleted", resp.Result)
}

func TestDeleteItemsReservation_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.DeleteItemsReserveRequest{
		StockID: uuid.New(),
		Items: []domain_model.DeleteItem{
			{
				ItemID: uuid.New(),
			},
		},
	}

	reservedItemModel := serializer.ToReservedItemsModelFromDeleteRequest(*params)

	storages.EXPECT().DeleteReservation(gomock.Any(), reservedItemModel).Return(custom_error.ErrNotFound)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.DeleteItemsReservation(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNotFound)
	require.Equal(t, nil, resp.Result)
}
