package wms

import (
	"freezonex/openiiot/biz/dal/mysql"
	"sync"
)

type WmsService struct {
	db *mysql.MySQL
}

var (
	service *WmsService
	once    sync.Once
)

func NewTenantService(db *mysql.MySQL) *WmsService {
	once.Do(func() {
		service = &WmsService{
			db: db,
		}
	})
	return service
}

func DefaultTenantService() *WmsService {
	return service
}
