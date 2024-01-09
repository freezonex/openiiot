package edge

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type EdgeService struct {
	db *mysql.MySQL
}

var (
	service *EdgeService
	once    sync.Once
)

func NewEdgeService(db *mysql.MySQL) *EdgeService {
	once.Do(func() {
		service = &EdgeService{
			db: db,
		}
	})
	return service
}
