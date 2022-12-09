package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Prompt = prompt{}

type prompt struct {
}

// List 获取提示列表
func (p *prompt) List(ctx context.Context, req *define.PromptListReq) (res *define.PromptListRes, err error) {
	res = &define.PromptListRes{}
	res, err = service.Prompt.List(ctx, req)
	return
}

// Store 创建一个新的提示
func (p *prompt) Store(ctx context.Context, req *define.PromptStoreReq) (res *define.PromptStoreRes, err error) {
	err = service.Prompt.Store(ctx, req)
	return
}

// Info 获取提示信息
func (p *prompt) Info(ctx context.Context, req *define.PromptInfoReq) (res *define.PromptInfoRes, err error) {
	res, err = service.Prompt.Info(ctx, req)
	return
}

// Update 更新一个提示
func (p *prompt) Update(ctx context.Context, req *define.PromptUpdateReq) (res *define.PromptUpdateRes, err error) {

	err = service.Prompt.Update(ctx, req)
	return
}

// Destroy 删除一个提示
func (p *prompt) Destroy(ctx context.Context, req *define.PromptDestroyReq) (res *define.PromptDestroyRes, err error) {
	err = service.Prompt.Destroy(ctx, req)
	return
}
