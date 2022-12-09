// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Nodes is the golang structure for table _nodes.
type Nodes struct {
	Id             uint        `json:"id"               ` //
	Name           string      `json:"name"             ` // 节点名称
	Keyword        string      `json:"keyword"          ` // 节点关键词
	Description    string      `json:"description"      ` // 节点描述
	Detail         string      `json:"detail"           ` // 节点详情
	Img            string      `json:"img"              ` // 节点图片
	ParentId       uint        `json:"parent_id"        ` // 父节点id
	IsIndex        int         `json:"is_index"         ` // 是否首页显示
	IsDisabledEdit int         `json:"is_disabled_edit" ` // 是否禁用编辑和删除,1是 0否
	Sort           int         `json:"sort"             ` // 显示顺序越小越靠前
	CreatedAt      *gtime.Time `json:"created_at"       ` // 创建时间
	DeletedAt      *gtime.Time `json:"deleted_at"       ` // 删除时间
	IsVirtual      int         `json:"is_virtual"       ` //
	CategoryId     int         `json:"category_id"      ` //
}
