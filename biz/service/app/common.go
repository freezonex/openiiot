package app

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type AppService struct {
	db *mysql.MySQL
}

var (
	service *AppService
	once    sync.Once
)

func NewAppService(db *mysql.MySQL) *AppService {
	once.Do(func() {
		service = &AppService{
			db: db,
		}
	})
	return service
}

func DefaultAppService() *AppService {
	return service
}
