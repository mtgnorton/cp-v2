package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Login = login{}

type login struct {
}

func (n *login) LoginPage(ctx context.Context, req *define.LoginPageReq) (res *define.LoginPageRes, err error) {

	service.View().Render(ctx, define.View{
		Title:    "登录",
		Template: "user/login.html",
		Data: g.Map{
			"captchaInfo": model.CommonGenerateCaptchaOutput{
				CaptchaId:     "",
				CaptchaBase64: "",
			},
		},
	})
	return
}

func (n *login) Login(ctx context.Context, req *define.LoginReq) (res *define.LoginRes, err error) {
	res = &define.LoginRes{}
	res.LoginOutput, err = service.UserNoNeedLogin.Login(ctx, req.LoginInput)
	return
}
