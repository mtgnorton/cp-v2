package shared

import (
	"context"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/utility"
	"strconv"
	"strings"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

var Download = download{}

type download struct {
}

// Image 下载远程图片
func (d *download) Image(ctx context.Context, in *model.DownloadImageInput) (out *model.DownloadImageOutput, err error) {
	out = &model.DownloadImageOutput{}
	r, err := g.Client().Get(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = r.Close()
		if err != nil {
			g.Log().Warning(ctx, fmt.Sprintf("下载远程图片出错，%s", err))
		}
	}()

	if in.SaveFilename == "" {
		name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
		in.SaveFilename = name + gfile.Ext(in.Url)
	}
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		uploadPath = "upload"
	}
	uploadPath = gstr.Trim(uploadPath, "/")

	rootPath := utility.GetServerPath()

	g.Dump(rootPath, uploadPath, "backend", in.Dir, in.SaveFilename)
	out.AbsolutePath = gfile.Join(rootPath, uploadPath, "backend", in.Dir, in.SaveFilename)

	out.RelativePath = "/" + gfile.Join(uploadPath, "backend", in.Dir, in.SaveFilename)

	out.Filename = in.SaveFilename

	rs := g.RequestFromCtx(ctx)

	schema := "http://"

	if rs != nil && rs.TLS != nil {
		schema = "https://"
		out.Url = schema + rs.Host + out.RelativePath
	}

	err = gfile.PutBytes(out.AbsolutePath, r.ReadAll())

	return
}
