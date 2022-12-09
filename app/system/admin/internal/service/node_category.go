package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
)

var NodeCategory = nodeCategory{}

type nodeCategory struct {
}

// List 获取节点分类列表
func (p *nodeCategory) List(ctx context.Context, req *define.NodeCategoryListReq) (res *define.NodeCategoryListRes, err error) {
	d := dao.NodeCategories.Ctx(ctx)
	res = &define.NodeCategoryListRes{}

	res.Page = req.Page
	res.Size = req.Size

	res.Total, err = d.Count()
	if err != nil {
		return
	}
	err = d.Page(req.Page, req.Size).Scan(&res.List)

	return
}

// Store 创建一个新的节点分类
func (p *nodeCategory) Store(ctx context.Context, req *define.NodeCategoryStoreReq) (err error) {
	d := dao.NodeCategories.Ctx(ctx)
	count, err := d.
		Where(dao.NodeCategories.Columns().Name, req.Name).
		Where(dao.NodeCategories.Columns().ParentId, req.ParentId).
		Count()

	if err != nil {
		return
	}
	if count > 0 {
		return response.NewError("名称在该父级下重复")
	}
	_, err = d.Insert(g.Map{
		dao.NodeCategories.Columns().Name:              req.Name,
		dao.NodeCategories.Columns().ParentId:          req.ParentId,
		dao.NodeCategories.Columns().IsIndexNavigation: req.IsIndexNavigation,
		dao.NodeCategories.Columns().Sort:              req.Sort,
	})

	return
}

// Info 获取节点分类信息
func (p *nodeCategory) Info(ctx context.Context, req *define.NodeCategoryInfoReq) (res *define.NodeCategoryInfoRes, err error) {
	d := dao.NodeCategories.Ctx(ctx)
	res = &define.NodeCategoryInfoRes{}
	err = d.
		Where(dao.NodeCategories.Columns().Id, req.Id).
		Scan(res)

	return
}

// Update 更新节点分类
func (p *nodeCategory) Update(ctx context.Context, req *define.NodeCategoryUpdateReq) (err error) {
	d := dao.NodeCategories.Ctx(ctx)
	count, err := d.
		WhereNot(dao.NodeCategories.Columns().Id, req.Id).
		Where(dao.NodeCategories.Columns().Name, req.Name).
		Where(dao.NodeCategories.Columns().ParentId, req.ParentId).
		Count()

	if err != nil {
		return
	}
	if count > 0 {
		return response.NewError("名称在该父级下重复")
	}
	_, err = d.
		Where(dao.NodeCategories.Columns().Id, req.Id).
		Update(g.Map{
			dao.NodeCategories.Columns().Name:              req.Name,
			dao.NodeCategories.Columns().ParentId:          req.ParentId,
			dao.NodeCategories.Columns().IsIndexNavigation: req.IsIndexNavigation,
			dao.NodeCategories.Columns().Sort:              req.Sort,
		})

	return
}

// Destroy 删除节点分类
func (p *nodeCategory) Destroy(ctx context.Context, req *define.NodeCategoryDestroyReq) (err error) {
	_, err = dao.NodeCategories.Ctx(ctx).
		Where(dao.NodeCategories.Columns().Id, req.Id).
		Delete()

	return
}
