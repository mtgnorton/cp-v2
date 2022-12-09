//go:build !codeanalysis

package shared

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/container/gvar"

	"bou.ke/monkey"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

var (
	tokenHandler = TokenHandler{}
	userKey      = "admin"
	data         = g.Map{
		"username": "admin",
		"id":       1,
		"nickname": "admin",
		"status":   "normal",
		"remark":   "备注",
	}
)

func TestTokenHandler_TokenGenerateAndSaveData(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tokenHandler.Init(ctx)
		token, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		fmt.Printf("TestTokenHandler_TokenGenerateAndSaveData Token is %s\n", token)
	})
	gtest.C(t, func(t *gtest.T) { //测试redis缓存
		tokenHandler.CacheMode = CacheModeRedis
		tokenHandler.Init(ctx)
		token, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		fmt.Printf("TestTokenHandler_TokenGenerateAndSaveData Token is %s\n", token)
	})

}

//func TestTokenHandler_Validate(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		tokenHandler.CacheMode = CacheModeRedis
//		tokenHandler.Init(ctx)
//		r, err := tokenHandler.Validate(ctx, "VYgiJ280XDOohxfueR6bE/P3w6qoIJaWcABC4T2YNeZGv07NOK3TGriPQoJYRm8V")
//		t.AssertNil(err)
//		fmt.Printf("TestTokenHandler_TokenValidate result is %s\n", &r)
//	})
//}

func TestTokenHandler_TokenValidate(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		tokenHandler.Init(ctx)
		token, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		r, err := tokenHandler.Validate(ctx, token)
		t.AssertNil(err)
		fmt.Printf("TestTokenHandler_TokenValidate result is %s\n", &r)
	})

	/*多点登录*/
	gtest.C(t, func(t *gtest.T) {
		tokenHandler.MultiLogin = true
		tokenHandler.Init(ctx)
		token1, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		r, err := tokenHandler.Validate(ctx, token1)
		t.AssertNil(err)
		fmt.Printf("TestTokenHandler_TokenValidate result is %s\n", &r)
		token2, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.Assert(token1, token2)
		t.AssertNil(err)

	})
}

func TestTokenHandler_TokenRemove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tokenHandler.Init(ctx)
		token, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		err = tokenHandler.Remove(ctx, token)
		t.AssertNil(err)
	})
}

func TestTokenHandler_TokenGetData(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tokenHandler.CacheMode = CacheModeRedis
		tokenHandler.Init(ctx)
		token, err := tokenHandler.GenerateAndSaveData(ctx, userKey, data)
		t.AssertNil(err)
		fmt.Printf("test Token is %s\n", token)

		data, err := tokenHandler.GetData(ctx, userKey)
		t.AssertNil(err)

		fmt.Printf("TestTokenHandler_TokenGetData data is %s", &data)
	})

}

func setup() {
	// 使用monkey模拟redis的赋值和取值操作
	keyValueMap := make(map[string]*gvar.Var)
	monkey.PatchInstanceMethod(reflect.TypeOf(g.Redis()), "Do", func(redis *gredis.Redis, ctx context.Context, command string, args ...interface{}) (*gvar.Var, error) {
		if command == "setex" {
			keyValueMap[args[0].(string)] = gvar.New(args[2])
			return nil, nil
		}
		if command == "set" {
			keyValueMap[args[0].(string)] = gvar.New(args[1])
			return nil, nil
		}

		if command == "get" {
			return gvar.New(keyValueMap[args[0].(string)]), nil
		}
		return nil, nil
	})
}

func teardown() {
	monkey.UnpatchAll()
}

// 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
