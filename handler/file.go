package handler

import (
	"cloudstorage/v1/db"
	"log"
	"net/http"
)

// FileDataQuery 文件接口
func FileDataQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		token := r.FormValue("token")
		// 校验token
		ok := db.IsValidateToken(username, token)
		if !ok {
			log.Println("token validate failed")
			w.Write([]byte("token validate failed"))
			return
		}

		// 获取用户文件信息

	}
}
