package wms_rfid_material

import (
	"freezonex/openiiot/biz/config"
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsRfidRfidMaterialService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsRfidRfidMaterialService
	once    sync.Once
)

func NewWmsRfidRfidMaterialService(db *mysql.MySQL) *WmsRfidRfidMaterialService {
	once.Do(func() {
		service = &WmsRfidRfidMaterialService{
			db: db,
		}
	})
	return service
}

func DefaultWmsRfidRfidMaterialService() *WmsRfidRfidMaterialService {
	return service
}
