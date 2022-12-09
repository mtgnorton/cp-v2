package shared

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

var Middleware = middleware{}

type middleware struct {
}

// 自定义上下文对象
func (s *middleware) Ctx(r *ghttp.Request) {
	g.Log().Debug(r.Context(), "ctx中间件开始执行")
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Data: make(g.Map),
	}
	Context.Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *middleware) TokenInCookieToHeader(r *ghttp.Request) {
	v := r.Cookie.Get("cp-v2-token")
	g.Dump("cp-v2-token", v.String())
	if v.String() != "" {
		r.Header.Set("Authorization", "Bearer "+v.String())
	} else {
		if r.Cookie.Get("cp-v2-unique-id").String() == "" {
			r.Cookie.Set("cp-v2-unique-id", gconv.String(gtime.Now().Timestamp())+grand.S(14))
		}
	}
	// 如果用户未登录,给每个未登录的用户分配一个唯一的id

	r.Middleware.Next()
}
