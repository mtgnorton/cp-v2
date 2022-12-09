package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type NoAuthLoginCaptchaReq struct {
	g.Meta `path:"/captcha-get" method:"get" summary:"获取登录验证码" tags:"全局"`
}

type NoAuthLoginCaptchaRes struct {
	model.CommonGenerateCaptchaOutput
}

type NoAuthIndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
}

type NoAuthIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
