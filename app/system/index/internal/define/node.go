package define

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NodePageReq struct {
	g.Meta  `path:"/node-page/{keyword}/*page" method:"get" tags:"节点" summary:"节点页面"`
	Keyword string `json:"keyword"`
	Page    int    `json:"page" d:"1"`
}

type NodePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type NodeListPageReq struct {
	g.Meta `path:"/node-list-page" method:"get" tags:"节点" summary:"节点列表页面"`
}

type NodeListPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
