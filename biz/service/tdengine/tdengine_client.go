package tdengine

import (
	db "database/sql"
	"log"
	"time"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

type TDEngineClient struct {
	DSN string
	//taos *sql.DB
}

func (t TDEngineClient) TestConnect(DSN string) bool {
	//var taosUri = "root:taosdata@http(localhost:6041)/"
	var taosDSN = DSN
	taos, err := db.Open("taosRestful", taosDSN)
	if err != nil {
		logs.Errorf("failed to connect TDengine, %v", err)
		return false
	}
	defer taos.Close()
	return true
}

func (t TDEngineClient) Exec(DSN string, sql string) (int64, error) {
	//var taosDSN = "root:taosdata@http(localhost:6041)/"
	var taosDSN = DSN
	taos, err := db.Open("taosRestful", taosDSN)
	if err != nil {
		logs.Errorf("failed to connect TDengine, err: %v", err)
		return 0, err
	}
	defer taos.Close()
	result, err := taos.Exec(sql)
	if err != nil {
		logs.Errorf("failed to Exec %v, err: %v", sql, err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logs.Errorf("failed to get affected rows, err: %v", err)
		return 0, err
	}
	logs.Info("RowsAffected", rowsAffected)
	return rowsAffected, nil
}

func (t TDEngineClient) Query(DSN string, sql string) ([]TDEngineRow, error) {
	//var taosDSN = "root:taosdata@http(localhost:6041)/power"
	var taosDSN = DSN
	taos, err := db.Open("taosRestful", taosDSN)
	if err != nil {
		logs.Error("failed to connect TDengine, err:", err)
	}
	defer taos.Close()

	rows, err := taos.Query(sql)
	if err != nil {
		logs.Error("failed to select from table, err:", err)
	}
	defer rows.Close()
	var results []TDEngineRow
	for rows.Next() {
		var r TDEngineRow
		var ts time.Time
		err := rows.Scan(&ts, &r.Value)
		if err != nil {
			log.Printf("scan error: %v\n", err)
			return nil, err
		}
		r.Time = ts.Format(time.RFC3339) // Or any other format you prefer
		results = append(results, r)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows error: %v\n", err)
		return nil, err
	}

	log.Printf("query %d data\n", len(results))
	return results, nil
}
