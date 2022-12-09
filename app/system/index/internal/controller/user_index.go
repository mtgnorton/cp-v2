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

var UserIndex = userIndex{}

type userIndex struct {
}

// Index 用户中心首页
func (u *userIndex) Index(ctx context.Context, req *define.UserIndexReq) (res *define.UserIndexRes, err error) {

	showUser, err := shared.User.Info(ctx, &model.UserInfoInput{
		Username: req.Username,
	})
	if err != nil {
		return
	}
	if showUser.Id == 0 {
		return res, response.NewError("用户不存在")
	}
	user, err := service.User.GetTemplateShow(ctx)

	if err != nil {
		return
	}
	posts, err := shared.Post.List(ctx, &model.PostListInput{
		Usernames: []string{req.Username},
	})
	if err != nil {
		return
	}

	replies, err := shared.Reply.List(ctx, &model.ReplyListInput{
		UserId:   showUser.Id,
		WithPost: true,
	})
	if err != nil {
		return
	}
	t := "posts"

	if req.Type == "replies" {
		t = "replies"
	}

	whetherFollow, err := shared.User.WhetherFollowUser(ctx, user.Id, showUser.Id)

	if err != nil {
		return
	}
	whetherShield, err := shared.User.WhetherShieldUser(ctx, user.Id, showUser.Id)
	if err != nil {
		return
	}

	service.View().Render(ctx, define.View{
		Title:    "个人中心",
		User:     user,
		Template: "user/index.html",
		Data: g.Map{
			"showUser":      showUser,
			"type":          t,
			"posts":         posts,
			"replies":       replies,
			"whetherFollow": whetherFollow,
			"whetherShield": whetherShield,
		},
	})
	return
}
