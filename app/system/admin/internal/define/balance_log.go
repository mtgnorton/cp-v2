package define

import (
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/frame/g"
)

type BalanceLogListReq struct {
	g.Meta `path:"/balance-log-list" method:"get" tags:"余额变动列表" summary:"余额变动列表"`
	model.BalanceChangeLogListInput
}
type BalanceLogListRes struct {
	*model.BalanceChangeLogListOutput
}

type BalanceLogDestroyReq struct {
	g.Meta `path:"/balance-log-destroy" method:"delete" tags:"余额变动列表" summary:"删除余额变动"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type BalanceLogDestroyRes struct {
}
