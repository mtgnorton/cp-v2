package shared

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/utility"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"
)

var Upload = upload{}

type upload struct {
}

// Single 上传单个文件
//todo  gf/v2@v2.1.4/net/ghttp/ghttp_server_config.go:239 有关上传大小的限制为1MB
func (u *upload) Single(ctx context.Context, in *model.UploadInput) (out *model.UploadOutput, err error) {
	out = &model.UploadOutput{}

	if in.File == nil {
		return nil, gerror.New("上传文件不存在")
	}
	if in.UploadPosition == "" {
		return nil, gerror.New("上传位置不能为空")
	}
	mimeType, err := u.GetMimeType(ctx, in.File)
	if err != nil {
		return
	}
	var fileExt string
	if fileExt = model.AllImages[mimeType]; mimeType == "" || fileExt == "" {
		return nil, gerror.New("上传文件类型错误")
	}

	// 上传大小限制
	if in.LimitSize != "" && in.File.Size > gfile.StrToSize(in.LimitSize) {
		return nil, gerror.New("上传文件大小超过限制")
	}

	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		uploadPath = "upload"
	}
	uploadPath = gstr.Trim(uploadPath, "/")

	rootPath := utility.GetServerPath()

	if in.SaveFilename == "" {
		in.File.Filename = u.randomFilename(ctx, in.File)
	} else {
		in.File.Filename = in.SaveFilename + gfile.Ext(in.File.Filename)
	}
	out.AbsolutePath = gfile.Join(rootPath, uploadPath, string(in.UploadPosition), in.Dir)
	out.Filename, err = in.File.Save(out.AbsolutePath, false)

	out.AbsolutePath = gfile.Join(out.AbsolutePath, out.Filename)
	out.RelativePath = "/" + gfile.Join(uploadPath, string(in.UploadPosition), in.Dir, out.Filename)

	// 从ctx获取到request
	r := g.RequestFromCtx(ctx)

	schema := "http://"

	if r.TLS != nil {
		schema = "https://"
	}
	out.Url = schema + r.Host + out.RelativePath

	out.MimeType = mimeType

	out.Size = gfile.FormatSize(in.File.Size)

	g.Dump(out, "UploadOutput")

	return
}

func (u *upload) randomFilename(ctx context.Context, file *ghttp.UploadFile) string {
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	name = name + gfile.Ext(file.Filename)
	return name
}

func (u *upload) GetMimeType(ctx context.Context, file *ghttp.UploadFile) (string, error) {
	fileHeader := make([]byte, 512)

	originFile, err := file.Open()
	if err != nil {
		return "", err
	}
	// Copy the headers into the FileHeader buffer
	if _, err := originFile.Read(fileHeader); err != nil {
		return "", err
	}

	return http.DetectContentType(fileHeader), nil
}
