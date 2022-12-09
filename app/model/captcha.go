package model

type CommonGenerateCaptchaOutput struct {
	CaptchaId     string `json:"captcha_id"`
	CaptchaBase64 string `json:"captcha_base64"`
}
