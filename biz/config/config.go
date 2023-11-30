package config

import (
	"errors"
	"os"
	"path"
	"runtime"

	yaml "gopkg.in/yaml.v2"
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
}

func Init() (*Config, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("no caller information")
	}
	var c Config
	configFile := path.Join(path.Dir(filename), "conf/config.yml")
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

	return &c, nil
}
