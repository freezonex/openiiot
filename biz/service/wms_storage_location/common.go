package storagelocation

import (
	"freezonex/openiiot/biz/dal/mysql"
	"sync"
)

type StorageLocationService struct {
	db *mysql.MySQL
}

var (
	service *StorageLocationService
	once    sync.Once
)

func DefaultStorageLocationService() *StorageLocationService {
	return service
}

func NewStorageLocationService(db *mysql.MySQL) *StorageLocationService {
	once.Do(func() {
		service = &StorageLocationService{
			db: db,
		}
	})
	return service
}
