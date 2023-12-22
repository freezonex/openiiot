package tenant

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type TenantService struct {
	db *mysql.MySQL
}

var (
	service *TenantService
	once    sync.Once
)

func NewTenantService(db *mysql.MySQL) *TenantService {
	once.Do(func() {
		service = &TenantService{
			db: db,
		}
	})
	return service
}
