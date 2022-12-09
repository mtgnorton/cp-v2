package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var Administrator = administratorApi{}

type administratorApi struct {
}

func (a *administratorApi) List(ctx context.Context, req *define.AdministratorListReq) (res *define.AdministratorListRes, err error) {
	var input *define.AdministratorListInput
	res = &define.AdministratorListRes{}

	err = gconv.Scan(req, &input)
	if err != nil {
		return
	}
	output, err := service.Administrator.List(ctx, input)
	if err != nil {
		return
	}
	err = gconv.Scan(output, &res)
	return res, err
}

func (a *administratorApi) Store(ctx context.Context, req *define.AdministratorStoreReq) (res *define.AdministratorStoreRes, err error) {
	res = &define.AdministratorStoreRes{}

	err = service.Administrator.Store(ctx, &req.AdministratorStoreInput)

	return
}

func (a *administratorApi) Info(ctx context.Context, req *define.AdministratorInfoReq) (res *define.AdministratorInfoRes, err error) {
	res = &define.AdministratorInfoRes{}
	res.AdministratorInfoOutput, err = service.Administrator.Info(ctx, &req.AdministratorInfoInput)
	return
}

func (a *administratorApi) Update(ctx context.Context, req *define.AdministratorUpdateReq) (res *define.AdministratorUpdateRes, err error) {
	res = &define.AdministratorUpdateRes{}
	var input *define.AdministratorUpdateInput
	err = gconv.Scan(req, &input)
	if err != nil {
		return
	}
	err = service.Administrator.Update(ctx, input)
	return
}

func (a *administratorApi) Destroy(ctx context.Context, req *define.AdministratorDestroyReq) (res *define.AdministratorDestroyRes, err error) {
	var input *define.AdministratorDestroyInput
	err = gconv.Scan(req, &input)
	if err != nil {
		return
	}
	err = service.Administrator.Destroy(ctx, input)
	return
}
