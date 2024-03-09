package wms_storage_location

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsStorageLocationService struct {
	db *mysql.MySQL
}

var (
	service *WmsStorageLocationService
	once    sync.Once
)

func DefaultStorageLocationService() *WmsStorageLocationService {
	return service
}

func NewStorageLocationService(db *mysql.MySQL) *WmsStorageLocationService {
	once.Do(func() {
		service = &WmsStorageLocationService{
			db: db,
		}
	})
	return service
}
