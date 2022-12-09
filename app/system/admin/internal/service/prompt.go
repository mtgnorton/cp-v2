package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
)

var Prompt = prompt{}

type prompt struct {
}

// List 获取提示列表
func (p *prompt) List(ctx context.Context, req *define.PromptListReq) (res *define.PromptListRes, err error) {
	d := dao.Prompts.Ctx(ctx)
	res = &define.PromptListRes{}

	res.Page = req.Page
	res.Size = req.Size

	res.Total, err = d.Count()
	if err != nil {
		return
	}
	err = d.Page(req.Page, req.Size).Scan(&res.List)

	return
}

// Store 创建一个新的提示
func (p *prompt) Store(ctx context.Context, req *define.PromptStoreReq) (err error) {
	d := dao.Prompts.Ctx(ctx)
	count, err := d.
		Where(dao.Prompts.Columns().Position, req.Position).
		Where(dao.Prompts.Columns().IsDisabled, 0).
		Count()

	if err != nil {
		return
	}
	if req.IsDisabled == 0 && count > 0 {
		return response.NewError("该位置已经存在使用中的提示")
	}
	_, err = d.Insert(g.Map{
		dao.Prompts.Columns().Position:    req.Position,
		dao.Prompts.Columns().Content:     req.Content,
		dao.Prompts.Columns().Description: req.Description,
		dao.Prompts.Columns().IsDisabled:  req.IsDisabled,
	})

	return
}

// Info 获取提示信息
func (p *prompt) Info(ctx context.Context, req *define.PromptInfoReq) (res *define.PromptInfoRes, err error) {
	d := dao.Prompts.Ctx(ctx)
	res = &define.PromptInfoRes{}
	err = d.
		Where(dao.Prompts.Columns().Id, req.Id).
		Scan(res)

	return
}

// Update 更新提示
func (p *prompt) Update(ctx context.Context, req *define.PromptUpdateReq) (err error) {
	d := dao.Prompts.Ctx(ctx)
	count, err := d.
		WhereNot(dao.Prompts.Columns().Id, req.Id).
		Where(dao.Prompts.Columns().Position, req.Position).
		Where(dao.Prompts.Columns().IsDisabled, 0).
		Count()

	if err != nil {
		return
	}
	if req.IsDisabled == 0 && count > 0 {
		return response.NewError("该位置已经存在使用中的提示")
	}
	_, err = d.
		Where(dao.Prompts.Columns().Id, req.Id).
		Update(g.Map{
			dao.Prompts.Columns().Position:    req.Position,
			dao.Prompts.Columns().Content:     req.Content,
			dao.Prompts.Columns().Description: req.Description,
			dao.Prompts.Columns().IsDisabled:  req.IsDisabled,
		})

	return
}

// Destroy 删除提示语
func (p *prompt) Destroy(ctx context.Context, req *define.PromptDestroyReq) (err error) {
	_, err = dao.Prompts.Ctx(ctx).
		Where(dao.Prompts.Columns().Id, req.Id).
		Delete()

	return
}
