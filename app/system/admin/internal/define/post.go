package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/frame/g"
)

type PostListReq struct {
	g.Meta `path:"/post-list" method:"get" summary:"主题列表" tags:"主题管理"`
	Id     uint
	*model.PostListInput
	model.OrderFieldDirectionInput
}

type PostListRes struct {
	*model.PostListOutput
	FrontDomain string `json:"front_domain"`
}

type PostDestroyReq struct {
	g.Meta `path:"/post-destroy" method:"delete" summary:"删除主题" tags:"主题管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要删除的主题"`
}

type PostDestroyRes struct {
}

type PostAuditReq struct {
	g.Meta `path:"/post-audit" method:"post" summary:"审核主题" tags:"主题管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要审核的主题"`
}

type PostAuditRes struct {
}

type PostToggleTopInput struct {
	Id      uint       `json:"id"`
	EndTime gtime.Time `json:"end_time" dc:"置顶截止时间,为空时，代表取消置顶"`
}

type PostToggleTopReq struct {
	g.Meta `path:"/post-toggle-top" method:"put" summary:"置顶主题" tags:"主题管理"`
	*PostToggleTopInput
}

type PostToggleTopRes struct {
}

type PostInfoReq struct {
	g.Meta `path:"/post-info" method:"get" summary:"主题详情" tags:"主题管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要查看的主题"`
}
type PostInfoRes struct {
	entity.Posts
}

type PostUpdateReq struct {
	g.Meta     `path:"/post-update" method:"put" summary:"修改主题" tags:"主题管理"`
	Id         uint       `json:"id" dc:"id" v:"min:1#请选择需要修改的主题"`
	Status     int        `json:"status" dc:"状态" `
	TopEndTime gtime.Time `json:"top_end_time" dc:"置顶截止时间,为空时，代表取消置顶"`
}

type PostUpdateRes struct {
}
