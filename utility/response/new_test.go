package response

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_New_Wrap_Error(t *testing.T) {

	//用户自定义错误
	gtest.C(t, func(t *gtest.T) {
		err := NewError("自定义错误", map[string]interface{}{"name": "jack"}, []string{"card"})
		t.Assert(gerror.Current(err), "自定义错误")
		contextInfo, err := getContextVariables(err)

		t.Assert(err, nil)
		t.Assert(contextInfo, `[{"name":"jack"},["card"]]`)

	})
	//包裹自定义错误
	gtest.C(t, func(t *gtest.T) {
		err := NewError("自定义错误", map[string]interface{}{"name": "jack"}, []string{"card"})
		ew := WrapError(err, "包裹自定义错误", map[string]interface{}{"age": 18}, []string{"user"})
		t.Assert(gerror.Current(ew), "包裹自定义错误")
		contextInfo, err := getContextVariables(ew)
		t.Assert(err, nil)
		t.Assert(contextInfo, `[{"age":18},["user"]]`)
	})

	//包裹原始错误
	gtest.C(t, func(t *gtest.T) {
		e := errors.New("原始错误")

		ew := WrapError(e, "包裹自定义错误", map[string]interface{}{"name": "jack"})

		t.Assert(gerror.Current(ew), "包裹自定义错误")

		contextInfo, err := getContextVariables(ew)
		t.Assert(err, nil)
		t.Assert(contextInfo, `[{"name":"jack"}]`)
	})
	//包裹原始错误
	gtest.C(t, func(t *gtest.T) {

		err := gerror.New("goframe 错误")
		ew := WrapError(err, "")
		t.Assert(gerror.Current(ew), "未知错误")
		contextInfo, err := getContextVariables(ew)
		t.Assert(err, nil)
		t.Assert(contextInfo, `null`)

	})
}

func getContextVariables(err error) (string, error) {
	code := gerror.Code(err)
	c, err := json.Marshal(code.Detail())

	return string(c), err

}
