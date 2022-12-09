package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/encoding/gurl"

	"github.com/gogf/gf/v2/frame/g"
)

var Other = other{}

type other struct {
}

func (l *other) Prompt(ctx context.Context, req *define.PromptPageReq) (res *define.PromptPageRes, err error) {

	redirectUrl, err := gurl.Decode(req.RedirectUrl)
	if err != nil {
		return
	}
	redirectUrl = gstr.Replace(redirectUrl, "\\", "/")

	service.View().Render(ctx, define.View{
		Title:    "提示",
		Template: "prompt.html",
		Data: g.Map{
			"prompt":      req.Message,
			"redirectUrl": redirectUrl,
		},
	})
	return
}

// UpdateEmailActivePage 更新邮箱激活页面
func (n *other) UpdateEmailActivePage(ctx context.Context, req *define.UserUpdateEmailActivePageReq) (res *define.UerUpdateEmailActivePageRes, err error) {
	res = &define.UerUpdateEmailActivePageRes{}
	err = service.User.UpdateEmailActive(ctx, req)

	if err == nil {
		shared.Common.Prompt(ctx, "邮箱修改成功", "/")
	}

	return
}
