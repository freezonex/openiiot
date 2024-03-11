package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	yaml "gopkg.in/yaml.v2"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/service/utils/cache"
)

type DBConfig struct {
	Openiiot DBConfigData `yaml:"Openiiot"`
}

type DBConfigData struct {
	DSN      string `yaml:"DSN"`
	DB_DEBUG bool   `yaml:"DB_DEBUG"`
	CI       bool   `yaml:"CI"`
	DBNAME   string `yaml:"DBNAME"`
}

type SuposConfig struct {
	URL    string `yaml:"URL"`
	AK     string `yaml:"AK"`
	SK     string `yaml:"SK"`
	LOGOUT string `yaml:"LOGOUT"`
}

type K8sConfig struct {
	AuthorizationValue string `yaml:"AuthorizationValue"`
	K8SURL             string `yaml:"K8SURL"`
}

type GrafanaConfig struct {
	RedirectUri string `yaml:"RedirectUri"`
}

type TDEngineConfig struct {
	DSN string `yaml:"DSN"`
}

type HertzConfig struct {
	MaxRequestBodySize int `yaml:"MaxRequestBodySize"`
}

type GeneralConfig struct {
	OpeniiotUrl string `yaml:"OpeniiotUrl"`
}

type Config struct {
	IDC            string `yaml:"IDC"`
	DBConfig       `yaml:"DBConfig"`
	SuposConfig    `yaml:"SuposConfig"`
	GrafanaConfig  `yaml:"GrafanaConfig"`
	TDEngineConfig `yaml:"TDEngineConfig"`
	HertzConfig    `yaml:"HertzConfig"`
	GeneralConfig  `yaml:"GeneralConfig"`
	K8sConfig      `yaml:"K8sConfig"`
}

func Init() (*Config, error) {
	var c Config
	var configDir string

	args := os.Args
	fmt.Print(args)
	for _, v := range args {
		if strings.HasPrefix(v, "-conf-dir=") {
			configDir = v[10:]
		}
	}

	ctx := context.TODO()
	runtime_idc_name := os.Getenv("RUNTIME_IDC_NAME")
	if runtime_idc_name == "ci" || runtime_idc_name == "local" {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			return nil, errors.New("no caller information")
		}

		configDir = path.Join(path.Dir(filename), "conf")
	}

	configFile := path.Join(configDir, runtime_idc_name+".yml")
	logs.CtxInfof(ctx, "config file path: %v", configFile)
	cf, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(cf, &c); err != nil {
		return nil, err
	}

	// override by env param
	sdkAddress := os.Getenv("SDK_ADDRESS")
	if sdkAddress != "" {
		c.SuposConfig.URL = sdkAddress
	} else {
		suposApiAddress := os.Getenv("SUPOS_SUPOS_ADDRESS")
		if suposApiAddress != "" {
			c.SuposConfig.URL = suposApiAddress
		}
	}

	ak := os.Getenv("SUPOS_APP_ACCOUNT_ID")
	if ak != "" {
		c.SuposConfig.AK = ak
	}

	sk := os.Getenv("SUPOS_SUPOS_APP_SECRET_KEY")
	if sk != "" {
		c.SuposConfig.SK = sk
	}

	//only for debug purpose, will be override when connected SupOS
	currentUsername := os.Getenv("SUPOS_CURRENT_USERNAME")
	currentRoleName := os.Getenv("SUPOS_CURRENT_ROLENAME")

	cache.Set("CurrentUserId", currentUsername)
	cache.Set("CurrentUsername", currentUsername)
	cache.Set("CurrentUserRole", currentRoleName)

	return &c, nil
}
