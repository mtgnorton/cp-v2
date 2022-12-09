package service

import (
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestAuth(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		shared.Config.Sets(ctx, model.CONFIG_MODULE_FORUM, g.Map{
			model.CONFIG_TOKEN_REGISTER_GIVE: 1000,
			model.CONFIG_TOKEN_LOGIN_GIVE:    20,
		})
	})
}
