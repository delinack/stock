package stock_domain

import (
	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/mock/storage_mock"
	"github.com/delinack/stock/internal/pkg/model"
	"github.com/delinack/stock/internal/pkg/service"
	"github.com/delinack/stock/internal/pkg/service/serializer"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"testing"
)

func TestGetItemsQuantity_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.GetItemsQuantityRequest{
		StockID: uuid.New(),
	}

	items := []model.ItemStock{
		{
			ItemID:   uuid.New(),
			Quantity: null.IntFrom(1),
		},
		{
			ItemID:   uuid.New(),
			Quantity: null.IntFrom(2),
		},
		{
			ItemID:   uuid.New(),
			Quantity: null.IntFrom(5),
		},
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemsQuantity(gomock.Any(), params).Return(items, nil)

	services := service.NewService(storages)
	domain := NewStockDomain(services)

	resp := domain_model.Response{}
	err := domain.GetItemsQuantity(&http.Request{}, params, &resp.Result)

	expectedResp := serializer.ToGetItemsQuantityResponse(items)

	require.NoError(t, err)
	require.Equal(t, expectedResp, resp.Result)
}

func TestGetItemsQuantity_StockUnavailable(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.GetItemsQuantityRequest{
		StockID: uuid.New(),
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(false, custom_error.ErrUnavailableStock)

	services := service.NewService(storages)
	domain := NewStockDomain(services)

	resp := domain_model.Response{}
	err := domain.GetItemsQuantity(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrUnavailableStock)
	require.Equal(t, nil, resp.Result)
}

func TestGetItemsQuantity_StockNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.GetItemsQuantityRequest{
		StockID: uuid.New(),
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(false, custom_error.ErrNotFound)

	services := service.NewService(storages)
	domain := NewStockDomain(services)

	resp := domain_model.Response{}
	err := domain.GetItemsQuantity(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNotFound)
	require.Equal(t, nil, resp.Result)
}

func TestGetItemsQuantity_ItemNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.GetItemsQuantityRequest{
		StockID: uuid.New(),
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemsQuantity(gomock.Any(), params).Return(nil, custom_error.ErrNotFound)

	services := service.NewService(storages)
	domain := NewStockDomain(services)

	resp := domain_model.Response{}
	err := domain.GetItemsQuantity(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNotFound)
	require.Equal(t, nil, resp.Result)
}
