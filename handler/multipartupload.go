package handler

import (
	"cloudstorage/v1/cache"
	"cloudstorage/v1/utils"
	"context"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type MutilPartUploadInfo struct {
	FileSha1     string
	FileUploadID string
	FileSize     int
	PartSize     int
	PartCount    int
}

// InitMultipartUploadHandler 初始化分块上传
func InitMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析http param
	r.ParseForm()
	username := r.Form.Get("username")
	sha1 := r.Form.Get("sha1")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		log.Printf("upload file init failed: %s\n", err)
		w.Write(utils.NewRespMsg(-1, "invalid params", nil).JSONBytes())
		return
	}

	// 2.获取redis连接
	rds := cache.NewRedis()
	defer rds.Close()

	// 3.生成分块上传信息
	uploadinfo := MutilPartUploadInfo{
		FileSha1:     sha1,
		FileUploadID: username + strconv.FormatInt(time.Now().UnixNano(), 10),
		FileSize:     filesize,
		PartSize:     utils.MutilpartSize,
		PartCount:    int(math.Ceil(float64(filesize) / (utils.MutilpartSize))),
	}

	// 4.写入分块信息到redis
	ctx := context.Background()
	//rds.Do(ctx, "HSET", "mp_"+uploadinfo.FileUploadID, "PartCount", uploadinfo.PartCount)
	//rds.Do(ctx, "HSET", "mp_"+uploadinfo.FileUploadID, "FileSha1", uploadinfo.FileSha1)
	//rds.Do(ctx, "HSET", "mp_"+uploadinfo.FileUploadID, "FileSize", uploadinfo.FileSize)
	// HSET 4.0 以上新版本可以一次性设置多个值
	rds.Do(ctx, "HSET", "mp_"+uploadinfo.FileUploadID, "PartCount", uploadinfo.PartCount, "FileSha1",
		uploadinfo.FileSha1, "FileSize", uploadinfo.FileSize)
	// 5.返回响应
	w.Write(utils.NewRespMsg(1, "init multipart upload succeed", uploadinfo).JSONBytes())
}
