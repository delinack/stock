package item_domain

import (
	"storage/internal/pkg/service"
	"storage/internal/pkg/storage"
	"testing"
)

func TestReserveItemsForDelivery(t *testing.T) {

	storages := storage.NewStorage()
	services := service.NewService(storages)
	domain := NewItemDomain(services)
}
