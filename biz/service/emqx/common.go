package emqx

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type EmqxService struct {
	db *mysql.MySQL
	client *EmqxClient
}

var (
	service *EmqxService
	once    sync.Once
)

func NewEmqxService(db *mysql.MySQL) *EmqxService {
	once.Do(func() {
		service = &EmqxService{
			db: db,
		}
	})
	return service
}
