package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Menu = menu{}

type menu struct{}

func (m *menu) List(ctx context.Context, req *define.MenuListReq) (res *define.MenuListRes, err error) {
	res = &define.MenuListRes{}
	res.MenuListOutput, err = service.Menu.List(ctx, &req.MenuListInput)
	return
}

func (m *menu) Store(ctx context.Context, req *define.MenuStoreReq) (res *define.MenuStoreRes, err error) {
	if req.Identification != "" && string(req.Identification[0]) != "/" {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "权限标识必须以/开头")
		return
	}
	err = service.Menu.Store(ctx, req.MenuStoreInput)
	return
}
func (m *menu) Info(ctx context.Context, req *define.MenuInfoReq) (res *define.MenuInfoRes, err error) {
	res = &define.MenuInfoRes{}
	res.MenuInfoOutput, err = service.Menu.Info(ctx, req.Id)
	return
}

func (m *menu) Update(ctx context.Context, req *define.MenuUpdateReq) (res *define.MenuUpdateRes, err error) {
	if req.Identification != "" && string(req.Identification[0]) != "/" {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "权限标识必须以/开头")
		return
	}
	err = service.Menu.Update(ctx, req.MenuUpdateInput)
	return
}

func (m *menu) Destroy(ctx context.Context, req *define.MenuDestroyReq) (res *define.MenuDestroyRes, err error) {
	err = service.Menu.Destroy(ctx, req.Id)
	return
}
