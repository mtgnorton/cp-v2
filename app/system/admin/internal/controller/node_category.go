package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var NodeCategory = nodeCategory{}

type nodeCategory struct {
}

// List 获取节点分类列表
func (p *nodeCategory) List(ctx context.Context, req *define.NodeCategoryListReq) (res *define.NodeCategoryListRes, err error) {
	res = &define.NodeCategoryListRes{}
	res, err = service.NodeCategory.List(ctx, req)
	return
}

// Store 创建一个新的节点分类
func (p *nodeCategory) Store(ctx context.Context, req *define.NodeCategoryStoreReq) (res *define.NodeCategoryStoreRes, err error) {
	err = service.NodeCategory.Store(ctx, req)
	return
}

// Info 获取节点分类信息
func (p *nodeCategory) Info(ctx context.Context, req *define.NodeCategoryInfoReq) (res *define.NodeCategoryInfoRes, err error) {
	res, err = service.NodeCategory.Info(ctx, req)
	return
}

// Update 更新一个节点分类
func (p *nodeCategory) Update(ctx context.Context, req *define.NodeCategoryUpdateReq) (res *define.NodeCategoryUpdateRes, err error) {

	err = service.NodeCategory.Update(ctx, req)
	return
}

// Destroy 删除一个节点分类
func (p *nodeCategory) Destroy(ctx context.Context, req *define.NodeCategoryDestroyReq) (res *define.NodeCategoryDestroyRes, err error) {
	err = service.NodeCategory.Destroy(ctx, req)
	return
}
