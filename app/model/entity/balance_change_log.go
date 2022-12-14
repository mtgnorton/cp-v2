// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceChangeLog is the golang structure for table _balance_change_log.
type BalanceChangeLog struct {
	Id         uint        `json:"id"          ` //
	UserId     uint        `json:"user_id"     ` // 用户id
	Username   string      `json:"username"    ` // 用户名
	Type       string      `json:"type"        ` // 每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register
	Amount     int         `json:"amount"      ` // 金额
	Before     uint        `json:"before"      ` // 变动前余额
	After      uint        `json:"after"       ` // 变动后余额
	RelationId uint        `json:"relation_id" ` // 关联主题id或关联回复id
	Remark     string      `json:"remark"      ` // 备注
	CreatedAt  *gtime.Time `json:"created_at"  ` // 创建时间
}
