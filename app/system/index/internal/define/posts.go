package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
)

type PostsNewPageReq struct {
	g.Meta `path:"/post-new" method:"get" tags:"主题相关" summary:"发表主题页面"`
}

type PostsNewPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PostsListReq struct {
	g.Meta `path:"/posts" method:"get" tags:"主题相关" summary:"主题列表"`
}

type PostsListRes struct {
	List []model.PostWithoutContent `json:"list"`
}

type PostQuill struct {
	Ops []struct {
		Attributes struct {
			Bold   bool   `json:"bold,omitempty"`
			Indent int    `json:"indent,omitempty"`
			Script string `json:"script,omitempty"`
			Color  string `json:"color,omitempty"`
		} `json:"attributes,omitempty"`
		Insert string `json:"insert"`
	} `json:"ops"`
}

type PostsStoreReq struct {
	g.Meta      `path:"/posts-store" method:"post" summary:"创建主题" tags:"主题相关"`
	NodeId      uint   `v:"required#节点id不能为空" json:"node_id"`
	Title       string `v:"required#标题不能为空" json:"title"`
	Content     string `v:"required#内容不能为空" json:"content"`
	HtmlContent string `v:"required#html内容不能为空" json:"html_content"`
}

type PostsStoreRes struct {
	Id uint `json:"id"`
}

type PostsUploadImageReq struct {
	g.Meta `path:"/posts-upload-image" method:"post" summary:"上传图片" tags:"主题相关"`
	Image  *ghttp.UploadFile `v:"required#图片不能为空" json:"image"`
}

type PostsUploadImageRes struct {
	Url string `json:"url"`
}

type PostsUpdateReq struct {
	g.Meta  `path:"/posts-update" method:"post" summary:"更新主题" tags:"主题相关"`
	Id      uint   `v:"required#主题id不能为空" json:"id"`
	Title   string `v:"required#标题不能为空" json:"title"`
	Content string `v:"required#内容不能为空" json:"content"`
}

type PostsUpdateRes struct {
}

type PostsMoveReq struct {
	g.Meta `path:"/posts-move" method:"post" summary:"移动主题" tags:"主题相关"`
	Id     uint `v:"required#主题id不能为空" json:"id"`
	NodeId uint `v:"required#新节点id不能为空" json:"node_id"`
}

type PostsMoveRes struct {
}

type PostsToggleCollectReq struct {
	g.Meta `path:"/posts-toggle-collect" method:"post" summary:"收藏主题" tags:"收藏|关注|屏蔽|感谢"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
}

type PostsToggleCollectRes struct {
}

type PostsToggleShieldReq struct {
	g.Meta `path:"/posts-toggle-shield" method:"post" summary:"屏蔽主题" tags:"收藏|关注|屏蔽|感谢"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
}

type PostsToggleShieldRes struct {
}

type PostsThanksReq struct {
	g.Meta `path:"/posts-thanks" method:"post" summary:"感谢主题" tags:"收藏|关注|屏蔽|感谢"`
	PostId uint `v:"required#主题id不能为空" json:"post_id"`
}

type PostsThanksRes struct {
}

type ReplyThanksReq struct {
	g.Meta  `path:"/replies-thanks" method:"post" summary:"感谢回复" tags:"收藏|关注|屏蔽|感谢"`
	ReplyId uint `v:"required#回复id不能为空" json:"reply_id" dc:"回复id"`
}

type ReplyThanksRes struct {
}

type ReplyShieldReq struct {
	g.Meta  `path:"/replies-shield" method:"post" summary:"屏蔽回复" tags:"收藏|关注|屏蔽|感谢"`
	ReplyId uint `v:"required#回复id不能为空" json:"reply_id" dc:"回复id"`
}

type ReplyShieldRes struct {
}
