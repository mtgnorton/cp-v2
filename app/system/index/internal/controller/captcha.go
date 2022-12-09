package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
)

var Captcha = captcha{}

type captcha struct {
}

// 图形验证码
func (l *captcha) Captcha(ctx context.Context, req *define.CaptchaReq) (res *define.CaptchaRes, err error) {
	res = &define.CaptchaRes{}
	res.CommonGenerateCaptchaOutput, err = shared.Captcha.Generate(ctx)
	return
}
