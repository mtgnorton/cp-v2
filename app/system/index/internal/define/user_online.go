package define

import "github.com/gogf/gf/v2/frame/g"

type UserOnlinePageReq struct {
	g.Meta `path:"/online-users" method:"get" summary:"实时在线用户页面" tags:"实时在线用户页面"`
}

type UserOnlinePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
