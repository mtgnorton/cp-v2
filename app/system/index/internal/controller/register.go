package controller

import (
	context "context"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var Register = register{}

type register struct {
}

// RegisterPage 注册页面
func (r *register) RegisterPage(ctx context.Context, req *define.RegisterPageReq) (res *define.RegisterPageRes, err error) {

	service.View().Render(ctx, define.View{
		Title:    "注册",
		Template: "user/register.html",
		Data: g.Map{
			"captchaInfo": model.CommonGenerateCaptchaOutput{
				CaptchaId:     "",
				CaptchaBase64: "",
			},
		},
	})
	return
}

// Register 注册请求
func (n *register) Register(ctx context.Context, req *define.RegisterReq) (res *define.RegisterRes, err error) {
	res = &define.RegisterRes{}
	err = service.UserNoNeedLogin.Register(ctx, req.RegisterInput)

	return
}

// ActivePage 激活用户页面
func (n *register) ActivePage(ctx context.Context, req *define.ActivePageReq) (res *define.ActivePageRes, err error) {

	err = service.User.Active(ctx, req)
	err = nil
	prompt := fmt.Sprintf(`%s,您已经完成注册，现在可以<a href='/login-page'>登录</a> 了。`, req.Username)
	if err != nil {
		g.Log().Error(ctx, err)
		prompt = gerror.Current(err).Error()
	}

	service.View().Render(ctx, define.View{
		Title:    "激活",
		Template: "prompt.html",
		Data: g.Map{
			"prompt": prompt,
		},
	})
	return
}

// ResendActiveEmail  重新发送激活邮件页面
func (n *other) ResendActiveEmailHtml(ctx context.Context, req *define.ResendActiveEmailPageReq) (res *define.ResendActiveEmailPageRes, err error) {
	service.View().Render(ctx, define.View{
		Title:    "重新发送激活邮件",
		Template: "user/resend-active-email.html",
	})
	return
}

// ResendActiveEmail 重新发送激活邮件
func (n *other) ResendActiveEmail(ctx context.Context, req *define.ResendActiveEmailReq) (res *define.ResendActiveEmailRes, err error) {
	err = service.UserNoNeedLogin.ResendActiveEmail(ctx, req)
	return
}
