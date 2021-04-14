package handler

import (
	"cloudstorage/v1/db"
	"cloudstorage/v1/utils"
	"log"
	"net/http"
	"strconv"
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

// FileFastUpload 文件秒传功能
func FileFastUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		filesha1 := r.Form.Get("sha1")
		filesize, err := strconv.Atoi(r.Form.Get("filesize"))
		if err != nil {
			log.Printf("filesize convert to int64 failed: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		filename := r.Form.Get("filename")

		//  查询 filesha1 文件是否存在
		file, err := db.TabFileDataQuery(filesha1)
		if err != nil {
			log.Printf("Query Failed: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := utils.RespMsg{Code: -1, Msg: "文件妙传失败，请使用普通上传功能"}
		if file == nil { // 文件不存返回秒传失败
			log.Printf("请使用普通上传功能")
			w.Write(res.JSONBytes())
			return
		}
		// 文件秒传
		ok := db.TabUserFileInsert(username, filesha1, filename, int64(filesize))
		if !ok {
			log.Printf("文件妙传，插入用户文件表中数据失败")
			res.Msg = "文件妙传失败，请重试"
			w.Write(res.JSONBytes())
			return
		}

		res.Code = 0
		res.Msg = "文件秒传成功"
		w.Write(res.JSONBytes())
	}
}
