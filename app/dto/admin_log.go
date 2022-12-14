// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLogForDao is the golang structure of table ga_admin_log for DAO operations like Where/data.
type AdminLog struct {
	g.Meta          `orm:"dto:true"`
	Id              interface{} //
	AdministratorId interface{} // 管理员id
	Path            interface{} // 请求路径
	Method          interface{} // 请求方法
	PathName        interface{} // 请求路径名称
	Params          interface{} // 请求参数
	Response        interface{} // 响应结果
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
