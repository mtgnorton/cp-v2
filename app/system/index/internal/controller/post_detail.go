package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
)

var PostDetail = postDetail{}

type postDetail struct {
}

// 主题页面
func (p *postDetail) PostsDetailPage(ctx context.Context, req *define.PostsDetailPageReq) (res *define.PostsDetailRes, err error) {
	//判断主题是否存在
	exist, err := shared.Post.Exist(ctx, req.PostId)
	if err != nil {
		return
	}
	if !exist {
		service.View().Render404(ctx, define.View{Error: "主题不存在"})
		return
	}
	// 判断主题是否未审核

	exist, err = shared.Post.Exist(ctx, req.PostId, model.POST_STATUS_NORMAL)
	if err != nil {
		return
	}
	if !exist {
		service.View().Render404(ctx, define.View{Error: "主题正在审核中"})
		return
	}

	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	err = service.Post.Visit(ctx, req.PostId, user)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	postWithComments, err := service.Post.DetailWithNodeAndComments(ctx, model.PostWithNodeAndCommentsReq{
		Id:        req.PostId,
		SeeUserId: user.Id,
	})

	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}

	isCollectPost, err := shared.User.WhetherCollectPost(ctx, user.Id, req.PostId)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	isShieldPost, err := shared.User.WhetherShieldPost(ctx, user.Id, req.PostId)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	isThankPost, err := shared.User.WhetherThanksPost(ctx, user.Id, req.PostId)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	//获取当前用户感谢的所有回复id
	thanksReplyIds, err := shared.User.GetThanksReplyIds(ctx, user.Id, req.PostId)

	seoDesc, err := service.Post.GetSeoDesc(ctx, &postWithComments.Posts)
	if err != nil {
		return
	}

	replyPrompt, err := shared.Prompt.GetContent(ctx, model.PROMPT_POSTION_REPLY_NEW)

	service.View().Render(ctx, define.View{
		Title:       postWithComments.Title,
		Description: seoDesc,
		Template:    "post-detail.html",
		User:        user,
		Data: g.Map{
			"Post":           postWithComments,
			"IsCollectPost":  isCollectPost,
			"IsShieldPost":   isShieldPost,
			"IsThankPost":    isThankPost,
			"ThanksReplyIds": thanksReplyIds,
			"ReplyPrompt":    replyPrompt,
		},
	})
	return
}
