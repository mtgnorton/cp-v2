package define

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/index" method:"get" tags:"首页" summary:"首页"`
}

type IndexRes struct {
	TodayUserCount  int      `json:"today_user_count" dc:"今日新增用户数"`
	TodayPostCount  int      `json:"today_post_count" dc:"今日新增话题数"`
	TodayReplyCount int      `json:"today_reply_count" dc:"今日新增回复数"`
	AllUserCount    int      `json:"all_user_count" dc:"总用户数"`
	AllPostCount    int      `json:"all_post_count" dc:"总话题数"`
	AllReplyCount   int      `json:"all_reply_count" dc:"总回复数"`
	Day30UserInc    []int    `json:"day_30_user_inc" dc:"最近30天用户增长"`
	Day30PostInc    []int    `json:"day_30_post_inc" dc:"最近30天话题增长"`
	Day30ReplyInc   []int    `json:"day_30_reply_inc" dc:"最近30天回复增长"`
	Day30Date       []string `json:"day_30_date" dc:"最近30天日期"`
}
