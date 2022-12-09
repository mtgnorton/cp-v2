package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ReplyListReq struct {
	g.Meta `path:"/reply-list" method:"get" summary:"回复列表" tags:"回复管理"`
	*model.ReplyListInput
}

type ReplyListRes struct {
	*model.ReplyListOutput
}

type ReplyAuditReq struct {
	g.Meta `path:"/reply-audit" method:"post" summary:"审核回复" tags:"回复管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要审核的回复"`
}
type ReplyAuditRes struct {
}

type ReplyDestroyReq struct {
	g.Meta `path:"/reply-destroy" method:"delete" summary:"删除回复" tags:"回复管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要删除的回复"`
}
type ReplyDestroyRes struct {
}

type ReplyUpdateReq struct {
	g.Meta `path:"/reply-update" method:"put" summary:"更新回复" tags:"回复管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要更新的回复"`
	Status int  `json:"status" dc:"status,1 已审核 -1屏蔽"`
}
type ReplyUpdateRes struct {
}
