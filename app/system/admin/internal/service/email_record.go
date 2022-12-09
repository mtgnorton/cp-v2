package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
)

var EmailRecord = emailRecord{}

type emailRecord struct {
}

// List 获取邮件记录列表
func (e *emailRecord) List(ctx context.Context, req *define.EmailRecordListReq) (res *define.EmailRecordListRes, err error) {
	d := dao.EmailRecords.Ctx(ctx)
	res = &define.EmailRecordListRes{}
	if req.Username != "" {
		d = d.WhereLike(dao.EmailRecords.Columns().Username, "%"+req.Username+"%")
	}
	if req.Email != "" {
		d = d.WhereLike(dao.EmailRecords.Columns().Email, "%"+req.Email+"%")
	}
	res.Page = req.Page
	res.Size = req.Size
	res.Total, err = d.Count()
	if err != nil {
		return
	}
	err = d.Page(req.Page, res.Size).OrderDesc(dao.BalanceChangeLog.Columns().Id).Scan(&res.List)

	return
}

// Destroy 删除邮件记录
func (e *emailRecord) Destroy(ctx context.Context, req *define.EmailRecordDestroyReq) (err error) {
	d := dao.EmailRecords.Ctx(ctx)
	_, err = d.Delete(dao.EmailRecords.Columns().Id, req.Id)

	return
}
