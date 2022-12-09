package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type EmailRecordListReq struct {
	g.Meta   `path:"/email-record-list" method:"get" tags:"邮件记录列表" summary:"邮件记录列表"`
	Email    string
	Username string
	model.PageSizeInput
}
type EmailRecordListRes struct {
	List []*entity.EmailRecords `json:"list"`
	model.PageSizeOutput
}

type EmailRecordDestroyReq struct {
	g.Meta `path:"/email-record-destroy" method:"delete" tags:"邮件记录列表" summary:"删除邮件记录"`
	Id     int64 `p:"id" v:"required|min:1#id必须｜id必须大于0" dc:"id"`
}

type EmailRecordDestroyRes struct {
}
