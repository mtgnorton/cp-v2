package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Administrator = administratorService{
	noModifyPassword: "******", //当上传此密码时，代表密码无需修改
}

type administratorService struct {
	noModifyPassword string
}

// 管理员session登录
//func (a *administratorService) LoginSession(ctx context.Context, in define.AdministratorLoginPostInput) error {
//
//	entity, err := a.GetUserByPassportAndPassword(
//		ctx,
//		in.Username,
//		utility.EncryptPassword(in.Username, in.Password),
//	)
//	if err != nil {
//		return err
//	}
//
//	if entity == nil {
//		return gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
//	}
//	if err := Session.SetAdministrator(ctx, entity); err != nil {
//		return err
//	}
//	// 自动更新上线
//	Context.SetUser(ctx, &define.ContextAdministrator{
//		Id:       entity.Id,
//		Username: entity.Username,
//		Nickname: entity.Nickname,
//		Avatar:   entity.Avatar,
//	})
//	return nil
//}

// 管理员登录

func (a *administratorService) List(ctx context.Context, in *define.AdministratorListInput) (out *define.AdministratorListOutput, err error) {
	d := dao.Administrator.Ctx(ctx)

	if in.Username != "" {
		d = d.WhereLike(dao.Administrator.Columns.Username, fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Status != "" {
		d = d.WhereIn(dao.Administrator.Columns.Status, g.Slice{in.Status})
	}
	if in.BeginTime != "" {
		d = d.WhereGTE(dao.Administrator.Columns.CreatedAt, in.BeginTime)
	}
	if in.EndTime != "" {
		d = d.WhereLTE(dao.Administrator.Columns.CreatedAt, in.EndTime)
	}

	out = &define.AdministratorListOutput{}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, gerror.New(err.Error())
	}
	d = d.Page(in.Page, in.Size)
	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField, in.OrderDirection)
	}
	err = d.Scan(&out.List)

	return out, err
}

//新增管理员
func (a *administratorService) Store(ctx context.Context, in *define.AdministratorStoreInput) (err error) {
	err = dao.Administrator.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var administrator *entity.Administrator
		if err = gconv.Struct(in, &administrator); err != nil {
			return err
		}
		exist, err := dao.Administrator.Ctx(ctx).Where("username", administrator.Username).Count()
		if exist > 0 {
			return gerror.New("管理员用户名已经存在 ")
		}
		if err != nil {
			return err
		}
		administrator.Password = utility.EncryptPassword(administrator.Username, administrator.Password)
		id, err := dao.Administrator.Ctx(ctx).OmitEmptyData().InsertAndGetId(administrator)
		if id == 0 {
			return gerror.New("创建失败")
		}
		if err != nil {
			return err
		}
		err = a.syncUserRoles(ctx, uint(id), in.RoleIds)
		return err
	})
	return
}
func (a *administratorService) Info(ctx context.Context, in *define.AdministratorInfoInput) (out *define.AdministratorInfoOutput, err error) {
	out = &define.AdministratorInfoOutput{}
	out.Roles, err = Role.All(ctx)

	if in.Id == 0 {
		return
	}
	err = dao.Administrator.Ctx(ctx).WherePri(in.Id).Scan(&out.User)
	if err != nil {
		return
	}
	roleIds, err := dao.AdministratorRole.Ctx(ctx).Where(dao.AdministratorRole.Columns.AdministratorId, in.Id).Array(dao.AdministratorRole.Columns.RoleId)
	if err != nil {
		return
	}
	out.User.RoleIds = []uint{}
	for _, roleId := range roleIds {
		out.User.RoleIds = append(out.User.RoleIds, roleId.Uint())
	}

	return
}
func (a *administratorService) Update(ctx context.Context, in *define.AdministratorUpdateInput) (err error) {
	err = dao.Administrator.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		exist, err := dao.Administrator.Ctx(ctx).WhereNot(dao.Administrator.Columns.Id, in.Id).Where(dao.Administrator.Columns.Username, in.Username).Count()

		if err != nil {
			return err
		}
		if exist > 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "用户名已经存在")
		}

		oldPassword, err := dao.Administrator.Ctx(ctx).WherePri(in.Id).Value(dao.Administrator.Columns.Password)

		if err != nil {
			return err
		}

		if in.Password == a.noModifyPassword {
			in.Password = oldPassword.String()
		}
		if in.Password != oldPassword.String() {
			in.Password = utility.EncryptPassword(in.Username, in.Password)
		}
		_, err = dao.Administrator.Ctx(ctx).WherePri(in.Id).Update(in)
		if err != nil {
			return err
		}
		err = a.syncUserRoles(ctx, in.Id, in.RoleIds)

		return err
	})
	return err
}

