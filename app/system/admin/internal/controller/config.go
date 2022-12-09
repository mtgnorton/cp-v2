package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"

	"github.com/gogf/gf/v2/os/gfile"
)

var Config = configApi{}

type configApi struct {
}

func (c *configApi) List(ctx context.Context, req *define.ConfigListReq) (resp *define.ConfigListRes, err error) {
	resp = new(define.ConfigListRes)
	resp.ConfigListOutput, err = service.Config.GetModules(ctx, req.ConfigListInput)
	return
}

func (c *configApi) Update(ctx context.Context, req *define.ConfigUpdateReq) (resp *define.ConfigUpdateRes, err error) {
	resp = new(define.ConfigUpdateRes)
	err = service.Config.Update(ctx, req.ConfigUpdateInput)

	return
}

func (n *configApi) UploadLogo(ctx context.Context, req *define.UploadLogoReq) (res *define.UploadLogoRes, err error) {
	res = &define.UploadLogoRes{}
	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.Logo,
		UploadPosition: model.UPLOAD_POSITION_BACKEND,
		Dir:            "logo",
		UploadType:     model.UPLOAD_TYPE_IMAGE,
	})
	if err != nil {
		return
	}
	res.Url = out.RelativePath

	return
}

func (n *configApi) UploadFavicon(ctx context.Context, req *define.UploadFaviconReq) (res *define.UploadFaviconRes, err error) {
	res = &define.UploadFaviconRes{}
	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.Favicon,
		UploadPosition: model.UPLOAD_POSITION_BACKEND,
		Dir:            "favicon",
		UploadType:     model.UPLOAD_TYPE_IMAGE,
	})
	err = gfile.Copy(out.AbsolutePath, "public/favicon.ico")
	if err != nil {
		return
	}
	res.Url = out.RelativePath

	return
}

func (n *configApi) UploadDefaultAvatar(ctx context.Context, req *define.UploadDefaultAvatarReq) (res *define.UploadDefaultAvatarRes, err error) {
	res = &define.UploadDefaultAvatarRes{}
	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.Avatar,
		UploadPosition: model.UPLOAD_POSITION_FRONTED,
		Dir:            "avatar",
		UploadType:     model.UPLOAD_TYPE_IMAGE,
		SaveFilename:   "default",
	})
	res.Url = out.RelativePath

	return
}
