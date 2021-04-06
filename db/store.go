package db

import (
	mydb "cloudstorage/v1/db/mysql"
	"log"
)

// TabFileDataInsert  向表tabfile插入数据
func TabFileDataInsert(fileSha1 string, fileName string, fileSize int64, fileAddr string) bool {
	// 准备sql
	// 这里也有防止sql注入的情况
	stmt, err := mydb.DBConnect().Prepare("INSERT ignore into tabfile (`file_sha1`, `file_name`, `file_size`," +
		" `file_addr`, `status`) values (?, ?, ?, ?, 1)")
	if err != nil {
		log.Printf("Failed to prepare statement: %s\n", err)
		return false
	}

	// 关闭连接
	defer stmt.Close()

	// 执行sql
	ret, err := stmt.Exec(fileSha1, fileName, fileSize, fileAddr)
	if err != nil { // sql 执行失败
		log.Printf("err: %s\n", err)
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf < 0 { // rf < 0 sql 插入数据失败
			log.Printf("File with hash[%s] already exists\n", fileSha1)
			return false
		}
		return true
	}

	return false
}
