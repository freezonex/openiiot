package flow

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type FlowService struct {
	db *mysql.MySQL
}

var (
	service *FlowService
	once    sync.Once
)

func NewFlowService(db *mysql.MySQL) *FlowService {
	once.Do(func() {
		service = &FlowService{
			db: db,
		}
	})
	return service
}
