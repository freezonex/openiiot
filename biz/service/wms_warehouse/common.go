package wms_warehouse

import (
	"freezonex/openiiot/biz/config"
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsWarehouseService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsWarehouseService
	once    sync.Once
)

func NewWmsWarehouseService(db *mysql.MySQL) *WmsWarehouseService {
	once.Do(func() {
		service = &WmsWarehouseService{
			db: db,
		}
	})
	return service
}

func DefaultWmsWarehouseService() *WmsWarehouseService {
	return service
}
