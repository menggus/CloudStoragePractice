package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/cloudstore?charset=utf8mb4")
	db.SetMaxOpenConns(1024)
	err := db.Ping()
	if err != nil {
		log.Printf("db connect failed: %s", err)
		os.Exit(1)
	}
}

// DBConnect 返回数据库连接
func DBConnect() *sql.DB {
	return db
}

// ParseRows 解析 rows 返回字典数据
func ParseRows(rows *sql.Rows) []map[string]interface{} {

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
