package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PostsDetailPageReq struct {
	g.Meta `path:"/post/{post_id}" method:"get" summary:"主题详情页面" tags:"主题相关"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
}

type PostsDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type AppletPostsDetailReq struct {
	g.Meta `path:"/applet/post" method:"get" summary:"小程序主题详情" tags:"主题相关"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
	model.PageSizeInput
}

type AppletPostsDetailRes struct {
	PostWithComments *model.PostWithNodeAndCommentsRes `json:"post_with_comments"`
}
