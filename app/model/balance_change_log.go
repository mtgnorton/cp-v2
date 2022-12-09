package model

import "gf-admin/app/model/entity"

type BalanceChangeType string

const (
	BALANCE_CHANGE_TYPE_LOGIN                  BalanceChangeType = "login"
	BALANCE_CHANGE_TYPE_REGSITER               BalanceChangeType = "register"
	BALANCE_CHANGE_TYPE_ESTABLISH_POST_DEDUCT  BalanceChangeType = "establish_post_deduct"
	BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_DEDUCT BalanceChangeType = "establish_reply_deduct"
	BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_REWARD BalanceChangeType = "establish_reply_reward"
	BALANCE_CHANGE_TYPE_THANK_REPLY_DEDUCT     BalanceChangeType = "thanks_reply_deduct"
	BALANCE_CHANGE_TYPE_THANK_REPLY_REWARD     BalanceChangeType = "thanks_reply_reward"
	BALANCE_CHANGE_TYPE_THANK_POST_DEDUCT      BalanceChangeType = "thanks_post_deduct"
	BALANCE_CHANGE_TYPE_THANK_POST_REWARD      BalanceChangeType = "thanks_post_reward"
	BALANCE_CHANGE_TYPE_ACTIVITY               BalanceChangeType = "activity"
)

type BalanceChangeLogListInput struct {
	UserId   uint
	Type     BalanceChangeType
	Username string
	PageSizeInput
}

type BalanceLogItem struct {
	entity.BalanceChangeLog
	TypeShow string `json:"type_show"`
}

type BalanceChangeLogListOutput struct {
	List []*BalanceLogItem `json:"list"`
	PageSizeOutput
}
