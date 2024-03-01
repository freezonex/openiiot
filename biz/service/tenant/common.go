package tenant

import (
	"freezonex/openiiot/biz/config"
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type TenantService struct {
	db *mysql.MySQL
	S  *config.K8sConfig
}

var (
	service *TenantService
	once    sync.Once
)

func NewTenantService(db *mysql.MySQL, s *config.K8sConfig) *TenantService {
	once.Do(func() {
		service = &TenantService{
			db: db,
			S:  s,
		}
	})
	return service
}

func DefaultTenantService() *TenantService {
	return service
}
