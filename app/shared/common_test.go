package shared

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/frame/g"
)

func TestCommon_TransformReferLink(t *testing.T) {
	r := Common.TransformReferLink("hello @gf,@sss 你好")
	g.Dump(r)

}

func TestCommon_TruncateDatabase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := Common.TruncateDatabase(gctx.New())
		t.AssertNil(err)
	})
}
