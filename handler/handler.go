/*
	web接口
*/
package handler

import (
	"cloudstorage/v1/meta"
	"cloudstorage/v1/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// FileUploadHandler 文件接口
func FileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// GET 返回文件上传 page
		// ioutil 包读取文件
		page, err := ioutil.ReadFile("static/view/index.html")
		if err != nil {
			log.Fatal(err)
			io.WriteString(w, "访问的页面不触存在")
			return
		}
		io.WriteString(w, string(page))

	} else if r.Method == "POST" {
		// POST 接收文件存放道本地
		ff, head, err := r.FormFile("file")
		if err != nil {
			log.Printf("Faild recive file data: %s\n", err)
			return
		}
		defer ff.Close()
		// 上传文件元信息
		fileMeta := meta.FileMeta{
			FileName:   head.Filename,
			FilePath:   "tmp/" + head.Filename,
			UploadTime: time.Now().Format("2006-01-02 15:04:05"),
		}

		nf, err := os.Create(fileMeta.FilePath)
		if err != nil {
			log.Printf("Failed create local new file: %s\n", err)
			return
		}
		defer nf.Close()
		// 上传文件大小
		fileMeta.FileSize, err = io.Copy(nf, ff)
		if err != nil {
			log.Printf("Failed save upload file: %s\n", err)
			return
		}
		// 上传文件的sha1值
		nf.Seek(0, 0) // 游标重新回到文件头部
		fileMeta.FileSha1 = utils.FileSha1(nf)

		// 上传存储元信息
		meta.UpdateFileMetas(fileMeta)

		http.Redirect(w, r, "/msg/succed", http.StatusFound)
	}
}

func SuccedHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "文件上传成功")
}

// QueryFileInfoHandler 查询 file 信息接口
// url： /file/meta?sha1=3bc5f45eb1cf75eff7f3e56c514748a11e84cdba
func QueryFileInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sha1 := r.FormValue("sha1")
		filemeta := meta.GetFileMeta(sha1)

		fj, err := json.Marshal(filemeta)
		if err != nil {
			log.Printf("Failed change to json: %s\n", err)
			return
		}
		io.WriteString(w, string(fj))
	}
}

// DownloadFileHandler 下载文件接口
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		sha1 := r.FormValue("sha1")
		filemeta := meta.GetFileMeta(sha1)

		data, err := ioutil.ReadFile(filemeta.FilePath)
		if err != nil {
			log.Printf("File Not Found: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 返回文件数据设置header
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("content-disposition", fmt.Sprintf("attachment;filename=\"%s\"", filemeta.FileName))
		w.Write(data)
	}
}
