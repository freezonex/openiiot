package tdengine

import (
	"fmt"
	"regexp"
	"sync"

	//_ "github.com/taosdata/driver-go/v3/taosRestful"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
)

type TDEngineService struct {
	db *mysql.MySQL
	c  *config.TDEngineConfig
	T  *TDEngineClient
}

type TDEngineDSN struct {
	Name     string `json:"name"`
	Host     string `json:"hoste"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TDEngineRow struct {
	Time  string
	Value float32
}

var (
	service *TDEngineService
	once    sync.Once
)

func NewTDEngineClient(dsn string) *TDEngineClient {
	return &TDEngineClient{
		DSN: dsn,
	}
}

func NewTDEngineService(db *mysql.MySQL, c *config.TDEngineConfig) *TDEngineService {
	once.Do(func() {
		tClient := NewTDEngineClient(c.DSN)
		service = &TDEngineService{
			db: db,
			c:  c,
			T:  tClient,
		}
	})
	return service
}

// ParseDSNString converts a DSN string to a TDEngineDSN struct
func ParseDSNString(dsn string) (*TDEngineDSN, error) {
	// Regular expression to extract the components from the DSN string
	re := regexp.MustCompile(`^(.+):(.+)@http\((.+):(\d+)\)/$`)
	matches := re.FindStringSubmatch(dsn)

	if len(matches) != 5 {
		return nil, fmt.Errorf("invalid DSN format")
	}

	return &TDEngineDSN{
		Username: matches[1],
		Password: matches[2],
		Host:     matches[3],
		// Assuming 'Name' field is not directly extracted from the DSN string
	}, nil
}

// ConstructDSNString constructs a DSN string from a TDEngineDSN struct
func ConstructDSNString(username string, password string, host string) string {
	return fmt.Sprintf("%s:%s@http(%s)/", username, password, host)
}

func DefaultTdengineService() *TDEngineService {
	return service
}
