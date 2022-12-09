package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Node = node{}

type node struct {
}

func (n *node) List(ctx context.Context, req *define.NodeListReq) (res *define.NodeListRes, err error) {
	res = &define.NodeListRes{}
	res.NodeListOutput, err = shared.Node.List(ctx, &req.NodeListInput)
	return
}

func (n *node) Store(ctx context.Context, req *define.NodeStoreReq) (res *define.NodeStoreRes, err error) {
	err = service.Node.Store(ctx, req.NodeStoreInput)
	return
}

func (n *node) Update(ctx context.Context, req *define.NodeUpdateReq) (res *define.NodeUpdateRes, err error) {
	err = service.Node.Update(ctx, req.NodeUpdateInput)
	return
}

func (n *node) Destroy(ctx context.Context, req *define.NodeDestroyReq) (res *define.NodeDestroyRes, err error) {
	err = service.Node.Destroy(ctx, req.Id)
	return
}

func (n *node) UploadImg(ctx context.Context, req *define.NodeUploadImgReq) (res *define.NodeUploadImgRes, err error) {
	res = &define.NodeUploadImgRes{}
	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.NodeImage,
		UploadPosition: model.UPLOAD_POSITION_BACKEND,
		Dir:            "node",
		UploadType:     model.UPLOAD_TYPE_IMAGE,
	})
	res.Url = out.RelativePath

	return
}

// DownloadImg 下载远程图片
func (n *node) DownloadImg(ctx context.Context, req *define.NodeDownloadImgReq) (res *define.NodeDownloadImgRes, err error) {
	res = &define.NodeDownloadImgRes{}
	out, err := shared.Download.Image(ctx, &model.DownloadImageInput{
		Url: req.ImageUrl,
		Dir: "node",
	})
	res.Url = out.RelativePath
	return
}

func (n *node) Info(ctx context.Context, req *define.NodeInfoReq) (res *define.NodeInfoRes, err error) {
	res = &define.NodeInfoRes{}
	res.Nodes, err = shared.Node.Detail(ctx, model.NodeDetailInput{
		Id: req.Id,
	})
	return
}
