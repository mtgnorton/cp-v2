package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"time"
)

var Index = index{}

type index struct {
}

func (i *index) Statistics(ctx context.Context) (res *define.IndexRes, err error) {
	res = &define.IndexRes{}

	// 获取所有的用户，帖子，回复 数量
	res.AllUserCount, err = dao.Users.Ctx(ctx).Count()
	if err != nil {
		return
	}
	res.AllPostCount, err = dao.Posts.Ctx(ctx).Count()
	if err != nil {
		return
	}

	res.AllReplyCount, err = dao.Replies.Ctx(ctx).Count()
	if err != nil {
		return
	}

	// 获取最近30天内每日新增的用户，帖子，回复数量

	// select DATE_FORMAT(created_at,'%Y-%m-%d') days,count(id) count from forum_users where created_at >  group by days;

	day30Before := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	type dayCount struct {
		Day   string `json:"day"`
		Count int    `json:"count"`
	}

	var userCountSlice []*dayCount

	err = dao.Users.Ctx(ctx).Where("created_at > ?", day30Before).Fields("DATE_FORMAT(created_at,'%Y-%m-%d') day,count(id) count").Group("day").Scan(&userCountSlice)

	userCountMap := make(map[string]int)

	for _, v := range userCountSlice {
		userCountMap[v.Day] = v.Count
	}

	var postCountSlice []*dayCount
	err = dao.Posts.Ctx(ctx).Where("created_at > ?", day30Before).Fields("DATE_FORMAT(created_at,'%Y-%m-%d') day,count(id) count").Group("day").Scan(&postCountSlice)
	postCountMap := make(map[string]int)
	for _, v := range postCountSlice {
		postCountMap[v.Day] = v.Count
	}

	var replyCountSlice []*dayCount
	err = dao.Replies.Ctx(ctx).Where("created_at > ?", day30Before).Fields("DATE_FORMAT(created_at,'%Y-%m-%d') day,count(id) count").Group("day").Scan(&replyCountSlice)
	replyCountMap := make(map[string]int)

	for _, v := range replyCountSlice {
		replyCountMap[v.Day] = v.Count
	}

	for i := 29; i >= 0; i-- {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		if i == 0 {
			res.TodayUserCount = userCountMap[day]
			res.TodayPostCount = postCountMap[day]
			res.TodayReplyCount = replyCountMap[day]
		}
		res.Day30UserInc = append(res.Day30UserInc, userCountMap[day])
		res.Day30PostInc = append(res.Day30PostInc, postCountMap[day])
		res.Day30ReplyInc = append(res.Day30ReplyInc, replyCountMap[day])
		res.Day30Date = append(res.Day30Date, time.Now().AddDate(0, 0, -i).Format("01-02"))
	}
	return
}
