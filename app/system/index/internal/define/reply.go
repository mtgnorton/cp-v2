package define

import "github.com/gogf/gf/v2/frame/g"

type ReplyStoreReq struct {
	g.Meta  `path:"/comments-store" method:"post" tags:"新增回复" summary:"新增回复"`
	PostId  uint   `v:"required#主题id不能为空" json:"post_id" dc:"主题id"`
	Content string `v:"required#内容不能为空" json:"content" dc:"回复内容"`
}

type ReplyStoreRes struct {
}
