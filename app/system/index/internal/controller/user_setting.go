package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var UserSetting = userSetting{}

type userSetting struct {
}

// Index 用户设置首页
func (u *userSetting) Index(ctx context.Context, req *define.UserSettingPageReq) (res *define.UserSettingPageRes, err error) {
	user, err := service.User.GetTemplateShow(ctx)
	if err != nil {
		return
	}
	userFull, err := shared.User.Info(ctx, &model.UserInfoInput{
		Id: user.Id,
	})
	if err != nil {
		return
	}
	if req.Type == "" {
		req.Type = "index"
	}
	service.View().Render(ctx, define.View{
		Title:    "个人设置",
		User:     user,
		Template: "user/setting.html",
		Data: g.Map{
			"userFull": userFull,
			"type":     req.Type,
		},
	})
	return
}

// UpdateInfo 更新用户信息
func (u *userSetting) UpdateInfo(ctx context.Context, req *define.UserUpdateInfoReq) (res *define.UserUpdateInfoRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.User.UpdateInfo(ctx, user.Id, req)
	return
}

// UpdatePassword 更新用户密码
func (u *userSetting) UpdatePassword(ctx context.Context, req *define.UserUpdatePasswordReq) (res *define.UserUpdatePasswordRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = shared.User.UpdatePassword(ctx, user.Username, req.Password, req.OldPassword)
	return
}

// UpdateEmail 更新用户邮箱
func (u *userSetting) UpdateEmail(ctx context.Context, req *define.UserUpdateEmailReq) (res *define.UserUpdateEmailRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.User.UpdateEmail(ctx, user.Username, req)
	return
}

// UploadAvatar  上传头像
func (u *userSetting) UploadAvatar(ctx context.Context, req *define.UserUploadAvatarReq) (res *define.UserUploadAvatarRes, err error) {
	user, err := service.FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return
	}
	err = service.User.UploadAvatar(ctx, user.Username, req)
	return
}
