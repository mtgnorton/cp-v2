package shared

import (
	"gf-admin/app/model"
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestNode_List(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		res, err := Node.List(gctx.New(), &model.NodeListInput{
			NeedChildren: false,
		})
		g.Dump(res, err)
	})
	gtest.C(t, func(t *gtest.T) {
		res, err := Node.List(gctx.New(), &model.NodeListInput{
			NeedChildren: true,
		})
		g.Dump(res, err)
	})
}
