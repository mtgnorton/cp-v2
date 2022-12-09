package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

//自定义错误状态码
const CustomCodeNumber = 1000

func NewError(message string, contextArgs ...interface{}) error {
	if message == "" {
		message = "未知错误"
	}
	code := gcode.New(CustomCodeNumber, message, contextArgs)
	return gerror.NewCodeSkip(code, 1)
}

func NewSuccess(message string, contextArgs ...interface{}) error {
	if message == "" {
		message = "操作成功"
	}
	code := gcode.New(0, message, contextArgs)

	return gerror.NewCodeSkip(code, 100)
}

func WrapError(err error, message string, contextArgs ...interface{}) error {
	if message == "" {
		message = "未知错误"
	}
	code := gcode.New(CustomCodeNumber, message, contextArgs)

	return gerror.WrapCodeSkip(code, 1, err)
}
