package supos

import (
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

type SuposService struct {
	db *mysql.MySQL
	c  *config.SuposConfig
}

type ApiResponse struct {
	HttpCode     int    `json:"httpCode"`
	ResponseBody string `json:"responseBody"`
}

type Pagination struct {
	Total     int `json:"total"`
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

type User struct {
	Username    string `json:"username"`
	UserDesc    string `json:"userDesc"`
	AccountType int    `json:"accountType"`
	LockStatus  int    `json:"lockStatus"`
	PersonCode  string `json:"personCode"`
	PersonName  string `json:"personName"`
	CreateTime  string `json:"createTime"`
	ModifyTime  string `json:"modifyTime"`
}

type UserRole struct {
	Name        string `json:"name"`
	Showname    string `json:"showname"`
	Description string `json:"description"`
}

type UserDetail struct {
	Username     string     `json:"username"`
	UserDesc     string     `json:"userDesc"`
	AccountType  int        `json:"accountType"`
	LockStatus   int        `json:"lockStatus"`
	Valid        int        `json:"valid"`
	PersonCode   string     `json:"personCode"`
	PersonName   string     `json:"personName"`
	ModifyTime   string     `json:"modifyTime"`
	CreateTime   string     `json:"createTime"`
	UserRoleList []UserRole `json:"userRoleList"`
	Avatar       string     `json:"avatar"`
}

type PageResult struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	Pagination Pagination `json:"pagination"`
	List       []User     `json:"list"`
}

var (
	service *SuposService
	once    sync.Once
)

func NewSuposService(db *mysql.MySQL, c *config.SuposConfig) *SuposService {
	once.Do(func() {
		service = &SuposService{
			db: db,
			c:  c,
		}
	})
	return service
}
