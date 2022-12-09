package service

import (
	"bytes"
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/dto"
	"gf-admin/boot"
	"gf-admin/utility/response"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// 中间件管理服务
var (
	Middleware = serviceMiddleware{}
)

type serviceMiddleware struct {
	IgnoreAuthUrls []string
}

// 返回处理中间件
func (s *serviceMiddleware) ResponseHandler(r *ghttp.Request) {

	buffers := bytes.NewBuffer([]byte(""))
	params := r.GetMap()

	//如果params里面包含文件，就忽略掉
	for k, v := range params {
		if reflect.TypeOf(v) != nil && reflect.TypeOf(v).String() == "*multipart.FileHeader" {
			delete(params, k)
		}
	}

	//g.DumpTo(buffers, r.GetMap(), gutil.DumpOption{})

	g.Log().Infof(r.Context(), "请求的url为：%s,客户端端传递过来的参数如下", r.URL.Path)
	g.Log().Infof(r.Context(), "%s", buffers)

	r.Middleware.Next()

	//系统运行时错误
	if err := r.GetError(); err != nil {
		r.Response.Status = 200
		r.Response.ClearBuffer()
		response.JsonErrorLogExit(r, err)
	}

	//如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		res  interface{}
		code gcode.Code = gcode.CodeOK
	)

	res = r.GetHandlerResponse()

	if msg == "" {
		if strings.Contains(r.URL.Path, "-update") {
			msg = "更新成功"
		} else if strings.Contains(r.URL.Path, "-delete") {
			msg = "删除成功"
		} else if strings.Contains(r.URL.Path, "-store") {
			msg = "保存成功"
		} else if strings.Contains(r.URL.Path, "-info") || strings.Contains(r.URL.Path, "-list") {
			msg = "获取成功"
		}
	}

	if res == nil || reflect.ValueOf(res).IsNil() {
		response.JsonExit(r, code.Code(), msg, g.Map{})
	}

	response.JsonExit(r, code.Code(), msg, res)

}

func (s *serviceMiddleware) GetIgnoreAuthUrls() []string {
	if s.IgnoreAuthUrls == nil {
		s.IgnoreAuthUrls = g.Cfg().MustGet(context.TODO(), "casbin.ignoreUrls").Strings()
	}
	return s.IgnoreAuthUrls
}

func (s *serviceMiddleware) OperationLog(r *ghttp.Request) {

	r.Middleware.Next()

	if r.Method == "GET" {
		return
	}

	path := r.URL.Path
	pathName, err := dao.AdminMenu.Ctx(r.Context()).Where(g.Map{
		dao.AdminMenu.Columns.Identification: strings.Replace(path, g.Cfg().MustGet(r.Context(), "server.prefix").String(), "", -1),
	}).Value(dao.AdminMenu.Columns.Name)
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}
	//params := r.GetMap()

	params := g.Map{}
	responseContent := r.GetHandlerResponse()

	if err != nil {
		responseContent = err.Error()
	}

	adminID, err := AdminTokenInstance.GetAdministratorId(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, err)

	}

	_, err = dao.AdminLog.Ctx(r.Context()).Insert(dto.AdminLog{
		AdministratorId: adminID,
		Path:            path,
		Method:          r.Method,
		PathName:        pathName,
		Params:          params,
		Response:        responseContent,
	})
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}
}

func (s *serviceMiddleware) Auth(r *ghttp.Request) {

	g.Log("auth").Debug(r.Context(), "是否登录验证中间件开始执行")
	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}

	if administrator.Id == 0 {
		response.JsonErrorLogExit(r, response.NewError("未登录或会话已过期，请您登录后再继续", g.Map{"administrator": administrator}))

	}

	r.Middleware.Next()
}

func (s *serviceMiddleware) Permission(r *ghttp.Request) {

	administrator, err := AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		response.JsonErrorLogExit(r, response.WrapError(err, "没有权限", g.Map{"administrator": administrator}))

	}

	// 如果是本地开发环境，则直接获取所有的路由和菜单
	if boot.EnvName == "local" {
		r.Middleware.Next()
		return
	}

	url := r.Request.URL.Path
	method := r.Request.Method
	prefix, err := g.Cfg().Get(r.Context(), "server.prefix")
	if err != nil {
		response.JsonErrorLogExit(r, err)
	}

	path := gstr.Replace(url, prefix.String(), "", 1)
	g.Log("auth").Infof(r.Context(), "权限认证,用户为:%s,path为:%s,method为:%s", administrator.Username, path, method)

	if garray.NewStrArrayFrom(s.GetIgnoreAuthUrls(), true).ContainsI(path) {
		r.Middleware.Next()
		return
	}

	isAllow, err := Enforcer.Auth(administrator.Username, path, method)
	if err != nil {

		response.JsonErrorLogExit(r, response.WrapError(err, "没有权限", g.Map{"administrator": administrator, "request path": path}))
	}
	if !isAllow {
		response.JsonErrorLogExit(r, response.NewError("没有权限", g.Map{"administrator": administrator, "request path": path}))
	}
	r.Middleware.Next()
}
