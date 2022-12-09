package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/database/gdb"
)

var Message = message{}

type message struct {
}

// List 获取消息列表
func (m *message) List(ctx context.Context, in *model.MessageListInput) (out *model.MessageListOutput, err error) {
	out = &model.MessageListOutput{}
	d := dao.Messages.Ctx(ctx)

	if in.RepliedUserId != 0 {
		d = d.Where(dao.Messages.Columns().RepliedUserId, in.RepliedUserId)
	}
	if in.IsRead != "" {
		d = d.Where(dao.Messages.Columns().IsRead, in.IsRead)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return
	}
	err = d.Page(in.Page, in.Size).
		OrderDesc(dao.Messages.Columns().Id).
		ScanList(&out.List, "Message")
	if err != nil {
		return
	}

	err = dao.Posts.Ctx(ctx).
		WherePri(gdb.ListItemValuesUnique(out.List, "Message", "PostId")).
		ScanList(&out.List, "Post", "Message", "id:PostId")
	if err != nil {
		return
	}

	err = dao.Replies.Ctx(ctx).
		WherePri(gdb.ListItemValuesUnique(out.List, "Message", "ReplyId")).
		ScanList(&out.List, "Reply", "Message", "id:ReplyId")
	if err != nil {
		return
	}

	err = dao.Users.Ctx(ctx).
		WherePri(gdb.ListItemValuesUnique(out.List, "Message", "UserId")).
		ScanList(&out.List, "User", "Message", "id:UserId")
	if err != nil {
		return
	}
	return
}

// GetUserNoReadCount 获取未读消息数量
func (m *message) GetUserNoReadCount(ctx context.Context, userId uint) (count int, err error) {
	count, err = dao.Messages.Ctx(ctx).
		Where(dao.Messages.Columns().RepliedUserId, userId).
		Where(dao.Messages.Columns().IsRead, 0).
		Count()
	return
}

// SetMessagesRead 设置消息已读
func (m *message) SetMessagesRead(ctx context.Context, ids []interface{}) (err error) {
	_, err = dao.Messages.Ctx(ctx).
		Where(dao.Messages.Columns().Id, ids).
		Update(g.Map{
			dao.Messages.Columns().IsRead: 1,
		})
	return
}
