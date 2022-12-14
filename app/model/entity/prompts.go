// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Prompts is the golang structure for table _prompts.
type Prompts struct {
	Id          uint        `json:"id"          ` //
	Position    string      `json:"position"    ` // 提示语位置
	Content     string      `json:"content"     ` // 提示语内容
	Description string      `json:"description" ` // 简介
	IsDisabled  int         `json:"is_disabled" ` // 状态：0 正常 1禁用
	CreatedAt   *gtime.Time `json:"created_at"  ` // 创建时间
	UpdateAt    *gtime.Time `json:"update_at"   ` // 更新时间
}
