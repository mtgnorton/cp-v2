package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/database/gdb"
)

var User = user{
	noModifyPassword: "******", //当上传此密码时，代表密码无需修改

}

type user struct {
	noModifyPassword string
}

func (u *user) Store(ctx context.Context, in *define.UserStoreReq) (out *define.UserStoreRes, err error) {
	out = &define.UserStoreRes{}
	_, err = shared.User.Register(ctx, &in.UserRegisterInput)
	if err != nil {
		return out, response.NewError(err.Error())
	}
	return
}

func (u *user) Update(ctx context.Context, in *define.UserUpdateReq) (out *define.UserUpdateRes, err error) {
	d := dao.Users.Ctx(ctx)
	// 判断记录是否存在
	err = dao.Users.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		exist, err := d.WhereNot(dao.Users.Columns().Id, in.Id).
			Where(dao.Users.Columns().Username, in.Username).
			Count()

		if err != nil {
			return err
		}
		if exist > 0 {
			return response.NewError("用户名已经存在")
		}

		oldPassword, err := d.WherePri(in.Id).Value(dao.Users.Columns().Password)

		if err != nil {
			return err
		}

		if in.Password == u.noModifyPassword {
			in.Password = oldPassword.String()
		}
		if in.Password != oldPassword.String() {
			in.Password = utility.EncryptPassword(in.Username, in.Password)
		}
		_, err = d.WherePri(in.Id).Update(in)

		return err
	})
	return
}

func (u *user) Info(ctx context.Context, in *define.UserInfoReq) (out *define.UserInfoRes, err error) {
	out = &define.UserInfoRes{}
	err = dao.Users.Ctx(ctx).WherePri(in.Id).Scan(&out)
	return
}

func (u *user) Destroy(ctx context.Context, in *define.UserDestroyReq) (out *define.UserDestroyRes, err error) {
	out = &define.UserDestroyRes{}
	_, err = dao.Users.Ctx(ctx).WherePri(in.Id).Delete()
	return
}
