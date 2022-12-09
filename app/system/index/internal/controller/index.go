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

var Index = index{}

type index struct {
}

func (n *index) IndexPage(ctx context.Context, req *define.IndexReq) (res *define.IndexRes, err error) {

	r := g.RequestFromCtx(ctx)

	//获取cookie里面index-node-keyword对应的节点
	cookieNode := r.Cookie.Get("index-node-keyword")

	// 判断节点是否存在
	exist, err := shared.Node.ExistByKeyword(ctx, cookieNode.String())

	if err != nil {
		return
	}
	if exist {
		r.Response.RedirectTo("/k/"+cookieNode.String(), 302)
		return
	}

	// 取出在首页的节点中，顺序第一(sort值最小)的节点
	firstNode, err := service.Node.GetIndexFirst(ctx)
	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	currentNodeKeyword := firstNode.Keyword
	r.Response.RedirectTo("/k/"+currentNodeKeyword, 302)
	return
}

func (n *index) IndexKeyWordPage(ctx context.Context, req *define.IndexNodeReq) (res *define.IndexNodeRes, err error) {
	currentNodeKeyword := req.Keyword

	// 将当前节点的keyword存入cookie
	r := g.RequestFromCtx(ctx)

	r.Cookie.Set("index-node-keyword", currentNodeKeyword)

	exist, err := shared.Node.ExistByKeyword(ctx, req.Keyword)

	if err != nil {
		return
	}
	if !exist {
		service.View().Render404(ctx)
		return
	}
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	nodes, err := shared.Node.List(ctx, &model.NodeListInput{
		IsIndex:      "1",
		NeedChildren: false,
	})
	if err != nil {
		return
	}

	posts, err := service.Post.ChooseListByKeyword(ctx, currentNodeKeyword)
	if err != nil {
		return
	}

	histories, err := shared.User.GetHistoryPostList(ctx, user.Id)
	if err != nil {
		return
	}

	// 获取今日热议主题,24小时内
	hots, err := service.Post.HotList(ctx, model.Day1, model.PageSizeInput{
		Page: 1,
		Size: 8,
	})
	// 是否显示置顶
	showTop := true

	if currentNodeKeyword == "follow-user" || currentNodeKeyword == "follow-node" {
		showTop = false
	}

	siteInfo, err := shared.Common.SiteStatisticsInfo(ctx)

	nodeCategories, err := service.NodeCategory.List(ctx, &define.NodeCategoryListInput{
		IsIndexNavigation: "1",
	})
	if err != nil {
		return
	}
	service.View().Render(ctx, define.View{
		User:     user,
		Template: "index.html",
		Data: g.Map{
			"nodes":              nodes,
			"showTop":            showTop,
			"posts":              posts,
			"nodeCategories":     nodeCategories,
			"currentNodeKeyword": currentNodeKeyword,
			"histories":          histories,
			"hots":               hots,
			"siteInfo":           siteInfo,
		},
	})
	return
}

func (n *index) SearchPage(ctx context.Context, req *define.SearchReq) (res *define.SearchRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}

	var posts *model.PostListOutput = &model.PostListOutput{}
	var replies *model.ReplyListOutput = &model.ReplyListOutput{}
	pager := &model.PagerRes{}
	g.Dump(req)
	if req.Type == "post" {
		posts, err = shared.Post.List(ctx, &model.PostListInput{
			FilterKeyword: req.Search,
			PageSizeInput: model.PageSizeInput{
				Page: req.Page,
				Size: 20,
			},
		})
		if err != nil {
			return
		}
		pager = service.Pager.Pager(&model.PagerReq{
			CurrentPage:    req.Page,
			Size:           20,
			TotalRow:       posts.Total,
			Url:            "/search/post/" + req.Search + "/%d",
			ShowPageAmount: 10,
		})
	} else if req.Type == "reply" {
		replies, err = shared.Reply.List(ctx, &model.ReplyListInput{
			Keyword:  req.Search,
			WithPost: true,
			PageSizeInput: model.PageSizeInput{
				Page: req.Page,
				Size: 20,
			},
		})
		if err != nil {
			return
		}
		pager = service.Pager.Pager(&model.PagerReq{
			CurrentPage:    req.Page,
			Size:           20,
			TotalRow:       posts.Total,
			Url:            "/search/reply/" + req.Search + "/%d",
			ShowPageAmount: 10,
		})
	}

	service.View().Render(ctx, define.View{
		User:     user,
		Template: "search.html",
		Data: g.Map{
			"posts":   posts,
			"replies": replies,
			"pager":   pager,
			"keyword": req.Search,
			"type":    req.Type,
		},
	})
	return
}
