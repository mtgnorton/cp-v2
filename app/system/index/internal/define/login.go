package define

import "github.com/gogf/gf/v2/frame/g"

type LoginPageReq struct {
	g.Meta `path:"/login-page" method:"get" tags:"用户相关" summary:"登录"`
}

type LoginPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录" tags:"用户相关"`
	*LoginInput
}

type LoginInput struct {
	Username string `v:"required#请输入用户名" dc:"用户名" d:"username" json:"username"`
	Password string `v:"required#请输入密码" dc:"密码" d:"password" json:"password"`

	Code      string `p:"captcha_code"  v:"" dc:"验证码"`
	CaptchaId string `p:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

type LoginRes struct {
	*LoginOutput
}
