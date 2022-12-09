package api

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
)

var User = user{}

type user struct {
}

func (u *user) Info(ctx context.Context, req *define.UserInfoReq) (res *define.UserInfoRes, err error) {
	res = &define.UserInfoRes{}
	res.UserSummary, err = service.FrontTokenInstance.GetUser(ctx)

	return
}

// Logout 退出登录
func (p *user) Logout(ctx context.Context, req *define.AppletLogoutReq) (res *define.AppletLogoutRes, err error) {
	res = &define.AppletLogoutRes{}
	err = service.User.Logout(ctx)
	return
}
