package wms_inbound

import (
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

type WmsInboundService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsInboundService
	once    sync.Once
)

func NewWmsInboundService(db *mysql.MySQL) *WmsInboundService {
	once.Do(func() {
		service = &WmsInboundService{
			db: db,
		}
	})
	return service
}

func DefaultWmsInboundService() *WmsInboundService {
	return service
}
