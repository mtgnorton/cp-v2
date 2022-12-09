package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"
)

var User = user{}

type user struct {
}

// BalanceLog  余额变动列表
func (u *user) BalanceLog(ctx context.Context, req *define.UserBalanceLogReq) (res *define.UserBalanceLogRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	log, err := shared.BalanceChangeLog.List(ctx, &model.BalanceChangeLogListInput{
		UserId: user.Id,
		PageSizeInput: model.PageSizeInput{
			Page: req.Page,
			Size: 30,
		},
	})
	if err != nil {
		return
	}

	pager := service.Pager.Pager(&model.PagerReq{
		CurrentPage:    req.Page,
		Size:           30,
		TotalRow:       log.Total,
		Url:            "/user/balance-log/%d",
		ShowPageAmount: 10,
	})

	service.View().Render(ctx, define.View{
		Title:    "余额记录",
		User:     user,
		Template: "user/balance_log.html",
		Data: g.Map{
			"log":   log,
			"pager": pager,
		},
	})
	return
}

// Message 消息提醒列表
func (u *user) Message(ctx context.Context, req *define.UserMessagePageReq) (res *define.UserMessagePageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	messages, err := shared.Message.List(ctx, &model.MessageListInput{
		RepliedUserId: user.Id,
		PageSizeInput: model.PageSizeInput{
			Page: req.Page,
			Size: 10,
		},
	})
	g.Dump(messages)
	if err != nil {
		return
	}
	// 将返回的消息设置为已读
	err = shared.Message.SetMessagesRead(ctx, gdb.ListItemValuesUnique(messages.List, "Message", "Id"))
	if err != nil {
		return
	}
	pager := service.Pager.Pager(&model.PagerReq{
		CurrentPage:    req.Page,
		Size:           10,
		TotalRow:       messages.Total,
		Url:            "/user/message/%d",
		ShowPageAmount: 10,
	})

	service.View().Render(ctx, define.View{
		Title:    "消息",
		User:     user,
		Template: "user/message.html",
		Data: g.Map{
			"messages": messages,
			"pager":    pager,
		},
	})
	return
}

// Logout 退出登录
func (p *user) Logout(ctx context.Context, req *define.LogoutReq) (res *define.LogoutRes, err error) {
	res = &define.LogoutRes{}
	err = service.User.Logout(ctx)

	g.RequestFromCtx(ctx).Response.RedirectTo("/")

	return
}
