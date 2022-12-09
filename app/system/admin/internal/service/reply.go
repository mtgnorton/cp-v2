package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Reply = reply{}

type reply struct {
}

// Audit 审核回复
func (r *reply) Audit(ctx context.Context, replyId uint) (err error) {
	// 判断是否需要审核
	d := dao.Replies.Ctx(ctx)

	count, err := d.Where(dao.Replies.Columns().Id, replyId).
		Where(dao.Replies.Columns().Status, model.REPLY_STATUS_NO_AUDIT).
		Count()
	if err != nil {
		return response.NewError("审核失败")
	}
	if count == 0 {
		return response.NewError("该回复不需要审核")
	}
	// 审核
	_, err = d.WherePri(replyId).Data(dao.Replies.Columns().Status, model.REPLY_STATUS_NORMAL).Update()
	if err != nil {
		return response.NewError("审核失败")
	}
	return shared.Reply.OfficialPublishHook(ctx, gconv.Int64(replyId))

}

// Update 更新主题
func (r *reply) Update(ctx context.Context, in *define.ReplyUpdateReq) (err error) {
	// 判断是否需要审核
	d := dao.Replies.Ctx(ctx)

	if in.Status == model.REPLY_STATUS_NO_AUDIT {
		return response.NewError("状态不能为未审核")
	}
	_, err = d.Where(dao.Replies.Columns().Id, in.Id).Update(g.Map{
		dao.Replies.Columns().Status: in.Status,
	})
	return
}
