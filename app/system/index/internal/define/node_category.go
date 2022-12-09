package define

import "gf-admin/app/model/entity"

type NodeCategoryListInput struct {
	IsIndexNavigation string  // 0否 1是 空为不查询
	Ids               []int64 // 为nil不查询
	ParentId          uint    //
}

type NodeCategoryItem struct {
	*entity.NodeCategories
	Nodes []*entity.Nodes `json:"nodes"`
}
type NodeCategoryListOutput struct {
	List      []*NodeCategoryItem `json:"list"`
	NodeTotal int                 `json:"node_total"`
}
