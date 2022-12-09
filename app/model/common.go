package model

type SiteInfo struct {
	UserCount  int `json:"user_count"`
	PostCount  int `json:"post_count"`
	ReplyCount int `json:"reply_count"`
}

const (
	STATUS_NORMAL   = "normal"
	STATUS_DISABLED = "disabled"
)
