// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLog is the golang structure for table admin_log.
type AdminLog struct {
	Id              uint        `json:"id"               ` //
	AdministratorId uint        `json:"administrator_id" ` // 管理员id
	Path            string      `json:"path"             ` // 请求路径
	Method          string      `json:"method"           ` // 请求方法
	PathName        string      `json:"path_name"        ` // 请求路径名称
	Params          string      `json:"params"           ` // 请求参数
	Response        string      `json:"response"         ` // 响应结果
	CreatedAt       *gtime.Time `json:"created_at"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       ` // 更新时间
}