package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleStoreReq struct {
	g.Meta `path:"/role-store" method:"post" summary:"添加角色" tags:"角色管理"`
	RoleStoreInput
}

type RoleStoreRes struct {
}

type RoleStoreInput struct {
	model.RoleBase
}

/*角色列表*/
type RoleListReq struct {
	g.Meta `path:"/role-list" method:"get" summary:"角色列表" tags:"角色管理"`
	RoleListInput
}

type RoleListRes struct {
	*RoleListOutput
}

type RoleListInput struct {
	Name           string ` dc:"角色名称"`
	Identification string `dc:"角色标志性符"`
	Status         string ` dc:"状态，normal正常，disabled 禁用"`
	model.PageSizeInput
}

type RoleListOutput struct {
	List []*entity.Role `json:"list"`
	model.PageSizeOutput
}

/*角色更新Get*/
type RoleInfoReq struct {
	g.Meta `path:"/role-info" method:"get" summary:"更新角色" tags:"角色管理"`
	RoleInfoInput
}

type RoleInfoRes struct {
	*RoleInfoOutput
}

type RoleInfoInput struct {
	Id int `v:"required#id参数必须传递" json:"id"`
}

type RoleInfoOutput struct {
	Role       model.RoleBase        `json:"role"`
	MenusTrees []*model.MenuMiniTree `json:"permission_tree"`
}

/*角色更新post*/
type RoleUpdateReq struct {
	g.Meta `path:"/role-update" method:"put" summary:"更新角色" tags:"角色管理"`
	*RoleUpdateInput
}

type RoleUpdateRes struct{}

type RoleUpdateInput struct {
	Id uint
	model.RoleBase
}

type RoleDestroyReq struct {
	g.Meta `path:"/role-destroy" method:"delete" summary:"删除角色" tags:"角色管理"`
	*RoleDestroyInput
}

type RoleDestroyRes struct{}

type RoleDestroyInput struct {
	Ids []uint `v:"required#ids参数必须" json:"ids"`
}
