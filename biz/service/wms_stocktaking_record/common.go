package wms_stocktaking_record

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type WmsStocktakingRecordService struct {
	db *mysql.MySQL
}

var (
	service *WmsStocktakingRecordService
	once    sync.Once
)

func DefaultStocktakingRecordService() *WmsStocktakingRecordService {
	return service
}

func NewStocktakingRecordService(db *mysql.MySQL) *WmsStocktakingRecordService {
	once.Do(func() {
		service = &WmsStocktakingRecordService{
			db: db,
		}
	})
	return service
}
