package shared

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"

	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/text/gregex"

	"github.com/gogf/gf/v2/text/gstr"
)

var Common = common{}

type common struct {
}

// CompleteUrl 把相对路径补全成域名,如果没有传递域名，则使用后台配置的
func (c *common) CompleteUrl(ctx context.Context, path string, domains ...string) (url string, err error) {

	var domain string
	if len(domains) == 0 {
		domain, err = Config.GetString(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_DOMAIN)
	}
	return gstr.Trim(domain, "/") + "/" + gstr.Trim(path, "/"), err
}

// TransformReferLink 将内容中的@username 转换为a链接
func (c *common) TransformReferLink(content string) string {

	rs, _ := gregex.ReplaceStringFuncMatch(`@([a-zA-Z0-9_]+)`, content, func(match []string) string {
		return `<a href="/user/` + match[1] + `" data-username="` + match[1] + `">@` + match[1] + `</a>`
	})
	return rs
}

// SiteStatisticsInfo 获取网站的注册会员数，主题数，回复数
func (c *common) SiteStatisticsInfo(ctx context.Context) (info *model.SiteInfo, err error) {
	info = &model.SiteInfo{}
	info.UserCount, err = dao.Users.Ctx(ctx).Count()
	if err != nil {
		return
	}
	info.PostCount, err = dao.Posts.Ctx(ctx).Count()
	if err != nil {
		return
	}
	info.ReplyCount, err = dao.Replies.Ctx(ctx).Count()
	return
}

// GetSiteNameAndLogo 获取网站的名称和logo
func (c *common) GetSiteNameAndLogo(ctx context.Context) (siteName string, logo string, err error) {
	siteName, err = Config.GetString(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_NAME)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	logo, err = Config.GetString(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_FORUM_SITE_LOGO)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	logo, err = Common.CompleteUrl(ctx, logo)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

// Prompt 用于提示信息
func (c *common) Prompt(ctx context.Context, message, redirectUrl string) {

	r := ghttp.RequestFromCtx(ctx)
	redirectUrl = gstr.Replace(redirectUrl, "/", "\\")

	r.Response.RedirectTo(fmt.Sprintf("/prompt/%s/%s", message, gurl.Encode(redirectUrl)))
}

// TruncateDatabase 清空数据库里面的部分业务数据
func (c *common) TruncateDatabase(ctx context.Context) (err error) {
	// 清空posts表
	_, err = g.DB().Exec(ctx, "truncate table forum_posts")
	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_replies")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_users")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_messages")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_balance_change_log")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_association")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_user_posts_histories")

	if err != nil {
		return
	}
	_, err = g.DB().Exec(ctx, "truncate table forum_email_records")

	if err != nil {
		return
	}

	_, err = g.DB().Exec(ctx, "truncate table ga_admin_log")

	if err != nil {
		return
	}
	return nil
}
