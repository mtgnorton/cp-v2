package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/os/gtime"
)

var EmailRecord = emailRecord{}

type emailRecord struct {
}

// GetUserLastEmailTime 获取某个邮箱最后一次发送邮件的时间
func (e *emailRecord) GetUserLastEmailTime(ctx context.Context, email string, emailTypes ...model.EmailType) (int64, error) {
	d := dao.EmailRecords.Ctx(ctx)

	d = d.Where(dao.EmailRecords.Columns().Email, email)
	d = d.Where(dao.EmailRecords.Columns().Error, "")
	emailType := ""
	if len(emailTypes) > 0 {
		emailType = string(emailTypes[0])
		d.Where(dao.EmailRecords.Columns().Type, emailType)
	}

	gVar, err := d.Value(dao.EmailRecords.Columns().CreatedAt)

	return gtime.NewFromStr(gVar.String()).Unix(), err
}
