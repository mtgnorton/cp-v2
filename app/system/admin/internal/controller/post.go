package controller

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

var Post = post{}

type post struct {
}

func (p *post) List(ctx context.Context, req *define.PostListReq) (res *define.PostListRes, err error) {
	res = &define.PostListRes{}
	req.PostListInput.OrderFunc = func(d *gdb.Model) *gdb.Model {
		if req.OrderDirection != "" && req.OrderField != "" {
			d = d.Order(req.OrderField + " " + req.OrderDirection)
		} else {
			d = d.OrderDesc(dao.Posts.Columns().Id)
		}
		return d
	}
	if req.Id != 0 {
		req.PostListInput.PostIds = []uint{req.Id}
	}
	frontDomain, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_DOMAIN)

	if err != nil {
		return
	}
	res.FrontDomain = frontDomain.String()
	res.PostListOutput, err = shared.Post.List(ctx, req.PostListInput)
	return
}

// Audit 审核主题
func (p *post) Audit(ctx context.Context, req *define.PostAuditReq) (res *define.PostAuditRes, err error) {
	err = service.Post.Audit(ctx, req.Id)
	return
}

// Destroy 删除主题
func (p *post) Destroy(ctx context.Context, req *define.PostDestroyReq) (res *define.PostDestroyRes, err error) {
	err = shared.Post.Destroy(ctx, req.Id)
	return
}

func (p *post) ToggleTop(ctx context.Context, req *define.PostToggleTopReq) (res *define.PostToggleTopRes, err error) {
	err = service.Post.ToggleTop(ctx, req.PostToggleTopInput)
	return
}

// Info 获取主题详情
func (p *post) Info(ctx context.Context, req *define.PostInfoReq) (res *define.PostInfoRes, err error) {
	res = &define.PostInfoRes{}
	res.Posts, err = shared.Post.Info(ctx, req.Id)
	return
}

// Update 更新主题
func (p *post) Update(ctx context.Context, req *define.PostUpdateReq) (res *define.PostUpdateRes, err error) {
	err = service.Post.Update(ctx, req)
	return
}
