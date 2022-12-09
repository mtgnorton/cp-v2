package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AssociationListReq struct {
	g.Meta `path:"/association-list" method:"get" tags:"收藏|关注|屏蔽|感谢列表" summary:"收藏|关注|屏蔽|感谢列表"`
	*model.AssociationListInput
}

type AssociationListRes struct {
	*model.AssociationListOutput
}
