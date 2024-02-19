package item_domain

import (
	"net/http"
	"testing"

	"github.com/delinack/stock/internal/pkg/custom_error"
	"github.com/delinack/stock/internal/pkg/domain_model"
	"github.com/delinack/stock/internal/pkg/mock/storage_mock"
	"github.com/delinack/stock/internal/pkg/service"
	"github.com/delinack/stock/internal/pkg/service/serializer"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestReserveItemsForDelivery_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}
	reserveItemModel := serializer.ToReserveItemsModelFromReserveRequest(*params)
	itemModel := serializer.ToItemsModelFromReserveRequest(*params)

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemQuantity(gomock.Any(), params.StockID, params.Items[0]).Return(int64(5), nil)
	storages.EXPECT().ReserveItems(gomock.Any(), reserveItemModel, itemModel).Return(nil)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.NoError(t, err)
	require.Equal(t, "items have been reserved", resp.Result)
}

func TestReserveItemsForDelivery_StockNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(false, custom_error.ErrNotFound)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNotFound)
	require.Equal(t, nil, resp.Result)
}

func TestReserveItemsForDelivery_StockUnavailable(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(false, custom_error.ErrUnavailableStock)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrUnavailableStock)
	require.Equal(t, nil, resp.Result)
}

func TestReserveItemsForDelivery_NullItemValue(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemQuantity(gomock.Any(), params.StockID, params.Items[0]).Return(int64(0), custom_error.ErrNullValue)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNullValue)
	require.Equal(t, nil, resp.Result)
}

func TestReserveItemsForDelivery_ExceededValue(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemQuantity(gomock.Any(), params.StockID, params.Items[0]).Return(int64(6), custom_error.ErrExceededValue)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrExceededValue)
	require.Equal(t, nil, resp.Result)
}

func TestReserveItemsForDelivery_ItemNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	storages := storage_mock.NewMockStore(ctrl)

	params := &domain_model.ReserveItemsOnStockForDeliveryRequest{
		StockID: uuid.New(),
		Items: []domain_model.ReserveItem{
			{
				ItemID:   uuid.New(),
				Quantity: int64(5),
			},
		},
	}
	reserveItemModel := serializer.ToReserveItemsModelFromReserveRequest(*params)
	itemModel := serializer.ToItemsModelFromReserveRequest(*params)

	storages.EXPECT().CheckAvailability(gomock.Any(), params.StockID).Return(true, nil)
	storages.EXPECT().GetItemQuantity(gomock.Any(), params.StockID, params.Items[0]).Return(int64(5), nil)
	storages.EXPECT().ReserveItems(gomock.Any(), reserveItemModel, itemModel).Return(custom_error.ErrNotFound)

	services := service.NewService(storages)
	domain := NewItemDomain(services)

	resp := domain_model.Response{}
	err := domain.ReserveItemsForDelivery(&http.Request{}, params, &resp.Result)

	require.ErrorIs(t, err, custom_error.ErrNotFound)
	require.Equal(t, nil, resp.Result)
}
