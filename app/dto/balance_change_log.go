// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceChangeLog is the golang structure of table forum_balance_change_log for DAO operations like Where/data.
type BalanceChangeLog struct {
	g.Meta     `orm:"table:forum_balance_change_log, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户id
	Username   interface{} // 用户名
	Type       interface{} // 每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register
	Amount     interface{} // 金额
	Before     interface{} // 变动前余额
	After      interface{} // 变动后余额
	RelationId interface{} // 关联主题id或关联回复id
	Remark     interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
}
