package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

type TimePeriod int

const (
	DayAll TimePeriod = 0
	Today  TimePeriod = -1 //代表今日0点到现在
	Day1   TimePeriod = 1  // 代表此时向前推24小时
	Day3   TimePeriod = 3  // 代表此时向前推3天
	Day7   TimePeriod = 7  // 代表此时向前推7天
	Month  TimePeriod = 30
	Year   TimePeriod = 365
)

const (
	POST_STATUS_NORMAL   = 1
	POST_STATUS_NO_AUDIT = 0
	POST_STATUS_SHIELD   = -1
)

type Post struct {
	gmeta.Meta `orm:"table:forum_posts"`
	entity.Posts
}

type PostContentType string

const (
	POST_CONTENT_TYPE_GENERAL  PostContentType = "general"
	POST_CONTENT_TYPE_QUILL    PostContentType = "quill"
	POST_CONTENT_TYPE_MARKDOWN PostContentType = "markdown"
)

type PostWithoutContent struct {
	Id                uint        `json:"id"                  `                   //
	NodeId            uint        `json:"node_id"             `                   // 节点id
	NodeName          string      `json:"node_name"           `                   // 节点名称
	NodeKeyword       string      `json:"node_keyword"           `                // 节点名称
	UserId            uint        `json:"user_id"             `                   // 用户id
	Username          string      `json:"username"            `                   // 用户名
	UserAvatar        string      `json:"user_avatar"  orm:"user_avatar"        ` // 用户头像
	Title             string      `json:"title"               `                   // 标题
	TopEndTime        *gtime.Time `json:"top_end_time"        `                   // 置顶截止时间,为空说明没有置顶
	CharacterAmount   uint        `json:"character_amount"    `                   // 字符长度
	VisitAmount       uint        `json:"visit_amount"       `                    // 访问次数
	CollectionAmount  uint        `json:"collection_amount"   `                   // 收藏次数
	ReplyAmount       uint        `json:"reply_amount"        `                   // 回复次数
	ThanksAmount      uint        `json:"thanks_amount"       `                   // 感谢次数
	ShieldedAmount    uint        `json:"shielded_amount"     `                   // 被屏蔽次数
	Status            int         `json:"status"              `                   // 状态：no_audit, normal, shielded
	Weight            int         `json:"weight"              `                   // 权重
	ReplyLastUserId   uint        `json:"reply_last_user_id"  `                   // 最后回复用户id
	ReplyLastUsername string      `json:"reply_last_username" `                   // 最后回复用户名
	CreatedAt         *gtime.Time `json:"created_at"          `                   // 主题创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"          `                   // 主题更新时间
	DeletedAt         *gtime.Time `json:"deleted_at"          `                   // 删除时间
	LastChangeTime    *gtime.Time `json:"last_change_time"   `                    // 最后变动时间
	IsTop             int         `json:"is_top"              `                   // 是否置顶
}

type PostListInput struct {
	Period        TimePeriod `json:"period" dc:"时间段，0:所有,1:1天，3:3天，7:7天，30:30天，365:365天"`
	Title         string     `json:"title" dc:"标题"`
	NodeIds       []uint     `json:"node_id" dc:"节点ids"`
	NodeKeyword   string     `json:"node_keyword" dc:"节点关键字"`
	NodeName      string     `json:"node_name" `
	UserIds       []uint     `json:"user_id" dc:"用户ids"`
	Usernames     []string   `json:"username" dc:"用户名数组"`
	FilterKeyword string     `json:"filter_keyword" dc:"过滤关键字"`
	Status        string     //0 未审核 1 正常  为空不筛选
	PostIds       []uint     `json:"post_ids" dc:"查询特定的id列表"`
	PageSizeInput            // 只有当page和size都存在时，才会进行分页
	OrderFunc     func(*gdb.Model) *gdb.Model
}

type PostListOutput struct {
	PageSizeOutput
	List []*PostWithoutContent `json:"list"`
}

type PostWithNodeAndCommentsReq struct {
	Id            uint `json:"id" dc:"主题id"`
	PageSizeInput      // 回复分页
	SeeUserId     uint `json:"see_user_id" dc:"查看主题的用户id,用户可以忽略回复,所以根据用户id来查询forum_thanks_or_shield_or_collect_content_relation表，获取到忽略的回复"`
}

type Replies struct {
	gmeta.Meta `orm:"table:forum_replies"`
	PageSizeOutput
	List []*ReplyWithPostItem `json:"list"`
}

type PostWithNodeAndCommentsRes struct {
	Post
	Node    Node    `orm:"with:id=node_id" json:"node" `
	User    User    `orm:"with:id=user_id" json:"user"`
	Replies Replies `json:"replies"`
}

type PagerReq struct {
	CurrentPage    int    `json:"current_page" dc:"当前页"`
	Size           int    `json:"size" dc:"每页大小"`
	TotalRow       int    `json:"total_row" dc:"总行数"`
	Url            string `json:"url" dc:"分页url"`
	ShowPageAmount int    `json:"show_page" dc:"显示页数"`
}

type PagerRes struct {
	CurrentPage    int    `json:"current_page" dc:"当前页"`
	Size           int    `json:"size" dc:"每页大小"`
	TotalRow       int    `json:"total_row" dc:"总行数"`
	ShowPageAmount int    `json:"show_page" dc:"显示页数"`
	BeginIndex     int    `json:"begin_index" dc:"记录开始位置"`
	EndIndex       int    `json:"end_index" dc:"记录结束位置"`
	Html           string `json:"html" dc:"分页html"`
}
