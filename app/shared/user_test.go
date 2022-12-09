package shared

import (
	"gf-admin/utility"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestUser_Email(t *testing.T) {
	g.View().BindFunc("TimeFormatDivide24Hour", utility.TimeFormatDivide24Hour)
	g.View().BindFunc("InSlice", utility.InSlice)
	g.View().BindFunc("toMap", utility.ToTemplateMap)
	gtest.C(t, func(t *gtest.T) {

		// 获取邮件模板 模板路径 app/system/index/internal/template/email.html
		c, err := g.View().Parse(ctx, "email.html", g.Map{
			"siteName":  "gf-admin",
			"username":  "gf-admin",
			"verifyUrl": "http://localhost:8080",
			"logo":      "http://localhost:8080",
		})
		g.Dump(c, err)
	})
}

func TestUser_SendActiveEmail(t *testing.T) {
	g.View().BindFunc("TimeFormatDivide24Hour", utility.TimeFormatDivide24Hour)
	g.View().BindFunc("InSlice", utility.InSlice)
	g.View().BindFunc("toMap", utility.ToTemplateMap)
	gtest.C(t, func(t *gtest.T) {
		err := User.SendActiveEmail(gctx.New(), 5)
		g.Dump(err)
	})
}
