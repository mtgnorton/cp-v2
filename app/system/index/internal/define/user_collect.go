package define

import "github.com/gogf/gf/v2/frame/g"

type UserCollectNodePageReq struct {
	g.Meta `path:"/my/nodes" method:"get" tags:"收藏|关注|屏蔽|感谢" summary:"收藏节点列表页面"`
}

type UserCollectNodePageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserCollectPostPageReq struct {
	g.Meta `path:"/my/posts" method:"get" tags:"收藏|关注|屏蔽|感谢" summary:"收藏主题列表页面"`
}
type UserCollectPostPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserFollowUserPageReq struct {
	g.Meta `path:"/my/following" method:"get" tags:"收藏|关注|屏蔽|感谢" summary:"关注用户列表页面"`
}

type UserFollowUserPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserToggleCollectNodeReq struct {
	g.Meta `path:"/node/toggle-collect" method:"post" tags:"收藏|关注|屏蔽|感谢" summary:"收藏节点"`
	NodeId uint `json:"node_id" v:"required#节点ID不能为空"`
}

type UserToggleCollectNodeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserToggleFollowUserReq struct {
	g.Meta   `path:"/user/follow" method:"post" tags:"收藏|关注|屏蔽|感谢" summary:"关注|取消 用户"`
	TargetId uint `p:"target_id"`
}

type UserToggleFollowUserRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
type UserToggleShieldUserReq struct {
	g.Meta   `path:"/user/shield" method:"post" tags:"收藏|关注|屏蔽|感谢" summary:"屏蔽|取消 屏蔽"`
	TargetId uint `p:"target_id"`
}

type UserToggleShieldUserRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
