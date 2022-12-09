package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RegisterPageReq struct {
	g.Meta `path:"/register-page" method:"get" tags:"用户相关" summary:"注册"`
}

type RegisterPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type RegisterInput struct {
	*model.UserRegisterInput
	Code      string `p:"captcha_code"  v:"" dc:"验证码"`
	CaptchaId string `p:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" summary:"注册" tags:"用户相关"`
	*RegisterInput
}

type RegisterRes struct {
}

type ResendActiveEmailPageReq struct {
	g.Meta `path:"/resend-active-email-page" method:"get" tags:"用户相关" summary:"重新发送激活邮件页面"`
}
type ResendActiveEmailPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ActivePageReq struct {
	g.Meta   `path:"/register/verify/:username/:time/:proof" method:"get" tags:"用户相关" summary:"激活"`
	Username string
	Proof    string
	Time     int64
}

type ActivePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ResendActiveEmailReq struct {
	g.Meta    `path:"/resend-active-email" method:"post" tags:"用户相关" summary:"重新发送激活邮件"`
	Email     string `p:"email" v:"required|email|max-length:50#邮箱必须｜邮箱格式错误｜邮箱长度超过限制" dc:"邮箱"`
	Code      string `p:"captcha_code"  v:"" dc:"验证码"`
	CaptchaId string `p:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type ResendActiveEmailRes struct {
}
