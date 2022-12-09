package model

import "gf-admin/app/model/entity"

const (
	REPLY_STATUS_NORMAL   = 1
	REPLY_STATUS_NO_AUDIT = 0
	REPLY_STATUS_SHIELD   = -1
)

type ReplyListInput struct {
	PostId        uint   // 可以为0,所属主题ID
	UserId        uint   // 可以为0,所属用户ID
	Username      string //
	SeeUserId     uint   //可以为0,查看回复的用户id，用户可能会屏蔽某些回复
	Keyword       string //为空不筛选
	Status        string //0 未审核 1 正常  为空不筛选
	WithPost      bool   //是否需要带上主题信息
	PageSizeInput        // 只有Page和Size都不为0时才分页
	OrderFieldDirectionInput
}

type ReplyWithPostItem struct {
	entity.Replies
	PostUsername string `json:"post_username"`
	PostTitle    string `json:"post_title"`
	PostId       uint   `json:"post_id"`
	NodeName     string `json:"node_name"`
	NodeKeyword  string `json:"node_keyword"`
	UserAvatar   string `json:"user_avatar"`
}
type ReplyListOutput struct {
	List []*ReplyWithPostItem `json:"list"`
	PageSizeOutput
}
