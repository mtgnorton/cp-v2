package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var Index = index{}

type index struct {
}

// Index 首页
func (i *index) Index(ctx context.Context, req *define.IndexReq) (res *define.IndexRes, err error) {
	res, err = service.Index.Statistics(ctx)
	return
}
