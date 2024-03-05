package wms_outbound

import (
	"freezonex/openiiot/biz/config"
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsOutboundService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsOutboundService
	once    sync.Once
)

func NewWmsOutboundService(db *mysql.MySQL) *WmsOutboundService {
	once.Do(func() {
		service = &WmsOutboundService{
			db: db,
		}
	})
	return service
}

func DefaultWmsOutboundService() *WmsOutboundService {
	return service
}
