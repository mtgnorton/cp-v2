package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AdministratorLogoutReq struct {
	g.Meta `path:"/logout" method:"post" summary:"登录退出" tags:"全局"`
}

type AdministratorLogoutRes struct {
}

/*添加管理员*/

type AdministratorStoreReq struct {
	g.Meta `path:"/administrator-store" method:"post" summary:"添加管理员" tags:"管理员管理"`
	AdministratorStoreInput
}

type AdministratorStoreRes struct {
}

type AdministratorStoreInput struct {
	model.AdministratorBase
}

/*管理员list*/
type AdministratorListReq struct {
	g.Meta `path:"/administrator-list" method:"get" summary:"管理员列表" tags:"管理员管理"`
	AdministratorListInput
}
type AdministratorListRes struct {
	AdministratorListOutput
}
type AdministratorListInput struct {
	Username  string `dc:"用户名查询"`
	Status    string `dc:"状态查询" v:"in:normal,disabled#状态错误"`
	BeginTime string `dc:"创建开始时间"`
	EndTime   string `dc:"创建结束时间"`
	model.PageSizeInput
	model.OrderFieldDirectionInput
}

type AdministratorListOutput struct {
	List []*entity.Administrator `json:"list"`
	model.PageSizeOutput
}

/*获取用户信息*/
type AdministratorInfoInput struct {
	Id uint `json:"id" dc:"参数可选"`
}

type AdministratorInfoOutput struct {
	User struct {
		model.AdministratorBase
	} `json:"user"`
	Roles []*entity.Role `json:"roles"`
}

type AdministratorInfoReq struct {
	g.Meta `path:"/administrator-info" method:"get" summary:"管理员新增或更新获取相关信息" tags:"管理员管理"`
	AdministratorInfoInput
}
type AdministratorInfoRes struct {
	*AdministratorInfoOutput
}

/*管理员更新post*/
type AdministratorUpdateReq struct {
	g.Meta `path:"/administrator-update" method:"put" summary:"管理员更新" tags:"管理员管理"`
	AdministratorUpdateInput
}

type AdministratorUpdateRes struct {
}

type AdministratorUpdateInput struct {
	Id uint `v:"required#管理员id不能为空" d:"34"`
	model.AdministratorBase
}

/*管理员删除post*/
type AdministratorDestroyReq struct {
	g.Meta `path:"/administrator-destroy" method:"delete" summary:"管理员删除" tags:"管理员管理"`
	AdministratorDestroyInput
}

type AdministratorDestroyRes struct {
}

type AdministratorDestroyInput struct {
	Ids []uint `v:"required#管理员id不能为空" d:"34" json:"ids"`
}

type AdministratorGetLoggedInfoReq struct {
	g.Meta `path:"/administrator-get-logged-info" method:"get" summary:"获取已经登录的管理员信息" tags:"全局"`
}
type AdministratorGetLoggedInfoRes struct {
	*model.AdministratorSummary
}
