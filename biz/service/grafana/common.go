package grafana

import (
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

type GrafanaService struct {
	db *mysql.MySQL
	c  *config.GrafanaConfig
}

type OAuthRequest struct {
	Code        string `form:"code"`
	GrantType   string `form:"grant_type"`
	RedirectURI string `form:"redirect_uri"`
}

var (
	service *GrafanaService
	once    sync.Once
)

func NewGrafanaService(db *mysql.MySQL, c *config.GrafanaConfig) *GrafanaService {
	once.Do(func() {
		service = &GrafanaService{
			db: db,
			c:  c,
		}
	})
	return service
}
