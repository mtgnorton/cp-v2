package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Post = post{}

type post struct {
}

func (p *post) ToggleTop(ctx context.Context, in *define.PostToggleTopInput) (err error) {

	err = p.ExistById(ctx, in.Id)
	if err != nil {
		return err
	}
	if !in.EndTime.IsZero() && in.EndTime.Before(gtime.Now()) {
		return response.NewError("置顶截止时间不能小于当前时间")
	}
	_, err = dao.Posts.Ctx(ctx).Where(dao.Posts.Columns().Id, in.Id).Update(g.Map{
		dao.Posts.Columns().TopEndTime: in.EndTime,
	})

	return
}

// ExistById 主题是否存在
func (p *post) ExistById(ctx context.Context, Id uint) (err error) {

	exist, err := dao.Posts.Ctx(ctx).WherePri(Id).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "主题不存在")
	}
	return nil
}

// Audit 审核主题
func (p *post) Audit(ctx context.Context, postId uint) (err error) {
	// 判断是否需要审核
	d := dao.Posts.Ctx(ctx)

	count, err := d.Where(dao.Posts.Columns().Id, postId).
		Where(dao.Posts.Columns().Status, model.POST_STATUS_NO_AUDIT).
		Count()
	if err != nil {
		return response.NewError("审核失败")
	}
	if count == 0 {
		return response.NewError("该主题不需要审核")
	}
	// 审核
	_, err = d.WherePri(postId).Data(dao.Posts.Columns().Status, model.REPLY_STATUS_NORMAL).Update()
	if err != nil {
		return response.NewError("审核失败")
	}
	return shared.Post.OfficialPublishHook(ctx, gconv.Int64(postId))
}

// Update 更新主题
func (p *post) Update(ctx context.Context, in *define.PostUpdateReq) (err error) {
	err = p.ExistById(ctx, in.Id)
	if err != nil {
		return err
	}
	if in.Status == model.POST_STATUS_NO_AUDIT {
		return response.NewError("状态不能为未审核")
	}
	_, err = dao.Posts.Ctx(ctx).WherePri(in.Id).Update(
		g.Map{
			dao.Posts.Columns().Status:     in.Status,
			dao.Posts.Columns().TopEndTime: in.TopEndTime,
		})
	return
}
