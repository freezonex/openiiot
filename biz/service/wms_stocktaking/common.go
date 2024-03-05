package wms_stocktaking

import (
	"freezonex/openiiot/biz/dal/mysql"
	"sync"
)

type WmsStocktakingService struct {
	db *mysql.MySQL
}

var (
	service *WmsStocktakingService
	once    sync.Once
)

func DefaultStocktakingService() *WmsStocktakingService {
	return service
}

func NewStocktakingService(db *mysql.MySQL) *WmsStocktakingService {
	once.Do(func() {
		service = &WmsStocktakingService{
			db: db,
		}
	})
	return service
}
