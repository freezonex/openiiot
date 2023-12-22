package mysql

import (
	"context"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	//"code.byted.org/gopkg/env"
	//"github.com/sirupsen/logrus"
	"freezonex/openiiot/biz/config"

	"gorm.io/driver/mysql"
	//"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/dal/query_openiiot"
)

var (
	dbOpeniiot  *gorm.DB
	dsnOpeniiot = "root:root1234@tcp(127.0.0.1:3306)/openiiot?charset=utf8mb4&parseTime=True&loc=Local"
	//dbOpeniiotName = "openiiot_db"
	//psmOpeniiot    = "freezonex.mysql.openiiot_db"
)

type MySQL struct {
	DBOpeniiotQuery *query_openiiot.Query
}

func Init(c *config.DBConfig) (*MySQL, error) {
	var m MySQL
	// CI environment
	if c.Openiiot.CI {
		dbOpeniiot = connectSQLiteOpeniiot().Debug()
		m.DBOpeniiotQuery = query_openiiot.Use(dbOpeniiot)
		return &m, nil
	}
	populateDBConfig(c)

	/*DBOpeniiot, err := gorm.Open(bytedgorm.MySQL(psmOpeniiot, dbOpeniiotName).With(func(conf *bytedgorm.DBConfig) {
		conf.DSN = c.Openiiot.DSN

	}), bytedgorm.WithDefaults())*/
	//dsnOpeniiot = "root:root1234@tcp(127.0.0.1:6033)/openiiot?charset=utf8mb4&parseTime=True&loc=Local"

	DBOpeniiot, err := gorm.Open(mysql.Open(dsnOpeniiot), &gorm.Config{})

	// check connect error
	if err != nil {
		return nil, err
	}

	dbOpeniiot = DBOpeniiot
	idc := os.Getenv("RUNTIME_IDC_NAME")
	if idc == "boe" || idc == "boei18n" {
		dbOpeniiot = dbOpeniiot.Debug()
	}
	if c.Openiiot.DB_DEBUG {
		dbOpeniiot = dbOpeniiot.Debug()
	}
	m.DBOpeniiotQuery = query_openiiot.Use(dbOpeniiot)
	return &m, nil
}

// Read write separation (openiiot db)
func DBOpeniiot(ctx context.Context) *gorm.DB {
	return dbOpeniiot.WithContext(ctx)
}

func populateDBConfig(c *config.DBConfig) {
	populateDBOpeniiotConfig(c)
}

func populateDBOpeniiotConfig(c *config.DBConfig) {
	if len(c.Openiiot.DSN) != 0 {
		dsnOpeniiot = c.Openiiot.DSN
	}
	/*if len(c.Openiiot.DBNAME) != 0 {
		dbOpeniiotName = c.Openiiot.DBNAME
	}*/
	//logs.Info("Openiiot Database config, psm = %v, dbName = %v ", psmOpeniiot, dbOpeniiotName)
}

func connectSQLiteOpeniiot() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared&parseTime=True"))
	if err != nil {
		//logs.Fatalf("error connecting to sqlite DB, err: %s", err)
	}
	//logs.Info("init sqlite success")
	/*err = DB.AutoMigrate(
		&model_openiiot.TestDataNamespace{},
		&model_openiiot.TestDatum{},
		&model_openiiot.TestDataHistory{},
		&model_openiiot.TestDataOperator{},
		&model_openiiot.TestDataTag{},
		&model_openiiot.TestDataTagMapping{},
		&model_openiiot.TestDataDirectory{},
	)
	if err != nil {
		//logs.Fatalf("error while doing DB auto migrate, err: %s", err)
	}*/
	return DB
}
