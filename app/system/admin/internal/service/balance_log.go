package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
)

var BalanceLog = balanceLog{}

type balanceLog struct {
}

// Destroy 删除余额变动记录
func (e *balanceLog) Destroy(ctx context.Context, req *define.BalanceLogDestroyReq) (err error) {
	d := dao.BalanceChangeLog.Ctx(ctx)
	_, err = d.Delete(dao.BalanceChangeLog.Columns().Id, req.Id)

	return
}
