package controller

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var ForgetPassword = forgetPassword{}

type forgetPassword struct {
}

func (n *forgetPassword) ForgetPasswordHtml(ctx context.Context, req *define.ForgetPasswordPageReq) (res *define.ForgetPasswordPageRes, err error) {
	service.View().Render(ctx, define.View{
		Title:    "忘记密码",
		Template: "user/forget-password.html",
	})
	return
}

func (n *forgetPassword) ForgetPassword(ctx context.Context, req *define.ForgetPasswordReq) (res *define.ForgetPasswordRes, err error) {

	err = service.UserNoNeedLogin.ForgetPassword(ctx, req)
	return
}

func (n *forgetPassword) ResetPasswordHtml(ctx context.Context, req *define.ResetPasswordPageReq) (res *define.ResetPasswordPageRes, err error) {
	service.View().Render(ctx, define.View{
		Title:    "重置密码",
		Template: "user/reset-password.html",
		Data: g.Map{
			"req": req,
		},
	})
	return
}

func (n *forgetPassword) ResetPassword(ctx context.Context, req *define.ResetPasswordReq) (res *define.ResetPasswordRes, err error) {

	err = service.UserNoNeedLogin.ResetPassword(ctx, req)

	return
}
