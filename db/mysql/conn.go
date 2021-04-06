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
