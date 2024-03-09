package wms_rfid_material

import (
	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
	"sync"
)

type WmsRfidMaterialService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsRfidMaterialService
	once    sync.Once
)

func NewWmsRfidMaterialService(db *mysql.MySQL) *WmsRfidMaterialService {
	once.Do(func() {
		service = &WmsRfidMaterialService{
			db: db,
		}
	})
	return service
}

func DefaultWmsRfidMaterialService() *WmsRfidMaterialService {
	return service
}
