package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PromptPageReq struct {
	g.Meta      `path:"/prompt/:message/*redirectUrl" method:"get" tags:"全局" summary:"提示" dc:"提示信息页面"`
	Message     string `json:"message" dc:"提示信息"`
	RedirectUrl string `json:"redirect_url" dc:"显示完提示后，重定向的地址,如果为空则返回上一页"`
}
type PromptPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserUpdateEmailActivePageReq struct {
	g.Meta   `path:"/user/update/email-active/:username/:email/:time/:proof" method:"get" tags:"个人中心" summary:"更新邮箱激活页面"`
	Username string
	Email    string
	Time     int64
	Proof    string
}

type UerUpdateEmailActivePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AuthInfoReq struct {
	g.Meta `path:"/auth-info" method:"get" summary:"获取用户信息" tags:"用户相关"`
}

type AuthInfoRes struct {
	*AuthInfoOutput
}

type AuthInfoOutput struct {
	model.UserInfoWithoutPass
}
