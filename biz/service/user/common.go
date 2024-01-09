package user

import (
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type UserService struct {
	db *mysql.MySQL
}

var (
	service *UserService
	once    sync.Once
)

func NewUserService(db *mysql.MySQL) *UserService {
	once.Do(func() {
		service = &UserService{
			db: db,
		}
	})
	return service
}

type SuposService struct {
	db *mysql.MySQL
}

var (
	service1 *SuposService
	once1    sync.Once
)

func NewSuposService(db *mysql.MySQL) *SuposService {
	once.Do(func() {
		service1 = &SuposService{
			db: db,
		}
	})
	return service1
}
