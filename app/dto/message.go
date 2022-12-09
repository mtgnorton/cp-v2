// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table forum_message for DAO operations like Where/data.
type Message struct {
	g.Meta          `orm:"table:forum_message, do:true"`
	Id              interface{} //
	UserId          interface{} // 用户id
	Username        interface{} // 用户名
	RepliedUserId   interface{} // 被回复用户id,用户a向用户b回复，用户b为 被回复用户id
	RepliedUsername interface{} // 被回复用户名
	PostsId         interface{} // 关联主题id
	RepliesId       interface{} // 关联回复id
	IsRead          interface{} // 是否已读，否: 0,是: 1
	CreatedAt       *gtime.Time // 创建时间
	DeletedAt       *gtime.Time // 删除时间
}
