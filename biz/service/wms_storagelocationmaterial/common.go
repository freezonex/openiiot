package wms_storagelocationmaterial

import (
	"freezonex/openiiot/biz/dal/mysql"
	"sync"
)

type WmsStorageLocationMaterialService struct {
	db *mysql.MySQL
}

var (
	service1 *WmsStorageLocationMaterialService
	once1    sync.Once
)

func DefaultStorageLocationMaterialService() *WmsStorageLocationMaterialService {
	return service1
}

func NewStorageLocationMaterial(db *mysql.MySQL) *WmsStorageLocationMaterialService {
	once1.Do(func() {
		service1 = &WmsStorageLocationMaterialService{
			db: db,
		}
	})
	return service1
}
