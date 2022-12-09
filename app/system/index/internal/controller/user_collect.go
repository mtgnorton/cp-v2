package controller

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var UserCollect = userCollect{}

type userCollect struct {
}

// MyNodesPage 收藏的节点
func (u *userCollect) MyNodesPage(ctx context.Context, req *define.UserCollectNodePageReq) (res *define.UserCollectNodePageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	collectNodes, err := service.User.GetCollectNodes(ctx, user.Id)
	if err != nil {
		return
	}
	service.View().Render(ctx, define.View{
		Title:    "收藏的节点",
		User:     user,
		Template: "user/collect.html",
		Data: g.Map{
			"type":         "nodes",
			"collectNodes": collectNodes,
		},
	})
	return
}

// MyPosts 收藏的主题
func (u *userCollect) MyPosts(ctx context.Context, req *define.UserCollectPostPageReq) (res *define.UserCollectPostPageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	posts, err := service.User.GetCollectPosts(ctx, user.Id)
	if err != nil {
		return
	}
	service.View().Render(ctx, define.View{
		Title:    "收藏的帖子",
		User:     user,
		Template: "user/collect.html",
		Data: g.Map{
			"type":  "posts",
			"posts": posts,
		},
	})
	return
}

// MyFollowing 关注的用户
func (u *userCollect) MyFollowing(ctx context.Context, req *define.UserFollowUserPageReq) (res *define.UserFollowUserPageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}

	posts, err := service.User.GetFollowUserPosts(ctx, user.Id)
	if err != nil {
		return
	}
	service.View().Render(ctx, define.View{
		Title:    "收藏的用户",
		User:     user,
		Template: "user/collect.html",
		Data: g.Map{
			"type":  "following",
			"posts": posts,
		},
	})
	return
}

// ToggleNodeCollect 切换节点收藏
func (u *userCollect) ToggleNodeCollect(ctx context.Context, req *define.UserToggleCollectNodeReq) (res *define.UserToggleCollectNodeRes, err error) {

	res = &define.UserToggleCollectNodeRes{}
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.Node.ToggleCollect(ctx, req.NodeId, user.Id, user.Username)
	return
}

// ToggleFollowUser  切换关注用户
func (u *userCollect) ToggleFollowUser(ctx context.Context, req *define.UserToggleFollowUserReq) (res *define.UserToggleFollowUserRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.User.ToggleFollowUser(ctx, user.Id, user.Username, req.TargetId)

	return
}

// ToggleShieldUser 切换屏蔽用户
func (u *userCollect) ToggleShieldUser(ctx context.Context, req *define.UserToggleShieldUserReq) (res *define.UserToggleShieldUserRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.User.ToggleShieldUser(ctx, user.Id, user.Username, req.TargetId)
	return
}
