package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type NodeCategoryListReq struct {
	g.Meta `path:"/node-category-list" method:"get" tags:"节点分类" summary:"节点分类列表"`
	model.PageSizeInput
}
type NodeCategoryListRes struct {
	List []*entity.NodeCategories `json:"list"`
	model.PageSizeOutput
}

type NodeCategoryStoreReq struct {
	g.Meta            `path:"/node-category-store" method:"post" tags:"节点分类" summary:"保存节点分类"`
	Name              string `p:"name" v:"required#名称必须" dc:"名称"`
	ParentId          int64  `p:"parent_id"  dc:"父级id"`
	IsIndexNavigation int    `p:"is_index_navigation"  dc:"是否首页导航"`
	Sort              int    `p:"sort"  dc:"排序"`
}

type NodeCategoryStoreRes struct {
}

type NodeCategoryInfoReq struct {
	g.Meta `path:"/node-category-info" method:"get" tags:"节点分类" summary:"节点分类详情"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type NodeCategoryInfoRes struct {
	entity.NodeCategories
}

type NodeCategoryUpdateReq struct {
	g.Meta            `path:"/node-category-update" method:"put" tags:"节点分类" summary:"更新节点分类"`
	Id                uint   `p:"id" v:"required#id必须" dc:"id"`
	Name              string `p:"name" v:"required#名称必须" dc:"名称"`
	ParentId          int64  `p:"parent_id"  dc:"父级id"`
	IsIndexNavigation int    `p:"is_index_navigation"  dc:"是否首页导航"`
	Sort              int    `p:"sort"  dc:"排序"`
}

type NodeCategoryUpdateRes struct {
}

type NodeCategoryDestroyReq struct {
	g.Meta `path:"/node-category-destroy" method:"delete" tags:"节点分类" summary:"删除节点分类"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type NodeCategoryDestroyRes struct {
}
