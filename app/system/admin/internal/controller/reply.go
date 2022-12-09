package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Reply = reply{}

type reply struct {
}

func (r *reply) List(ctx context.Context, req *define.ReplyListReq) (res *define.ReplyListRes, err error) {
	res = &define.ReplyListRes{}
	req.ReplyListInput.WithPost = true
	res.ReplyListOutput, err = shared.Reply.List(ctx, req.ReplyListInput)
	return
}

func (r *reply) Audit(ctx context.Context, req *define.ReplyAuditReq) (res *define.ReplyAuditRes, err error) {
	res = &define.ReplyAuditRes{}
	err = service.Reply.Audit(ctx, req.Id)
	return
}

func (r *reply) Update(ctx context.Context, req *define.ReplyUpdateReq) (res *define.ReplyUpdateRes, err error) {
	res = &define.ReplyUpdateRes{}
	err = service.Reply.Update(ctx, req)
	return
}
func (r *reply) Destroy(ctx context.Context, req *define.ReplyDestroyReq) (res *define.ReplyDestroyRes, err error) {
	err = shared.Reply.Destroy(ctx, req.Id)
	return
}
