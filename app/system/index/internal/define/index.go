package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页" dc:"首页页面显示的节点默认会显示上次点击的节点，如果没有点击过则显示首页的第一个节点,确定好节点后会跳转到首页节点页面"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type IndexNodeReq struct {
	g.Meta  `path:"/k/{keyword}" method:"get" tags:"首页" summary:"首页节点"`
	Keyword string `in:"path" name:"keyword" default:"" desc:"关键字"`
}

type IndexNodeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AppletIndexReq struct {
	g.Meta  `path:"/applet/index" method:"get" tags:"小程序首页" summary:"小程序首页"`
	Keyword string `in:"path" name:"keyword" default:"" desc:"关键字"`
	Page    string `in:"path" name:"page" default:"1" desc:"页码"`
}

type AppletIndexRes struct {
	NodeList       *model.NodeListOutput `json:"node_list"`
	PostList       *model.PostListOutput `json:"post_list"`
	CurrentKeyword string                `json:"current_keyword"`
}

type SearchReq struct {
	g.Meta `path:"/search/:type/:search/:page" method:"get" tags:"全局" summary:"搜索"`
	Search string `in:"path" name:"search" default:"" desc:"搜索关键字"`
	Type   string `in:"path" name:"type" default:"post" desc:"搜索类型"`
	Page   int    `in:"path" name:"page" default:"1" desc:"页码"`
}

type SearchRes struct {
}
