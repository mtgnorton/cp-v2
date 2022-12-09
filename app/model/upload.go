package model

import "github.com/gogf/gf/v2/net/ghttp"

type UploadPosition string

const (
	UPLOAD_POSITION_BACKEND UploadPosition = "backend"
	UPLOAD_POSITION_FRONTED UploadPosition = "frontend"
)

type UploadType string

var (
	// 类型来源 1.18/src/net/http/sniff_test.go:44
	// "image/png", "image/jpeg", "image/gif", "image/bmp", "image/webp",
	AllImages = map[string]string{
		"image/png":  "png",
		"image/jpeg": "jpg",
		"image/gif":  "gif",
		"image/bmp":  "bmp",
		"image/webp": "webp",
	}
)

const (
	UPLOAD_TYPE_IMAGE UploadType = "image"
)

// 存储路径
type UploadInput struct {
	File           *ghttp.UploadFile
	UploadPosition UploadPosition //前台还是后台
	Dir            string         //存储目录,可以为空
	UploadType     UploadType     //为空不限制
	LimitSize      string         //为空不限制,限制大小 支持人性化的单位，如kb，mb，gb，见 gfile.StrToSize
	SaveFilename   string         //如果为空则自动生成文件名

}

type UploadOutput struct {
	Filename     string
	AbsolutePath string
	RelativePath string
	Url          string
	MimeType     string
	Size         string
}
