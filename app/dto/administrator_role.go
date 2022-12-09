// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdministratorRoleForDao is the golang structure of table ga_administrator_role for DAO operations like Where/data.
type AdministratorRole struct {
	g.Meta          `orm:"dto:true"`
	Id              interface{} //
	AdministratorId interface{} // 管理员id
	RoleId          interface{} // 角色id
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
