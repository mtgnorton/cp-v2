package controller

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var UserOnline = userOnline{}

type userOnline struct {
}

func (u *userOnline) OnlinePage(ctx context.Context, req *define.UserOnlinePageReq) (res *define.UserOnlinePageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)

	if err != nil {
		return
	}

	service.View().Render(ctx, define.View{
		Title:    "实时在线用户",
		User:     user,
		Template: "online-users.html",
		Data:     g.Map{},
	})
	return
}
