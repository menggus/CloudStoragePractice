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
	// scanArgs:  [指针地址1，指针地址2，指针地址3]

	// record := make(map[string]interface{})
	// 上述放到循环外层会出现bug，因为record为map类型，引用类型，当record改变，对其引用的slice也会改变其值
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		record := make(map[string]interface{})
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
