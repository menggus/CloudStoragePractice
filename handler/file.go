package handler

import (
	"cloudstorage/v1/db"
	"cloudstorage/v1/utils"
	"net/http"
)

// FileDataQuery 文件接口
func FileDataQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		// 校验token
		// 获取用户文件信息
		data, err := db.TabUserFileQueryRows(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := utils.RespMsg{
			Code: 0,
			Msg:  "succeed",
			Data: data,
		}

		w.Write(res.JSONBytes())
	}
}
