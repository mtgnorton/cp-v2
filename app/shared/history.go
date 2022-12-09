package shared

import (
	"context"
	"gf-admin/app/dao"
)

var History = history{}

type history struct {
}

// GetPostIds 根据用户id获取用户浏览过的主题id
func (h *history) GetPostIds(ctx context.Context, userId uint, amounts ...int) (postIds []uint, err error) {
	amount := 5
	if len(amounts) > 0 {
		amount = amounts[0]
	}
	items, err := dao.UserPostsHistories.Ctx(ctx).
		Where(dao.UserPostsHistories.Columns().UserId, userId).
		OrderDesc(dao.UserPostsHistories.Columns().Id).
		Limit(amount).
		Distinct().
		Array(dao.UserPostsHistories.Columns().PostsId)

	if err != nil {
		return
	}
	for _, item := range items {
		postIds = append(postIds, item.Uint())

	}

	return
}
