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

var Node = node{}

type node struct {
}

// NodePage 节点页面
func (n *node) NodePage(ctx context.Context, req *define.NodePageReq) (res *define.NodePageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	node, err := shared.Node.Detail(ctx, model.NodeDetailInput{
		Keyword: req.Keyword,
	})
	if err != nil {
		return
	}
	if node.Id == 0 {
		service.View().Render404(ctx)
		return
	}

	var posts *model.PostListOutput

	if node.IsVirtual == 1 {
		posts, err = service.Post.AllPostList(ctx, model.PageSizeInput{
			Page: req.Page,
			Size: 20,
		})
	} else {
		posts, err = service.Post.NodeList(ctx, req.Keyword, model.PageSizeInput{
			Page: req.Page,
			Size: 20,
		})
	}

	if err != nil {
		return
	}
	pager := service.Pager.Pager(&model.PagerReq{
		CurrentPage:    req.Page,
		Size:           20,
		TotalRow:       posts.Total,
		Url:            "/node-page/" + req.Keyword + "/%d",
		ShowPageAmount: 10,
	})

	isCollectNode, err := shared.User.WhetherCollectNode(ctx, user.Id, node.Id)
	if err != nil {
		return
	}
	nodeCollectAmount, err := shared.Node.GetCollectAmount(ctx, node.Id)

	if err != nil {
		return
	}

	service.View().Render(ctx, define.View{
		Title:       node.Name,
		Description: node.Description,
		Template:    "node.html",
		User:        user,
		Data: g.Map{
			"node":              node,
			"posts":             posts,
			"pager":             pager,
			"isCollectNode":     isCollectNode,
			"nodeCollectAmount": nodeCollectAmount,
		},
	})
	return
}

func (n *node) NodeListPage(ctx context.Context, req *define.NodeListPageReq) (res *define.NodeListPageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	nodeCategories, err := service.NodeCategory.List(ctx, &define.NodeCategoryListInput{
		ParentId: 0,
	})
	if err != nil {
		return
	}
	service.View().Render(ctx, define.View{
		Title:    "节点列表",
		Template: "node-list.html",
		User:     user,
		Data: g.Map{
			"nodeCategories": nodeCategories,
		},
	})
	return
}
