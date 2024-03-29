package wms_material

import (
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

type WmsMaterialService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *WmsMaterialService
	once    sync.Once
)

func NewWmsMaterialService(db *mysql.MySQL) *WmsMaterialService {
	once.Do(func() {
		service = &WmsMaterialService{
			db: db,
		}
	})
	return service
}

func DefaultWmsMaterialService() *WmsMaterialService {
	return service
}
