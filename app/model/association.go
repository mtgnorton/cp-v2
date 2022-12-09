package model

import "gf-admin/app/model/entity"

const (
	//感谢主题:thanks_posts,感谢回复: thanks_reply,屏蔽主题: shield_posts,屏蔽回复: shield_reply,收藏主题:collect_posts,收藏节点: collect_nodes,follow_user:关注用户,shield_user:屏蔽用户

	AssociationTypeThanksPost  = "thanks_post"
	AssociationTypeThanksReply = "thanks_reply"
	AssociationTypeShieldPost  = "shield_post"
	AssociationTypeShieldReply = "shield_reply"
	AssociationTypeCollectPost = "collect_post"
	AssociationTypeCollectNode = "collect_node"
	AssociationTypeFollowUser  = "follow_user"
	AssociationTypeShieldUser  = "shield_user"
)

type AssociationListInput struct {
	Username       string
	Type           string
	TargetId       uint
	TargetUsername string
	AdditionalId   uint
	PageSizeInput
}
type AssociationItem struct {
	entity.Association
	TargetContent string `json:"target_content"`
}
type AssociationListOutput struct {
	List []*AssociationItem `json:"list"`
	PageSizeOutput
}
