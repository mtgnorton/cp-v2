package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"

	"github.com/gogf/gf/v2/frame/g"
)

var OperationLogService = operationLogService{}

type operationLogService struct {
}

func (o *operationLogService) List(ctx context.Context, in *define.OperationLogInput) (out *define.OperationLogOutput, err error) {

	d := dao.AdminLog.Ctx(ctx)

	condition := g.Map{}
	out = &define.OperationLogOutput{}

	if in.AdministratorId != 0 {
		d = d.Where(dao.AdminLog.Columns.AdministratorId, in.AdministratorId)
	}

	if in.Path != "" {
		d = d.WhereLike(dao.AdminLog.Columns.Path, fmt.Sprintf("%%%s%%", in.Path))
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count(condition)
	if err != nil {
		return
	}

	err = d.Order(dao.AdminLog.Columns.Id+" desc").Page(in.Page,
		in.Size).Scan(&out.List)
	if err != nil {
		return nil, err
	}
	records, err := dao.Administrator.Ctx(ctx).Fields(dao.Administrator.Columns.Id,
		dao.Administrator.Columns.Username).All()

	if err != nil {
		return
	}
	out.AdministratorMap = make(map[uint]string)
	for _, record := range records {
		out.AdministratorMap[record["id"].Uint()] = record["username"].String()
	}
	for key, record := range out.List {
		out.List[key].AdminName = out.AdministratorMap[record.AdministratorId]
	}
	return
}
