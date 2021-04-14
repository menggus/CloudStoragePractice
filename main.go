package main

import (
	"cloudstorage/v1/handler"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/file", handler.FileHandler)
	http.HandleFunc("/msg/succed", handler.SuccedHandler)
	http.HandleFunc("/file/meta", handler.QueryFileInfoHandler)
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
	http.HandleFunc("/file/rname", handler.RenameHandler)
	http.HandleFunc("/file/delete", handler.DeleteFileHandler)
	http.HandleFunc("/user/signup", handler.UserRegisterHandler)
	http.HandleFunc("/user/signin", handler.UserLoginHandler)
	http.HandleFunc("/user/home", handler.UserHomeHandler)
	http.HandleFunc("/user/info", handler.TokenHandler(handler.UserInfoHandler))
	http.HandleFunc("/file/query", handler.TokenHandler(handler.FileDataQuery))
	http.HandleFunc("/file/fastupload", handler.TokenHandler(handler.FileFastUpload))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err)
	}
}
