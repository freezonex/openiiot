package user

import (
	"context"
	"freezonex/openiiot/biz/service/callback_mgr"
	"sync"
	"time"

	"freezonex/openiiot/biz/dal/mysql"
)

func init() {
	callback_mgr.RegisterCallBack("GetUserByTokenDB", func(ctx context.Context, usertoken string) (*time.Time, string, error) {
		return DefaultUserService().GetUserByTokenDB(ctx, usertoken)
	})
}

type UserService struct {
	db *mysql.MySQL
}

var (
	service *UserService
	once    sync.Once
)

func DefaultUserService() *UserService {
	return service
}

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
