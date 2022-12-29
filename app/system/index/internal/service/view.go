package service

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

type sView struct{}

var insView = sView{}

// 视图管理服务
func View() *sView {
	return &insView
}

// 渲染指定模板页面
func (s *sView) RenderTpl(ctx context.Context, tpl string, data ...define.View) {
	var (
		viewObj  = define.View{}
		viewData = make(g.Map)
		request  = g.RequestFromCtx(ctx)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}

	// 获取logo
	logo, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_LOGO)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	viewObj.Logo = logo.String()

	// 获取网站名称
	siteName, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_NAME)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	viewObj.SiteName = siteName.String()

	// 获取网站描述
	siteDesc, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_DESCRIPTION)

	if err != nil {
		return
	}

	// 获取网站slogan
	siteSlogan, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_SLOGAN)

	if err != nil {
		return
	}

	viewObj.SiteSlogan = siteSlogan.String()

	if viewObj.Title == "" {
		viewObj.Title = viewObj.SiteName
	} else {
		viewObj.Title = viewObj.Title + "-" + viewObj.SiteName
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = ""
	}
	if viewObj.Description == "" {
		viewObj.Description = siteDesc.String()
	}

	// 去掉空数据
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}

	// 如果mainTpl以/开头，则表示是绝对路径，不需要使用layout布局文件
	if gstr.HasPrefix(gconv.String(viewData["Template"]), "/") {
		tpl = gstr.TrimLeft(gconv.String(viewData["Template"]), "/")
	}

	// 渲染模板
	_ = request.Response.WriteTpl(tpl, viewData)

	if gmode.IsDevelop() {

		//	g.Dump("viewData", viewData)
		_ = request.Response.WriteTplContent(`{{dump .}}`, viewData)
	}
}

// 渲染默认模板页面
func (s *sView) Render(ctx context.Context, data ...define.View) {

	s.RenderTpl(ctx, g.Cfg().MustGet(ctx, "front.templateLayout").String(), data...)
}

// 跳转中间页面
func (s *sView) Render302(ctx context.Context, data ...define.View) {
	view := define.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "页面跳转中"
	}
	view.Template = s.getViewFolderName(ctx) + "/pages/302.html"
	s.Render(ctx, view)
}

// 401页面
func (s *sView) Render401(ctx context.Context, data ...define.View) {
	view := define.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.Template = "fail/401.html"
	s.Render(ctx, view)
}

// 403页面
func (s *sView) Render403(ctx context.Context, data ...define.View) {
	view := define.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.Template = "fail/403.html"
	s.Render(ctx, view)
}

// 404页面
func (s *sView) Render404(ctx context.Context, data ...define.View) {
	view := define.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	view.Template = "fail/404.html"
	s.Render(ctx, view)

}

// 500页面
func (s *sView) Render500(ctx context.Context, data ...define.View) {
	view := define.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	view.Template = "fail/500.html"
	s.Render(ctx, view)
}

// 获取视图存储目录
func (s *sView) getViewFolderName(ctx context.Context) string {
	return gstr.Split(g.Cfg().MustGet(ctx, "front.templateLayout").String(), "/")[0]
}
