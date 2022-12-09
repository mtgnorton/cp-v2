package response

import (
	"gf-admin/utility/custom_log"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code     int         `json:"code"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message  string      `json:"message"`  // 提示信息
	Data     interface{} `json:"data"`     // 返回数据(业务接口定义具体数据结构)
	Redirect string      `json:"redirect"` // 引导客户端跳转到指定路由
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})

}

func JsonErrorLogExit(r *ghttp.Request, err error) {

	r.Response.ClearBuffer()

	code := gerror.Code(err)

	var data interface{}
	data = g.Map{}

	if code.Code() != gcode.CodeOK.Code() {
		go custom_log.Log(r, err)
	} else {
		// 如果返回的code为0，说明是返回的成功消息,不记录日志，并且把相关数据返回给调用方
		data = r.GetHandlerResponse()
	}
	if code == gcode.CodeNil && err != nil {
		code = gcode.CodeInternalError
	}

	JsonExit(r, code.Code(), gerror.Current(err).Error(), data)
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

// JsonRedirect 返回标准JSON数据引导客户端跳转。
func JsonRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonRes{
		Code:     code,
		Message:  message,
		Data:     responseData,
		Redirect: redirect,
	})

}

// JsonRedirectExit 返回标准JSON数据引导客户端跳转，并退出当前HTTP执行函数。
func JsonRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JsonRedirect(r, code, message, redirect, data...)
	r.Exit()
}
