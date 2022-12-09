// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Replies is the golang structure for table _replies.
type Replies struct {
	Id              uint        `json:"id"                ` //
	PostsId         uint        `json:"posts_id"          ` // 主题id
	UserId          uint        `json:"user_id"           ` // 用户id
	Username        string      `json:"username"          ` // 用户名
	RelationUserIds string      `json:"relation_user_ids" ` // 涉及用户ids，多个以逗号分隔
	Content         string      `json:"content"           ` // 内容
	CharacterAmount uint        `json:"character_amount"  ` // 字符长度
	ThanksAmount    uint        `json:"thanks_amount"     ` // 感谢次数
	ShieldedAmount  uint        `json:"shielded_amount"   ` // 被屏蔽次数
	Status          int         `json:"status"            ` // 状态：no_audit, normal, shielded
	CreatedAt       *gtime.Time `json:"created_at"        ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"        ` // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at"        ` // 删除时间
}
