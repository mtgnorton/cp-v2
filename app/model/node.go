package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/util/gmeta"
)

type Node struct {
	gmeta.Meta `orm:"table:forum_nodes"`
	entity.Nodes
}

// 获取节点详情，可以通过Id ｜ keyword 查询
type NodeDetailInput struct {
	Id      uint
	Keyword string
}

type NodePageListOutput struct {
	List []*entity.Nodes `json:"list"`
	PageSizeOutput
}

type NodePageListInput struct {
	Name    string `json:"name"`
	IsIndex string `v:"in:1,0" json:"is_index"`
	PageSizeInput
}

type NodeListInput struct {
	Name         string
	IsIndex      string // 0否 1是 为空不查询
	IsVirtual    string // 0否 1是 为空不查询
	NeedChildren bool
}

type NodeTree struct {
	entity.Nodes
	Children     []*NodeTree
	NodeCategory entity.NodeCategories `json:"node_category"`
}
type NodeListOutput struct {
	List []*NodeTree `json:"list"`
}
