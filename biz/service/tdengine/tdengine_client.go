package tdengine

import (
	db "database/sql"
	"encoding/json"
	"strings"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

type TDEngineClient struct {
	DSN string
	//taos *sql.DB
}

func (t TDEngineClient) TestConnect(DSN string) bool {
	//DSN = "root:taosdata@http(localhost:6041)/"
	taos, err := db.Open("taosRestful", DSN)
	if err != nil {
		logs.Errorf("failed to initialize connection to TDengine, %v", err)
		return false
	}
	defer taos.Close()

	// Execute a simple query to test the connection
	_, err = taos.Exec("SELECT SERVER_STATUS()")
	if err != nil {
		logs.Errorf("failed to connect to TDengine, %v", err)
		return false
	}

	return true
}

func (t TDEngineClient) Exec(DSN string, sql string) (int64, error) {
	//DSN = "root:taosdata@http(localhost:6041)/"
	taos, err := db.Open("taosRestful", DSN)
	if err != nil {
		logs.Errorf("failed to initialize connection to TDengine, err: %v", err)
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

// use transaction to batch execute sql split by ";"
func (t TDEngineClient) BatchExec(DSN string, sqlBatch string) (int64, error) {
	taos, err := db.Open("taosRestful", DSN)
	if err != nil {
		logs.Errorf("failed to initialize connection to TDengine, err: %v", err)
		return 0, err
	}
	defer taos.Close()

	tx, err := taos.Begin()
	if err != nil {
		logs.Errorf("failed to begin transaction, err: %v", err)
		return 0, err
	}

	var totalRowsAffected int64 = 0
	sqlStatements := strings.Split(sqlBatch, ";")
	for _, sql := range sqlStatements {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		result, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			logs.Errorf("failed to Exec %v, err: %v", sql, err)
			return totalRowsAffected, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			logs.Errorf("failed to get affected rows for %v, err: %v", sql, err)
			return totalRowsAffected, err
		}
		totalRowsAffected += rowsAffected
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		logs.Errorf("failed to commit transaction, err: %v", err)
		return totalRowsAffected, err
	}

	return totalRowsAffected, nil
}

func (t TDEngineClient) Query(DSN string, sql string) (string, error) {
	//DSN = "root:taosdata@http(localhost:6041)/power"
	taos, err := db.Open("taosRestful", DSN)
	if err != nil {
		logs.Error("failed to connect TDengine, err:", err)
	}
	defer taos.Close()

	rows, err := taos.Query(sql)
	if err != nil {
		logs.Error("failed to select from table, err:", err)
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logs.Error("failed to get columns:", err)
		return "", err
	}

	var results []map[string]interface{}
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			logs.Error("failed to scan row:", err)
			return "", err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			row[col] = v
		}
		results = append(results, row)
	}

	jsonString, err := json.Marshal(results)
	if err != nil {
		return "", err
	}

	jsonStr := string(jsonString)
	return jsonStr, nil
}
