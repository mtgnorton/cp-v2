package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var OperationLog = &operationLogController{}

type operationLogController struct{}

func (o *operationLogController) List(ctx context.Context, req *define.OperationLogListReq) (res *define.OperationLogListRes, err error) {
	res = &define.OperationLogListRes{}
	res.OperationLogOutput, err = service.OperationLogService.List(ctx, req.OperationLogInput)
	return
}
