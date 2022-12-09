package define

import "github.com/gogf/gf/v2/frame/g"

type ForgetPasswordPageReq struct {
	g.Meta `path:"/forget-password" method:"get" tags:"用户相关" summary:"忘记密码页面"`
}
type ForgetPasswordPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ForgetPasswordReq struct {
	g.Meta    `path:"/forget-password" method:"post" tags:"用户相关" summary:"忘记密码"`
	Username  string `p:"username" v:"required#用户名必须" dc:"用户名"`
	Email     string `p:"email" v:"required|email#邮箱必须｜邮箱格式错误" dc:"邮箱"`
	Code      string `p:"captcha_code"  v:"" dc:"验证码"`
	CaptchaId string `p:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type ForgetPasswordRes struct {
}
type ResetPasswordPageReq struct {
	g.Meta   `path:"/reset-password/:username/:email/:time/:proof" method:"get" tags:"用户相关" summary:"重置密码页面"`
	Username string
	Email    string
	Time     string
	Proof    string
}
type ResetPasswordPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
type ResetPasswordReq struct {
	g.Meta    `path:"/reset-password" method:"post" tags:"用户相关" summary:"重置密码"`
	Username  string `p:"username" v:"required#用户名必须" dc:"用户名"`
	Email     string `p:"email" v:"required|email#邮箱必须｜邮箱格式错误" dc:"邮箱"`
	Time      int64  `p:"time" v:"required#时间戳必须" dc:"时间戳"`
	Proof     string `p:"proof" v:"required#校验码必须" dc:"校验码"`
	Password  string `v:"required|password#请输入密码|密码长度需要在长度在6~18之间" dc:"密码" d:"password" json:"password"`
	Password2 string `v:"required|same:password#请输入确认密码|两次密码不一致" dc:"确认密码" d:"password2" json:"password2"`
}
type ResetPasswordRes struct {
}
