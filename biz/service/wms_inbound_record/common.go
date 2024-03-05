package wms_inbound_record

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsInboundRecordService struct {
	db *mysql.MySQL
}

var (
	service *WmsInboundRecordService
	once    sync.Once
)

func NewWmsInboundRecordService(db *mysql.MySQL) *WmsInboundRecordService {
	once.Do(func() {
		service = &WmsInboundRecordService{
			db: db,
		}
	})
	return service
}

func DefaultWmsInboundRecordService() *WmsInboundRecordService {
	return service
}
