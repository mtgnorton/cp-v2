package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var Personal = personal{}

type personal struct {
}

func (a *personal) Info(ctx context.Context, req *define.PersonalInfoReq) (res *define.PersonalInfoRes, err error) {
	res = &define.PersonalInfoRes{}
	id, err := service.AdminTokenInstance.GetAdministratorId(ctx)
	if err != nil {
		return
	}
	res.PersonalInfoOutput, err = service.Personal.Info(ctx, id)
	return
}

func (a *personal) Update(ctx context.Context, req *define.PersonalUpdateReq) (res *define.PersonalUpdateRes, err error) {
	res = &define.PersonalUpdateRes{}
	administrator, err := service.AdminTokenInstance.GetAdministrator(ctx)
	if err != nil {
		return
	}
	err = service.Personal.Update(ctx, administrator, &req.PersonalUpdateInput)
	return
}

func (a *personal) UploadAvatar(ctx context.Context, req *define.PersonalAvatarReq) (res *define.PersonAvatarRes, err error) {

	var (
		request = g.RequestFromCtx(ctx)
		file    = request.GetUploadFile("avatar_file")
	)
	if file == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的头像")
	}

	res = &define.PersonAvatarRes{}
	id, err := service.AdminTokenInstance.GetAdministratorId(ctx)
	if err != nil {
		return
	}
	res.PersonAvatarOutput, err = service.Personal.Avatar(ctx, id, &define.PersonAvatarInput{
		AvatarFile: file,
	})

	return
}
