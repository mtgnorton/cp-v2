// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenuForDao is the golang structure of table ga_admin_menu for DAO operations like Where/data.
type AdminMenu struct {
	g.Meta             `orm:"dto:true"`
	Id                 interface{} //
	Name               interface{} // 菜单名称
	Path               interface{} // 前端路由地址，可以是外链
	ParentId           interface{} // 父id
	Identification     interface{} // 后端权限标识符
	Method             interface{} // 请求方法
	FrontComponentPath interface{} // 前端组件路径
	Icon               interface{} // 菜单图标
	Sort               interface{} // 显示顺序，越小越靠前
	Status             interface{} // 状态 normal 正常 disabled 禁用
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
	Type               interface{} // 菜单类型
	LinkType           interface{} //
}
