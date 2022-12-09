package admin

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/controller"
	"gf-admin/app/system/admin/internal/service"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/util/gmode"
)

func Run(ctx context.Context) {
	var (
		s   = g.Server("admin")
		oai = s.GetOpenApi()
	)

	// OpenApi自定义信息
	oai.Info.Title = `API Reference`
	oai.Config.CommonResponse = response.JsonRes{}
	oai.Config.CommonResponseDataField = `data`

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}
	if !gfile.IsDir(uploadPath) {
		err := gfile.Mkdir(uploadPath)
		if err != nil {
			return
		}
	}
	s.AddStaticPath("/upload", uploadPath)

	err := g.View().AddPath(g.Cfg().MustGet(ctx, "server.templatePath").String())
	if err != nil {
		g.Log().Fatal(ctx, "view path error, %s", err)
	}
	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}

	prefix, err := g.Cfg().Get(ctx, "server.prefix")
	if err != nil {
		g.Log().Fatalf(ctx, "get server admin prefix error,error info following : %s", err)
	}
	g.Dump(prefix, "prefix")

	service.AdminTokenInstance.Init(ctx)

	// 前台系统路由注册
	s.Group(prefix.String(), func(group *ghttp.RouterGroup) {

		// 使用传统路由方式绑定websocket请求
		group.ALL("/ws", controller.Ws.Ws)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(
				shared.Middleware.Ctx,
				service.Middleware.OperationLog,
				service.Middleware.ResponseHandler,
			)
			//无需登录验证的路由

			group.Bind(controller.NoAuth)

			group.Group("/", func(group *ghttp.RouterGroup) {

				//需要登录验证的路由
				service.AdminTokenInstance.LoadConfig().Middleware(group)

				group.Middleware(service.Middleware.Auth)

				group.Bind(controller.Personal)
				group.Bind(controller.Global)

				// 需要权限验证的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware.Permission)

					group.Bind(
						controller.Index,
						controller.Administrator,
						controller.Role,
						controller.Menu,
						controller.Config,
						controller.OperationLog,
						controller.Node,
						controller.User,
						controller.Post,
						controller.Reply,
						controller.SensitiveWord,
						controller.Prompt,
						controller.EmailRecord,
						controller.BalanceLog,
						controller.NodeCategory,
						controller.Association,
					)
				})
			})
		})

	})

	// 如果访问/重定向到/admin
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.RedirectTo("/admin")
	})
	// 自定义丰富文档
	enhanceOpenAPIDoc(s)
	sessionConfig(s)
	service.Enforcer.Register(ctx)

	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			//g.Log().Debug(ctx, ghttp.HookBeforeServe)
			//r.SetParam("key1", "v11")
			//r.GetRequest("key1")
		},
	})

	//controller.Ws.MonitorSystem(ctx)
	// 启动Http Server
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	s.Run()
}
func sessionConfig(s *ghttp.Server) {

	err := s.SetConfigWithMap(g.Map{
		"SessionStorage": gsession.NewStorageRedis(g.Redis("session")),
	})
	if err != nil {
		g.Log().Fatalf(gctx.New(), "init session driver error, %s", err)
	}
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `data`

	openapi.Components = goai.Components{
		SecuritySchemes: goai.SecuritySchemes{
			"ApiKeyAuth": goai.SecuritySchemeRef{
				Ref: "", // 暂时还不知道该值是干什么用的
				Value: &goai.SecurityScheme{
					Type: "apiKey",
					In:   "header",
					Name: "Authorization",
				},
			},
		},
	}
	openapi.Security = &goai.SecurityRequirements{
		goai.SecurityRequirement{"ApiKeyAuth": []string{}},
	}

	// API description.
	openapi.Info.Title = `forum`
	openapi.Info.Description = `后台接口文档`

}
