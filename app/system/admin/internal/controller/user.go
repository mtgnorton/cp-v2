package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var User = user{}

type user struct {
}

func (u *user) List(ctx context.Context, req *define.UserListReq) (res *define.UserListRes, err error) {
	res = &define.UserListRes{}
	res.UserListOutput, err = shared.User.List(ctx, req.UserListInput)
	return
}

func (u *user) Store(ctx context.Context, req *define.UserStoreReq) (res *define.UserStoreRes, err error) {
	res, err = service.User.Store(ctx, req)
	return
}
func (u *user) Update(ctx context.Context, req *define.UserUpdateReq) (res *define.UserUpdateRes, err error) {
	res, err = service.User.Update(ctx, req)
	return
}

// Info 获取用户信息
func (u *user) Info(ctx context.Context, req *define.UserInfoReq) (res *define.UserInfoRes, err error) {
	res, err = service.User.Info(ctx, req)
	return
}

func (u *user) Destroy(ctx context.Context, req *define.UserDestroyReq) (res *define.UserDestroyRes, err error) {
	res, err = service.User.Destroy(ctx, req)
	return
}
