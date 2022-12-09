package shared

import (
	"gf-admin/app/model"
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestDownload_Image(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		out, err := Download.Image(gctx.New(), &model.DownloadImageInput{
			Url: "https://cdn.v2ex.com/navatar/3b5d/ca50/819_xlarge.png?m=1633251402",
		})
		g.Dump(out, err)
	})
}