func (a *administratorService) Destroy(ctx context.Context, in *define.AdministratorDestroyInput) (err error) {

	administratorNameVar, err := dao.Administrator.Ctx(ctx).WherePri(in.Ids).Value(dao.Administrator.Columns.Username)
	if err != nil {
		return err
	}
	affected, err := dao.Administrator.Ctx(ctx).WherePri(in.Ids).Delete()
	if err != nil {
		return err
	}
	rows, err := affected.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "管理员不存在")
	}

	_, err = Enforcer.DeleteUser(administratorNameVar.String())

	return
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (a *administratorService) GetUserByPassportAndPassword(ctx context.Context, username, password string) (administrator *entity.Administrator, err error) {
	err = dao.Administrator.Ctx(ctx).Where(g.Map{
		dao.Administrator.Columns.Username: username,
		dao.Administrator.Columns.Password: password,
	}).Scan(&administrator)
	return
}

func (a *administratorService) syncUserRoles(ctx context.Context, administratorId uint, roleIds []uint) (err error) {
	return dao.AdministratorRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		administratorNameVar, err := dao.Administrator.Ctx(ctx).WherePri(administratorId).Value(dao.Administrator.Columns.Username)
		if err != nil {
			return err
		}
		_, err = Enforcer.DeleteUser(administratorNameVar.String())

		if err != nil {
			return err
		}
		for _, roleId := range roleIds {
			roleNameVar, err := dao.Role.Ctx(ctx).WherePri(roleId).Value(dao.Role.Columns.Identification)
			if err != nil {
				return err
			}
			_, err = Enforcer.AddRoleForUser(administratorNameVar.String(), roleNameVar.String())
			if err != nil {
				return err
			}
		}

		_, err = dao.AdministratorRole.Ctx(ctx).Where(dao.AdministratorRole.Columns.AdministratorId, administratorId).Delete()
		if err != nil {
			return err
		}
		m := g.Slice{}
		for _, roleId := range roleIds {
			m = append(m, g.Map{
				dao.AdministratorRole.Columns.AdministratorId: administratorId,
				dao.AdministratorRole.Columns.RoleId:          roleId,
			})
		}
		if len(m) > 0 {
			_, err = dao.AdministratorRole.Ctx(ctx).Insert(m)

		}
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *administratorService) GetAdministratorSummary(ctx context.Context, administratorId uint) (admin *model.AdministratorSummary, err error) {
	admin = &model.AdministratorSummary{}
	err = dao.Administrator.Ctx(ctx).WherePri(administratorId).Scan(&admin)
	if err != nil {
		return admin, err
	}
	roles, err := Role.GetByAdministratorId(ctx, administratorId)
	if err != nil {
		return admin, err
	}
	err = gconv.Scan(roles, &admin.Roles)
	if err != nil {
		return admin, err
	}
	menus, err := Menu.GetByAdministratorId(ctx, administratorId)
	if err != nil {
		return admin, err
	}
	err = gconv.Scan(menus, &admin.Menus)

	if err != nil {
		return admin, err
	}
	if admin.Roles == nil {
		admin.Roles = []*model.RoleShow{}
	}
	if admin.Menus == nil {
		admin.Menus = []*model.MenuShow{}
	}
	return admin, err
}
