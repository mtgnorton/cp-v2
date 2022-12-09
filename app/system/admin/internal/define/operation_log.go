package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type OperationLogListReq struct {
	g.Meta `path:"/operation-log-list" method:"GET" summary:"操作日志列表" tags:"系统管理"`
	*OperationLogInput
}

type OperationLogListRes struct {
	*OperationLogOutput
}

type OperationLogInput struct {
	Path            string `json:"path"`
	AdministratorId int    `json:"admin_id"`
	model.PageSizeInput
}

type OperationLogRecord struct {
	entity.AdminLog
	AdminName string `json:"admin_name"`
}

type OperationLogOutput struct {
	List             []OperationLogRecord `json:"list"`
	AdministratorMap map[uint]string      `json:"administrator_map"`
	Page             int                  `json:"page"`
	Size             int                  `json:"size"`
	Total            int                  `json:"total"`
}
