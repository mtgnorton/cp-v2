package controller

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
)

var Association = association{}

type association struct {
}

// List 获取关联列表
func (a *association) List(ctx context.Context, req *define.AssociationListReq) (res *define.AssociationListRes, err error) {
	res = &define.AssociationListRes{}
	res.AssociationListOutput, err = shared.Association.List(ctx, req.AssociationListInput)
	return
}
