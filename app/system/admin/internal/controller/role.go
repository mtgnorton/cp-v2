package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var Role = role{}

type role struct {
}

func (r *role) List(ctx context.Context, req *define.RoleListReq) (res *define.RoleListRes, err error) {
	res = &define.RoleListRes{}

	res.RoleListOutput, err = service.Role.List(ctx, &req.RoleListInput)

	return
}

func (r *role) Store(ctx context.Context, req *define.RoleStoreReq) (res *define.RoleStoreRes, err error) {
	var in *define.RoleStoreInput
	err = gconv.Scan(req, &in)
	if err != nil {
		return
	}
	err = service.Role.Store(ctx, in)
	return
}

//
func (r *role) Info(ctx context.Context, req *define.RoleInfoReq) (res *define.RoleInfoRes, err error) {
	res = &define.RoleInfoRes{}
	res.RoleInfoOutput, err = service.Role.Info(ctx, &req.RoleInfoInput)
	return
}

func (r *role) Update(ctx context.Context, req *define.RoleUpdateReq) (res *define.RoleUpdateRes, err error) {
	err = service.Role.Update(ctx, req.RoleUpdateInput)
	return
}

func (r *role) Destroy(ctx context.Context, req *define.RoleDestroyReq) (res *define.RoleDestroyRes, err error) {
	err = service.Role.Destroy(ctx, req.RoleDestroyInput)
	return
}
