package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var BalanceLog = balanceLog{}

type balanceLog struct {
}

// List 获取余额变动记录列表
func (e *balanceLog) List(ctx context.Context, req *define.BalanceLogListReq) (res *define.BalanceLogListRes, err error) {
	res = &define.BalanceLogListRes{}
	res.BalanceChangeLogListOutput, err = shared.BalanceChangeLog.List(ctx, &req.BalanceChangeLogListInput)
	return
}

// Destroy 删除余额变动记录
func (e *balanceLog) Destroy(ctx context.Context, req *define.BalanceLogDestroyReq) (res *define.BalanceLogDestroyRes, err error) {
	err = service.BalanceLog.Destroy(ctx, req)
	return
}
