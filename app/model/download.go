package model

type DownloadImageInput struct {
	Url          string
	Dir          string //存储目录,可以为空
	SaveFilename string //如果为空则自动生成文件名
}
type DownloadImageOutput struct {
	Filename     string
	Url          string // 拼接域名和relativePath
	RelativePath string
	AbsolutePath string
}
