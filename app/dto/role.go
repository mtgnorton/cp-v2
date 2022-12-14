// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleForDao is the golang structure of table ga_role for DAO operations like Where/data.
type Role struct {
	g.Meta         `orm:"dto:true"`
	Id             interface{} //
	Name           interface{} // 角色名
	Identification interface{} // 角色标识符
	Sort           interface{} // 显示顺序，越小越靠前
	Status         interface{} // 状态 normal 正常 disabled 禁用
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
