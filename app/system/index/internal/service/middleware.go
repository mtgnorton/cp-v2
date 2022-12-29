package service

import (
	"bytes"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility/response"
	"reflect"
	"strings"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 中间件管理服务
var (
	Middleware = middleware{}
)

type middleware struct {
}

func ErrorExit(r *ghttp.Request, err error) {

	isApplet := r.GetHeader("is-applet")

	if r.IsAjaxRequest() || isApplet == "1" {
		response.JsonErrorLogExit(r, err)
	} else {
		g.Log().Error(gctx.New(), err)
		View().Render500(r.Context(), define.View{
			Error: gerror.Current(err).Error(),
		})
		r.Exit()

	}
}

func (s *middleware) Auth(r *ghttp.Request) {

	g.Log("userNoNeedLogin").Debug(r.Context(), "是否登录验证中间件开始执行")
	user, err := FrontTokenInstance.GetUser(r.Context())
	g.Log("userNoNeedLogin").Debugf(r.Context(), "登录的用户为：%#v", user)

	if err != nil {
		ErrorExit(r, err)
	}

	if user.Id == 0 {
		View().Render403(r.Context(), define.View{
			Error: "未登录或会话已过期，请您登录后再继续",
		})
		r.Exit()
	}

	r.Middleware.Next()
}

// 返回处理中间件
func (s *middleware) ResponseHandler(r *ghttp.Request) {

	buffers := bytes.NewBuffer([]byte(""))
	params := r.GetMap()

	//如果params里面包含文件，就忽略掉
	for k, v := range params {
		if reflect.TypeOf(v).String() == "*multipart.FileHeader" {
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

		ErrorExit(r, err)
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
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), msg, g.Map{})
		} else {
			if r.GetForm("post_by_link").Int() == 1 {
				r.Response.RedirectBack()
			}
		}
	}
	response.JsonExit(r, code.Code(), msg, res)
}

// 允许跨域请求中间件
func (s *middleware) Cors(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// LoginGiveToken  每日登录赠送积分，并记录最后访问时间和访问ip
func (s *middleware) LoginGiveToken(r *ghttp.Request) {
	ctx := r.Context()
	user, err := FrontTokenInstance.GetUser(ctx)

	//g.Dump(user, "loginGiveToken")
	if err != nil {
		ErrorExit(r, err)
	}
	if user.Id == 0 {
		r.Middleware.Next()
		return
	}
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, user.Id).Update(g.Map{
		dao.Users.Columns().LastLoginIp:   r.GetClientIp(),
		dao.Users.Columns().LastLoginTime: gtime.Now(),
	})
	if err != nil {
		ErrorExit(r, err)
	}
	// 每日登录赠送积分
	count, err := dao.BalanceChangeLog.Ctx(ctx).
		Where(dao.BalanceChangeLog.Columns().UserId, user.Id).
		WhereGTE(dao.BalanceChangeLog.Columns().CreatedAt, gtime.Date()).
		Where(dao.BalanceChangeLog.Columns().Type, model.BALANCE_CHANGE_TYPE_LOGIN).Count()

	if err != nil {
		ErrorExit(r, err)
	}
	if count == 0 {
		settingLoginGiveAmount, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_LOGIN_GIVE)
		if err != nil {
			ErrorExit(r, err)
		}
		err = shared.User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     user.Id,
			Amount:     settingLoginGiveAmount.Int(),
			ChangeType: model.BALANCE_CHANGE_TYPE_LOGIN,
		})
		if err != nil {
			ErrorExit(r, err)
		}
	}

	r.Middleware.Next()
}
