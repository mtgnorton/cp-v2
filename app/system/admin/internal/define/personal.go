package define

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type PersonalLoginInfoReq struct {
	g.Meta `path:"/login-info" method:"GET" summary:"相关登录信息" tags:"全局"`
}

type PersonalLoginInfoOutput struct {
	Data map[string]*gvar.Var `json:"data"`
}
type PersonalLoginInfoRes struct {
	*PersonalLoginInfoOutput
}

type PersonalLoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"执行登录请求" tags:"全局"`
	PersonalLoginInput
}
type PersonalLoginRes struct {
	*PersonalLoginOutput
}

type PersonalLoginInput struct {
	Username string `json:"username" v:"required#请输入账号"   dc:"账号"`

	Password  string `json:"password" v:"required#请输入密码"   dc:"密码(明文)" d:"admin"`
	Code      string `json:"code"  v:"" dc:"验证码"`
	CaptchaId string `json:"captcha_id" v:"" dc:"后端返回的captcha标识符"`
}

type PersonalLoginOutput struct {
	Token string
}

type PersonalInfoReq struct {
	g.Meta `path:"/personal-info"  method:"get" summary:"获取管理员个人信息" tags:"个人中心"`
}

type PersonalInfoOutput struct {
	Id        uint        `json:"id"`
	Username  string      `json:"username"`
	Status    string      `d:"normal" dc:"管理员状态，可选值：normal,disabled" json:"status"`
	Nickname  string      `dc:"昵称" d:"nickname" json:"nickname"`
	Avatar    string      `dc:"管理员头像地址" d:"http://www.baidu.com" json:"avatar"`
	CreatedAt *gtime.Time `json:"created_at"`
}

type PersonalInfoRes struct {
	PersonalInfoOutput
}

type PersonalUpdateInput struct {
	OldPassword       string `json:"old_password" v:"required#旧密码必须"`
	Password          string `json:"password"      v:"password|same:password_confirmed#新密码长度需要在6-18之间｜两次密码输入不一致"  ` // MD5密码
	PasswordConfirmed string `json:"password_confirmed"`
	Nickname          string `json:"nickname"        ` // 昵称
	Avatar            string `json:"avatar"          ` // 头像地址
}

type PersonalUpdateReq struct {
	g.Meta `path:"/personal-update" method:"put" summary:"更新管理员的信息" tags:"个人中心"`
	PersonalUpdateInput
}

type PersonalUpdateRes struct {
}

type PersonalAvatarReq struct {
	g.Meta     `path:"/personal-avatar" method:"post" summary:"上传头像" tags:"个人中心"`
	AvatarFile *ghttp.UploadFile `json:"avatar_file" type:"file" v:"required#请上传头像文件"`
}

type PersonAvatarInput struct {
	AvatarFile *ghttp.UploadFile `json:"avatar_file"`
}

type PersonAvatarOutput struct {
	AvatarUrl string `json:"avatar_url"`
}
type PersonAvatarRes struct {
	PersonAvatarOutput
}

type WsReq struct {
	g.Meta `path:"/ws" method:"get" summary:"websocket通信" tags:"全局"`
}

type WsRes struct {
}
