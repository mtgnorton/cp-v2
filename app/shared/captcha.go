package shared

import (
	"context"
	"gf-admin/app/model"

	"github.com/mojocn/base64Captcha"
)

var Captcha = captcha{}

type captcha struct {
}

func (c *captcha) Generate(ctx context.Context) (out model.CommonGenerateCaptchaOutput, err error) {
	var store = base64Captcha.DefaultMemStore

	driver := base64Captcha.NewDriverDigit(50, 120, 5, 0, 50)

	instance := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := instance.Generate()

	return model.CommonGenerateCaptchaOutput{
		CaptchaId:     id,
		CaptchaBase64: b64s,
	}, err

}

func (c *captcha) Verify(ctx context.Context, code string, id string) bool {
	if code == "" || id == "" {
		return false
	}
	var store = base64Captcha.DefaultMemStore
	return store.Verify(id, code, true)

}
