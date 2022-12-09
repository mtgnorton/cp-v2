package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/captcha-get" method:"get" summary:"获取验证码" tags:"全局"`
}

type CaptchaRes struct {
	model.CommonGenerateCaptchaOutput
}

type AppletCaptchaReq struct {
	g.Meta `path:"/applet-captcha-get" method:"get" summary:"小程序获取验证码" tags:"全局"`
}

type AppletCaptchaRes struct {
	model.CommonGenerateCaptchaOutput
}
