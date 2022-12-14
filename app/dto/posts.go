// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure of table forum_posts for DAO operations like Where/data.
type Posts struct {
	g.Meta            `orm:"table:forum_posts, do:true"`
	Id                interface{} //
	NodeId            interface{} // 节点id
	UserId            interface{} // 用户id
	Username          interface{} // 用户名
	Title             interface{} // 标题
	ContentType       interface{} // quill为使用quill编辑器，markdown为使用markdown编辑器,general 为普通文本
	Content           interface{} // 内容
	HtmlContent       interface{} // html内容
	TopEndTime        *gtime.Time // 置顶截止时间,为空说明没有置顶
	CharacterAmount   interface{} // 字符长度
	VisitAmount       interface{} // 访问次数
	CollectionAmount  interface{} // 收藏次数
	ReplyAmount       interface{} // 回复次数
	ThanksAmount      interface{} // 感谢次数
	ShieldedAmount    interface{} // 被屏蔽次数
	Status            interface{} // 状态：0 未审核 1 已审核
	Weight            interface{} // 权重
	ReplyLastUserId   interface{} // 最后回复用户id
	ReplyLastUsername interface{} // 最后回复用户名
	LastChangeTime    *gtime.Time // 主题最后变动时间
	CreatedAt         *gtime.Time // 主题创建时间
	UpdatedAt         *gtime.Time // 主题更新时间
	DeletedAt         *gtime.Time // 删除时间
}
