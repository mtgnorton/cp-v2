package api

import (
	"context"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
)

var Login = login{}

type login struct {
}

func (n *login) Login(ctx context.Context, req *define.LoginReq) (res *define.LoginRes, err error) {
	res = &define.LoginRes{}
	res.LoginOutput, err = service.UserNoNeedLogin.Login(ctx, req.LoginInput)
	return
}
