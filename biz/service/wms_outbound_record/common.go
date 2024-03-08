package wms_outbound_record

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsOutboundRecordService struct {
	db *mysql.MySQL
}

var (
	service *WmsOutboundRecordService
	once    sync.Once
)

func NewWmsOutboundRecordService(db *mysql.MySQL) *WmsOutboundRecordService {
	once.Do(func() {
		service = &WmsOutboundRecordService{
			db: db,
		}
	})
	return service
}

func DefaultWmsOutboundRecordService() *WmsOutboundRecordService {
	return service
}
