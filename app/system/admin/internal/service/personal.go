package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Personal = personalService{}

type personalService struct {
	uploadAvatarPath string
}

func (p *personalService) LoginInfo(ctx context.Context) (out *define.PersonalLoginInfoOutput, err error) {
	out = &define.PersonalLoginInfoOutput{}
	out.Data, err = shared.Config.Gets(ctx, "backend", "is_open_verify_captcha")
	return
}

func (p *personalService) Login(ctx context.Context, in define.PersonalLoginInput) (out *define.PersonalLoginOutput, err error) {
	out = &define.PersonalLoginOutput{}

	isVerifyCaptcha, err := shared.Config.Get(ctx, shared.Config.BACKEND, "is_open_verify_captcha")
	if err != nil {
		return
	}
	if isVerifyCaptcha.Bool() && !shared.Captcha.Verify(ctx, in.Code, in.CaptchaId) {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "验证码错误")
		return
	}

	entity, err := Administrator.GetUserByPassportAndPassword(
		ctx,
		in.Username,
		utility.EncryptPassword(in.Username, in.Password),
	)
	if err != nil {
		return
	}

	if entity == nil {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
	}

	adminSummary, err := Administrator.GetAdministratorSummary(ctx, entity.Id)
	if err != nil {
		return
	}

	token, err := AdminTokenInstance.LoadConfig().TokenHandler.GenerateAndSaveData(ctx, in.Username, adminSummary)
	if err != nil {
		return
	}
	out.Token = token
	shared.Context.SetUser(ctx, adminSummary)

	return
}

func (p *personalService) Info(ctx context.Context, id uint) (out define.PersonalInfoOutput, err error) {
	out = define.PersonalInfoOutput{}
	err = dao.Administrator.Ctx(ctx).WherePri(id).Scan(&out)
	if err != nil {
		return
	}
	if out.Id == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "用户不存在")
	}
	return

}

func (p *personalService) Avatar(ctx context.Context, id uint, in *define.PersonAvatarInput) (out define.PersonAvatarOutput, err error) {

	out = define.PersonAvatarOutput{}

	uploadRes, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           in.AvatarFile,
		UploadPosition: model.UPLOAD_POSITION_BACKEND,
		Dir:            "avatar/" + gconv.String(id),
		UploadType:     model.UPLOAD_TYPE_IMAGE,
	})
	if err != nil {
		return
	}
	_, err = dao.Administrator.Ctx(ctx).WherePri(id).Update(g.Map{
		dao.Administrator.Columns.Avatar: uploadRes.RelativePath,
	})
	if err != nil {
		return
	}
	updatedAdministrator, err := Administrator.GetAdministratorSummary(ctx, id)
	if err != nil {
		return
	}
	err = AdminTokenInstance.TokenHandler.UpdateData(ctx, updatedAdministrator.Username, updatedAdministrator)
	if err != nil {
		return
	}

	out.AvatarUrl = uploadRes.RelativePath
	return
}
func (p *personalService) Update(ctx context.Context, administrator *model.AdministratorSummary, in *define.PersonalUpdateInput) (err error) {

	inputOldPassword := utility.EncryptPassword(administrator.Username, in.OldPassword)

	if inputOldPassword != administrator.Password {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "旧密码不正确")
		return
	}
	updateData := g.Map{
		dao.Administrator.Columns.Nickname: in.Nickname,
	}
	if in.Password != "" {
		updateData[dao.Administrator.Columns.Password] = utility.EncryptPassword(administrator.Username, in.Password)
	}
	result, err := dao.Administrator.Ctx(ctx).WherePri(administrator.Id).Fields().Update(updateData)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()

	if err != nil {
		return
	}
	if row == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "没有修改数据")
		return
	}

	updatedAdministrator, err := Administrator.GetAdministratorSummary(ctx, administrator.Id)

	if err != nil {
		return err
	}
	err = AdminTokenInstance.TokenHandler.UpdateData(ctx, updatedAdministrator.Username, updatedAdministrator)

	return
}
