package service

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/system/index/internal/define"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestReply_Store(t *testing.T) {
	before()
	gtest.C(t, func(t *gtest.T) {
		_, err := Reply.Store(ctx, &define.ReplyStoreReq{
			PostId:  2,
			Content: "test",
		})
		t.AssertNil(err)
	})
	after()
}

func before() {
	// 使用monkey模拟redis的赋值和取值操作
	monkey.PatchInstanceMethod(reflect.TypeOf(&FrontTokenInstance), "GetUser", func(_ *frontTokenHandle, ctx context.Context) (*model.UserSummary, error) {
		return &model.UserSummary{
			Id:       1,
			Username: "mtgnorton",
			Email:    "",
		}, nil
	})
}

func after() {
	monkey.UnpatchAll()

}
