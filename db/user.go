package db

import (
	mydb "cloudstorage/v1/db/mysql"
	"cloudstorage/v1/utils"
	"log"
)

// TabUserDataInsert 插入用户数据
func TabUserDataInsert(username string, passowrd string) bool {
	stmt, err := mydb.DBConnect().Prepare("INSERT IGNORE INTO tabuser (`user_name`, `user_pwd`) VALUES (?,?)")
	if err != nil {
		log.Printf("stmt Prepare failed: %s\n", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passowrd)
	if err != nil {
		log.Printf("写入用户数据失败 %s\n", err)
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf < 0 {
			log.Printf("写入用户数据失败，返回row：%d\n", rf)
			return false
		}
	}
	return true
}

// TabUserDataQuery 用户数据查询
func TabUserDataQuery(username string, password string) bool {
	stmt, err := mydb.DBConnect().Prepare("SELECT * FROM tabuser WHERE status=0 AND user_name=? limit 1")
	if err != nil {
		log.Printf("stmt Prepare failed: %s\n", err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Printf("查询失败 %s\n", err)
		return false
	} else if rows == nil {
		log.Printf("没有查询到用户")
		return false
	}

	// todo 校验没通过

	userpassword := utils.Sha1([]byte(password))

	log.Println(userpassword)
	return true
}

// TabTokenDataInsert 写入token
func TabTokenDataInsert(username string, token string) bool {
	stmt, err := mydb.DBConnect().Prepare("INSERT IGNORE INTO tabtoken (`user_name`, `user_token`) values (?,?)")
	if err != nil {
		log.Printf("token 写入失败： %s\n", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, token)
	if err != nil {
		log.Printf("写入token的sql执行失败：%s\n", err)
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf < 0 {
			log.Printf("写入token的sql执行RowsAffect < 0 ：%s\n", err)
			return false
		}
	}
	return true
}
