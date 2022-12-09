package service

import (
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"

	"github.com/gogf/gf/v2/frame/g"
)

func TestPosts_Store(t *testing.T) {
	err := shared.Config.Sets(ctx, model.CONFIG_MODULE_FORUM, g.Map{
		model.CONFIG_POSTS_CHARACTER_MAX:          300,
		model.CONFIG_POSTS_EVERY_DAY_MAX:          10,
		model.CONFIG_TOKEN_ESTABLISH_POSTS_DEDUCT: 10,
	})
	g.Dump(err)
}

func TestPosts_List(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Post.List(ctx, &model.PostListInput{
			IsIndex: true,
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}

func TestPosts_DetailWithNodeAndComments(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		detail, err := Post.DetailWithNodeAndComments(ctx, model.PostWithNodeAndCommentsReq{
			Id:        2,
			SeeUserId: 2,
		})
		t.Assert(err, nil)
		g.Dump(detail)
	})
}
