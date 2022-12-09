package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var Post = post{}

type post struct {
}

func (p *post) NewPage(ctx context.Context, req *define.PostsNewPageReq) (res *define.PostsNewPageRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	nodes, err := shared.Node.List(ctx, &model.NodeListInput{
		IsVirtual: "0",
	})
	prompt, err := shared.Prompt.GetContent(ctx, model.PROMPT_POSTION_POST_NEW)
	service.View().Render(ctx, define.View{
		Title:    "创建新主题",
		User:     user,
		Template: "post-new.html",
		Data: g.Map{
			"node":   nodes,
			"prompt": prompt,
		},
	})
	return
}

func (p *post) Store(ctx context.Context, req *define.PostsStoreReq) (res *define.PostsStoreRes, err error) {
	return service.Post.Store(ctx, req)
}

//UploadPostImage 上传主题中包含的图片
func (p *post) UploadPostImage(ctx context.Context, req *define.PostsUploadImageReq) (res *define.PostsUploadImageRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	return service.Post.UploadImage(ctx, user.Id, req)
}

func (p *post) Reply(ctx context.Context, req *define.ReplyStoreReq) (res *define.ReplyStoreRes, err error) {
	return service.Reply.Store(ctx, req)
}

func (p *post) ToggleCollect(ctx context.Context, req *define.PostsToggleCollectReq) (res *define.PostsToggleCollectRes, err error) {
	res = &define.PostsToggleCollectRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	err = service.Post.ToggleCollect(ctx, req.PostId, user.Id, user.Username)
	return
}

func (p *post) ToggleShield(ctx context.Context, req *define.PostsToggleShieldReq) (res *define.PostsToggleShieldRes, err error) {
	res = &define.PostsToggleShieldRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	err = service.Post.ToggleShield(ctx, req.PostId, user.Id, user.Username)
	return
}

func (p *post) ThanksPost(ctx context.Context, req *define.PostsThanksReq) (res *define.PostsThanksRes, err error) {
	res = &define.PostsThanksRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	err = service.Post.ThanksPost(ctx, req.PostId, user.Id, user.Username)
	return
}

func (p *post) ThanksReply(ctx context.Context, req *define.ReplyThanksReq) (res *define.ReplyThanksRes, err error) {
	res = &define.ReplyThanksRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	err = service.Post.ThanksReply(ctx, req.ReplyId, user.Id, user.Username)
	return
}

func (p *post) ShieldReply(ctx context.Context, req *define.ReplyShieldReq) (res *define.ReplyShieldRes, err error) {
	res = &define.ReplyShieldRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	err = service.Post.ShieldReply(ctx, req.ReplyId, user.Id, user.Username)
	return
}
