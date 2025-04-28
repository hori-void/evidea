package controllers

import (
	"database/sql"
	// "fmt"
	// "log"
	// "controllers"

	_ "github.com/lib/pq" // PostgreSQLドライバ（適宜変更）
)

// func ExecuteQuery(query string) ([]map[string]interface{}, error) {
func ExecuteQuery(query string, isSelect bool) (interface{}, error) {
	connStr := "host=host.docker.internal user=user password=evidea dbname=evidea_db port=5432 sslmode=disable connect_timeout=10"
	// DB接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close() // 処理後にDBを切断

	// クエリの種類によって処理を分ける
	if isSelect {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// 結果をスライスに格納
		var results []map[string]interface{}
		cols, _ := rows.Columns()
		for rows.Next() {
			// 各カラムの値を格納
			values := make([]interface{}, len(cols))
			valuePtrs := make([]interface{}, len(cols))
			for i := range values {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				return nil, err
			}

			// マップに変換
			rowMap := make(map[string]interface{})
			for i, col := range cols {
				rowMap[col] = values[i]
			}
			results = append(results, rowMap)
		}
		return results, nil

	} else {
		// SELECT以外（INSERT, UPDATE, DELETEなど）の場合
		result, err := db.Exec(query)
		if err != nil {
			return nil, err
		}

		affectedRows, err := result.RowsAffected()
		if err != nil {
			return nil, err
		}

		return affectedRows, nil
	}
}

// var results []map[string]interface{}

// for rows.Next() {
// 	// 各カラムの値を格納するスライス
// 	scanArgs := make([]interface{}, len(columns))
// 	values := make([]interface{}, len(columns))
// 	for i := range values {
// 		scanArgs[i] = &values[i]
// 	}

// 	if err := rows.Scan(scanArgs...); err != nil {
// 		return nil, err
// 	}

// 	// 1行分のデータをマップに格納
// 	rowMap := make(map[string]interface{})
// 	for i, colName := range columns {
// 		rowMap[colName] = values[i]
// 	}
// 	results = append(results, rowMap)
// }

// if err := rows.Err(); err != nil {
// 	return nil, err
// }

// return results, nil
