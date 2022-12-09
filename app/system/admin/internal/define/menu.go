package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

/*列表*/
type MenuListReq struct {
	g.Meta `path:"/menu-list" method:"get" summary:"菜单列表" tags:"菜单管理"`
	MenuListInput
}
type MenuListRes struct {
	*MenuListOutput
}

type MenuListInput struct {
	Name   string `json:"name"`
	Status string `v:"in:normal,disabled" json:"status"`
}

type MenuListOutput struct {
	List []*entity.AdminMenu
}

/*新增*/
type MenuStoreReq struct {
	g.Meta `path:"/menu-store" method:"post" summary:"添加菜单" tags:"菜单管理"`
	*MenuStoreInput
}
type MenuStoreRes struct{}

type MenuStoreInput struct {
	model.MenuBase
}

/*更新*/
type MenuInfoReq struct {
	g.Meta `path:"/menu-info" method:"get" summary:"更新菜单get" tags:"菜单管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要修改角色"`
}

type MenuInfoRes struct {
	*MenuInfoOutput
}

type MenuInfoOutput struct {
	Id uint `json:"id"`
	model.MenuBase
}

type MenuUpdateReq struct {
	g.Meta `path:"/menu-update" method:"put" summary:"更新菜单post" tags:"菜单管理"`
	*MenuUpdateInput
}

type MenuUpdateRes struct{}

type MenuUpdateInput struct {
	Id uint
	model.MenuBase
}

type MenuDestroyReq struct {
	g.Meta `path:"/menu-destroy" method:"delete" summary:"删除菜单" tags:"菜单管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择需要删除角色"`
}

type MenuDestroyRes struct{}

type RoutesReq struct {
	g.Meta `path:"/routes" method:"get" summary:"获取前台菜单路由组" tags:"全局"`
}

type RoutesRes struct {
	FrontRoutes []model.FrontRoute `json:"front_routes"`
}

type SiteInfoReq struct {
	g.Meta `path:"/site-info" method:"get" summary:"获取站点信息" tags:"全局"`
}
type SiteInfoRes struct {
	Logo     string `json:"logo"`
	SiteName string `json:"site_name"`
}
