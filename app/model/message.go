package model

import "gf-admin/app/model/entity"

type MessageType string

const (
	MESSAGE_TYPE_POST_OWNER MessageType = "post_owner"
	MESSAGE_TYPE_REPLY      MessageType = "reply"
)

type MessageListInput struct {
	RepliedUserId uint   //获取某个用户所有的消息提醒
	IsRead        string //"" 获取所有消息提醒，"1" 获取已读消息提醒，"0" 获取未读消息提醒
	PageSizeInput
}

type MessageListItem struct {
	Message *entity.Messages
	Post    *entity.Posts
	Reply   *entity.Replies
	User    *entity.Users
}

type MessageListOutput struct {
	List []*MessageListItem `json:"list"`
	PageSizeOutput
}
