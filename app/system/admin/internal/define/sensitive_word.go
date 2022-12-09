package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SensitiveWordListReq struct {
	g.Meta  `method:"get" path:"/sensitive-word-list" summary:"敏感词列表" tags:"敏感词管理"`
	Keyword string
	model.PageSizeInput
}

type SensitiveWordListRes struct {
	List []*entity.SensitiveWords `json:"list"`
	model.PageSizeOutput
}

type SensitiveWordStoreReq struct {
	g.Meta `method:"POST" path:"/sensitive-word-store" summary:"敏感词库存储" tags:"敏感词管理"`
	Word   string `v:"required#敏感词不能为空"`
	Type   string
}

type SensitiveWordStoreRes struct {
}

type SensitiveWordDestroyReq struct {
	g.Meta `method:"delete" path:"/sensitive-word-destroy" summary:"敏感词库删除" tags:"敏感词管理"`
	Ids    []int `v:"required#ids不能为空"`
}

type SensitiveWordDestroyRes struct {
}
