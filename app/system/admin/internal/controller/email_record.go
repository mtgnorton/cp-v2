package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var EmailRecord = emailRecord{}

type emailRecord struct {
}

// List 获取邮件记录列表
func (e *emailRecord) List(ctx context.Context, req *define.EmailRecordListReq) (res *define.EmailRecordListRes, err error) {
	res, err = service.EmailRecord.List(ctx, req)
	return
}

// Destroy 删除邮件记录
func (e *emailRecord) Destroy(ctx context.Context, req *define.EmailRecordDestroyReq) (res *define.EmailRecordDestroyRes, err error) {
	err = service.EmailRecord.Destroy(ctx, req)
	return
}
