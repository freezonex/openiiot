package core

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type CoreService struct {
	db *mysql.MySQL
}

var (
	service *CoreService
	once    sync.Once
)

func NewCoreService(db *mysql.MySQL) *CoreService {
	once.Do(func() {
		service = &CoreService{
			db: db,
		}
	})
	return service
}

func DefaultCoreService() *CoreService {
	return service
}
