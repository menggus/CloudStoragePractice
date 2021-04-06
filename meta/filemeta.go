package meta

import (
	db "cloudstorage/v1/db"
	"log"
)

type FileMeta struct {
	FileSha1   string
	FileName   string
	FileSize   int64
	FilePath   string
	UploadTime string
}

var fileMetas map[string]FileMeta // 存储所有上传文件的元信息

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMetas 新增or更新文件元信息
func UpdateFileMetas(f FileMeta) {
	fileMetas[f.FileSha1] = f
}

// UpdateFileMetasDB 新增or更新文件至mysql数据库
func UpdateFileMetasDB(f FileMeta) bool {

	return db.TabFileDataInsert(f.FileSha1, f.FileName, f.FileSize, f.FilePath)
}

// GetFileMeta 通过sha1值获取文件元信息
func GetFileMeta(sha string) FileMeta {
	return fileMetas[sha]
}

// GetFileMetaDB 从数据库获取文件元信息
func GetFileMetaDB(filesha1 string) FileMeta {
	file, err := db.TabFileDataQuery(filesha1)
	if err != nil {
		log.Printf("Query Failed: %s\n", err)
		return FileMeta{}
	}
	filemeta := FileMeta{
		FileSha1: file.FileSha1,
		FileName: file.FileName.String,
		FilePath: file.FileAddr.String,
		FileSize: file.FileSize.Int64,
	}

	return filemeta
}

// DeletaFileMeta 删除文件
func DeletaFileMeta(sha string) bool {

	delete(fileMetas, sha)

	return true
}
