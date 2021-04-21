package handler

import (
	"cloudstorage/v1/cache"
	"cloudstorage/v1/utils"
	"context"
	"io"
	"log"
	"math"
	"net/http"
	"os"
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

// UploadPartHandler 上传分块
func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析参数
	r.ParseForm()
	// username := r.Form.Get("username")
	id := r.Form.Get("uploadID") // 文件上传 文件id
	index := r.Form.Get("index") // 分块 id

	// 2.获取redis连接
	rds := cache.NewRedis()
	if rds == nil {
		log.Println("get redis server failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rds.Close()

	// 3.获取文件句柄，用于存储 分块内容
	fd, err := os.OpenFile("/data/"+id+"/"+index, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		log.Printf("get file object failed： %s\n", err)
		w.Write(utils.NewRespMsg(-1, "multipart upload failed", nil).JSONBytes())
		return
	}
	defer fd.Close()

	// 4.创建 buff 来缓存从客户端获取数据. 每次读取 1M
	buff := make([]byte, 1024*1024)

	// todo 这里少了一项文件数据MD5的检测，客户端生成的MD5值，与写入后的文件MD5值进行校验
	for {
		n, err := r.Body.Read(buff)

		if err != nil && err != io.EOF {
			log.Printf("read buff failed: %s\n", err)
			w.Write(utils.NewRespMsg(-1, "upload multi part failed", nil).JSONBytes())
			return
		}
		if n == 0 {
			break
		}
		fd.Write(buff[:n])
	}

	// 5. 写入redis记录分块上传状态
	ctx := context.Background()
	rds.Do(ctx, "HSET", "mp_"+id, "chkidx_"+index, 1)

	// 6. 返回消息
	w.Write(utils.NewRespMsg(1, "upload multi part success", nil).JSONBytes())
}

// MultipartMerageHandler 分块合并接口
//func MultipartMerageHandler(w http.ResponseWriter, r *http.Request) {
//	// 1.解析参数
//	r.ParseForm()
//	username := r.Form.Get("username")
//	id := r.Form.Get("uploadID")
//	filesha1 := r.Form.Get("sha1")
//	filesize := r.Form.Get("filesize")
//
//	// 2.获取redis连接
//	rds := cache.NewRedis()
//	if rds == nil {
//		log.Println("get redis server failed")
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	// 3. 获取分块数据，是否已经全部上传完成
//	ctx := context.Background()
//	res := rds.Do(ctx, "HGETALL", "mp_"+id)
//}
