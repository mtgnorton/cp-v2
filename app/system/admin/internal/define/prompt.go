package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PromptListReq struct {
	g.Meta `path:"/prompt-list" method:"get" tags:"提示管理" summary:"提示列表"`
	model.PageSizeInput
}
type PromptListRes struct {
	List []*entity.Prompts `json:"list"`
	model.PageSizeOutput
}

type PromptStoreReq struct {
	g.Meta      `path:"/prompt-store" method:"post" tags:"提示管理" summary:"保存提示"`
	Position    string `p:"position" v:"required#位置必须" dc:"位置"`
	Content     string `p:"content" v:"required#内容必须" dc:"内容"`
	Description string `p:"description"  dc:"简介"`
	IsDisabled  int    `  dc:"状态"`
}

type PromptStoreRes struct {
}

type PromptInfoReq struct {
	g.Meta `path:"/prompt-info" method:"get" tags:"提示管理" summary:"提示详情"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type PromptInfoRes struct {
	entity.Prompts
}

type PromptUpdateReq struct {
	g.Meta      `path:"/prompt-update" method:"put" tags:"提示管理" summary:"更新提示"`
	Id          uint   `p:"id" v:"required#id必须" dc:"id"`
	Position    string `p:"position" v:"required#位置必须" dc:"位置"`
	Content     string `p:"content" v:"required#内容必须" dc:"内容"`
	Description string `p:"description"  dc:"简介"`
	IsDisabled  int    ` dc:"状态"`
}

type PromptUpdateRes struct {
}

type PromptDestroyReq struct {
	g.Meta `path:"/prompt-destroy" method:"delete" tags:"提示管理" summary:"删除提示"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type PromptDestroyRes struct {
}
