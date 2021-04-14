package db

import (
	mydb "cloudstorage/v1/db/mysql"
	"log"
	"time"
)

type UserFile struct {
	UserName    string
	FileName    string
	FileSha1    string
	FileSize    string
	UploadAt    string
	LastUpdated string
}

// TabUserFileQueryRows 查询所有文件信息
func TabUserFileQueryRows(u string) ([]map[string]interface{}, error) {

	// todo 查询数量的控制
	stmt, err := mydb.DBConnect().Prepare("SELECT file_sha1, file_name, file_size, upload_at, last_update FROM tabuserfile " +
		"WHERE user_name=? AND status=0 limit 15")
	if err != nil {
		log.Printf("TabUserFileQueryRows sql prepare failed: %s\n", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(u)
	if err != nil {
		log.Printf("TabUserFileQueryRows sql query failed: %s\n", err)
		return nil, err
	}

	defer rows.Close()
	userfiles := mydb.ParseRows(rows)

	return userfiles, nil
}

// TabUserFileInsert  用户文件上传
func TabUserFileInsert(username, sha1, filename string, filesize int64) bool {
	stmt, err := mydb.DBConnect().Prepare("INSERT IGNORE INTO tabuserfile (user_name, file_name, file_sha1, " +
		"file_size, upload_at) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Printf("sql prepare failed: %s\n", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, filename, sha1, filesize, time.Now())
	if err != nil {
		log.Printf("sql exec failed: %s\n", err)
		return false
	}

	if row, err := ret.RowsAffected(); err == nil {

		if row < 0 { // row < 0 sql 插入数据失败，数据已存在
			log.Printf("File with hash[%s] already exists\n", sha1)

			return false
		}

		if row == 0 { // row < 0 sql 插入数据失败，数据已存在
			log.Printf("File with hash[%s] already exists\n", sha1)
			// todo 这里插入数据失败，可能文件已经上传过

			return true
		}
	}
	return true
}
