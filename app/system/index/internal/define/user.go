package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserIndexReq struct {
	g.Meta   `path:"/user/:username/*type" method:"get" tags:"个人中心" summary:"个人中心页面"`
	Username string
	Type     string `json:"type" d:"posts" dc:"posts为最近主题页面,replies为最近回复页面" `
}

type UserIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type CollectNodeItem struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Keyword       string `json:"keyword"`
	Img           string `json:"img"`
	CollectAmount int    `json:"collect_amount"`
}

type UserBalanceLogReq struct {
	g.Meta `path:"/user/balance-log/*page" method:"get" tags:"个人中心" summary:"余额记录页面"`
	Page   int
}

type UserBalanceLogRes struct {
	List []*entity.BalanceChangeLog `json:"list"`
}

type UserMessagePageReq struct {
	g.Meta `path:"/user/message/*page" method:"get" tags:"个人中心" summary:"消息页面"`
	Page   int
}

type UserMessagePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserInfoReq struct {
	g.Meta `path:"/applet/user/info" method:"get" tags:"个人中心" summary:"个人信息"`
}
type UserInfoRes struct {
	*model.UserSummary
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"get" summary:"退出登录" tags:"用户相关"`
}

type LogoutRes struct {
}

type AppletLogoutReq struct {
	g.Meta `path:"/applet/logout" method:"get" summary:"退出登录" tags:"用户相关"`
}

type AppletLogoutRes struct {
}
