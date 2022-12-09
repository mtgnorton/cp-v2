// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Association is the golang structure of table forum_association for DAO operations like Where/data.
type Association struct {
	g.Meta         `orm:"table:forum_association, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户id
	Username       interface{} // 用户名
	TargetId       interface{} // 被感谢｜屏蔽|收藏 主题id|回复id
	AdditionalId   interface{} // 附加id，当target_id为回复id时，additional_id为主题id
	TargetUserId   interface{} // 被感谢｜屏蔽|收藏 用户id
	TargetUsername interface{} // 被感谢用户名
	Type           interface{} // 类型 感谢主题: thanks_posts,感谢回复: thanks_reply,屏蔽主题: shield_posts,屏蔽回复: shield_reply,收藏主题:collect_posts
	CreatedAt      *gtime.Time // 创建时间
}
