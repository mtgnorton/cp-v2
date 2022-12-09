package model

import (
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gmeta"
	"github.com/gogf/gf/v2/frame/g"
)

type User struct {
	gmeta.Meta `orm:"table:forum_users"`
	entity.Users
}

const (
	USER_STATUS_DISABLE_LOGIN = 1 << iota //1
	USER_STATUS_DISABLE_POST              //2
	USER_STATUS_DISABLE_REPLY             //4
	USER_STATUS_NO_ACTIVE                 //8
)

//userId uint, amount int, changeType model.BalanceChangeType, relationId uint, remark string
type UserChangeBalanceInput struct {
	UserId     uint
	Amount     int // 大于0增加，小于0减少
	ChangeType BalanceChangeType
	RelationId uint
	Remark     string
}

type UserSummary struct {
	Id               uint   `json:"id" gconv:"id"`
	Username         string `json:"username" gconv:"username"`
	Email            string `json:"email" gconv:"email"`
	Avatar           string `json:"avatar"`
	Status           uint   `json:"status" gconv:"status"`
	CreatedAt        gtime.Time
	CollectNodeCount uint `json:"collect_node_count" gconv:"collect_node_count"`
	CollectPostCount uint `json:"collect_post_count" gconv:"collect_post_count"`
	FollowUserCount  uint `json:"collect_user_count" gconv:"collect_user_count"`
	Balance          int  `json:"balance" gconv:"balance"`
	NoReadCount      int
}

type UserInfoWithoutPass struct {
	Id                  uint        `json:"id"                     ` //
	Username            string      `json:"username"               ` // 用户名
	Email               string      `json:"email"                  ` // email
	Description         string      `json:"description"            ` // 简介
	Avatar              string      `json:"avatar"                 ` // 头像地址
	Status              string      `json:"status"                 ` // 状态
	PostsAmount         uint        `json:"posts_amount"           ` // 创建主题次数
	ReplyAmount         uint        `json:"reply_amount"           ` // 回复次数
	ShieldedAmount      uint        `json:"shielded_amount"        ` // 被屏蔽次数
	FollowByOtherAmount uint        `json:"follow_by_other_amount" ` // 被关注次数
	TodayActivity       uint        `json:"today_activity"         ` // 今日活跃度
	Remark              string      `json:"remark"                 ` // 备注
	LastLoginIp         string      `json:"last_login_ip"          ` // 最后登陆IP
	LastLoginTime       *gtime.Time `json:"last_login_time"        ` // 最后登陆时间
	CreatedAt           string      `json:"created_at"             ` // 注册时间
	UpdatedAt           gtime.Time  `json:"updated_at"             ` // 更新时间
	DeletedAt           *gtime.Time `json:"deleted_at"             ` // 删除时间
}

type UserRegisterInput struct {
	Username    string `v:"required|passport#请输入用户名|用户名只能字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"用户名" d:"username" json:"username"`
	Email       string `v:"required|email#请输入邮箱|邮箱格式错误" dc:"邮箱" d:"email" json:"email"`
	Password    string `v:"required|password#请输入密码|密码长度需要在长度在6~18之间" dc:"密码" d:"password" json:"password"`
	Password2   string `v:"required|same:password#请输入确认密码|两次密码不一致" dc:"确认密码" d:"password2" json:"password2"`
	Description string `json:"description"            ` // 简介
	Avatar      string `json:"avatar"                 ` // 头像地址
	Status      int    `json:"status"                 ` // 状态：disable_login | disable_posts | disable_reply
	Remark      string `json:"remark"                 ` // 备注

	IsNeedEmail bool // 是否需要邮箱验证
}

type UserListInput struct {
	g.Meta   `path:"/user-list" method:"get" summary:"用户列表" tags:"用户管理"`
	Username string `json:"username"`
	Status   string //status 1:禁止登陆 2:禁止发帖 4:禁止回复, 0:正常 为空时不筛选
	OrderFieldDirectionInput
	PageSizeInput
}

type UserListOutput struct {
	List []*entity.Users `json:"list"`
	PageSizeOutput
}

type UserInfoInput struct {
	Id       uint `json:"id"`
	Username string
}
