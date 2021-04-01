package meta

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

// GetFileMeta 通过sha1值获取文件元信息
func GetFileMeta(sha string) FileMeta {
	return fileMetas[sha]
}
