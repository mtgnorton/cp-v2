package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
	"gf-admin/app/system/admin/internal/define"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var Role = role{}

type role struct {
}

func (r *role) List(ctx context.Context, in *define.RoleListInput) (out *define.RoleListOutput, err error) {

	out = &define.RoleListOutput{}
	out.Page = in.Page
	out.Size = in.Size

	d := dao.Role.Ctx(ctx)
	if in.Name != "" {
		d = d.Where(dao.Role.Columns.Name, in.Name)
	}
	if in.Identification != "" {
		d = d.Where(dao.Role.Columns.Identification, in.Identification)
	}
	if in.Status != "" {
		d = d.Where(dao.Role.Columns.Status, in.Status)
	}
	out.Total, err = d.Count()

	if err != nil {
		return out, err
	}
	d.Page(in.Page, in.Size)

	err = d.Scan(&out.List)

	return out, err
}

func (r *role) Store(ctx context.Context, in *define.RoleStoreInput) (err error) {
	return dao.Role.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		exist, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns.Name, in.Name).WhereOr(dao.Role.Columns.Identification, in.Identification).Count()

		if err != nil {
			return err
		}
		if exist > 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "角色名或标识符重复")
		}
		roleId, err := dao.Role.Ctx(ctx).InsertAndGetId(in)
		if err != nil {
			return nil
		}
		err = r.syncRoleMenus(ctx, uint(roleId), in.MenuIds)
		return err
	})
}

func (r *role) Info(ctx context.Context, in *define.RoleInfoInput) (out *define.RoleInfoOutput, err error) {
	out = &define.RoleInfoOutput{}
	out.MenusTrees, err = Menu.MiniTree(ctx)

	if in.Id == 0 {
		return
	}
	exist, err := dao.Role.Ctx(ctx).WherePri(in.Id).Count()
	if err != nil {
		return
	}
	if exist == 0 {
		return out, gerror.NewCode(gcode.CodeInvalidParameter, "对应的角色不存在")
	}
	err = dao.Role.Ctx(ctx).WherePri(in.Id).Scan(&out.Role)
	if err != nil {
		return
	}
	menuIds, err := dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns.RoleId, in.Id).Array(dao.RoleMenu.Columns.MenuId)
	out.Role.MenuIds = make([]uint, 0)
	for _, menuId := range menuIds {
		out.Role.MenuIds = append(out.Role.MenuIds, menuId.Uint())
	}

	return
}

func (r *role) Update(ctx context.Context, in *define.RoleUpdateInput) (err error) {
	return dao.Role.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		exist, err := dao.Role.Ctx(ctx).WherePri(in.Id).Count()

		if err != nil {
			return err
		}
		if exist == 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "角色不存在")
		}
		exist, err = dao.Role.Ctx(ctx).Where("? != ? && (? = ? or ? = ?)", dao.Role.Columns.Id, in.Id, dao.Role.Columns.Name, in.Name, dao.Role.Columns.Identification, in.Identification).Count()

		if err != nil {
			return err
		}
		if exist > 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "角色名或标识符重复")
		}

		_, err = dao.Role.Ctx(ctx).Where(dao.Role.Columns.Id, in.Id).OmitEmptyData().Update(in)

		if err != nil {
			return err
		}
		err = r.syncRoleMenus(ctx, in.Id, in.MenuIds)

		return err
	})
}
func (r *role) Destroy(ctx context.Context, in *define.RoleDestroyInput) (err error) {

	dr := dao.Role.Ctx(ctx).WherePri(in.Ids)
	exist, err := dr.Count()

	if err != nil {
		return err
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "角色不存在")
	}
	_, err = dr.Delete()
	return

}

/*同步role_menu表的角色和菜单权限的关系以及casbin策略表的关系*/
func (r *role) syncRoleMenus(ctx context.Context, roleId uint, menuIds []uint) (err error) {

	err = dao.Role.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		roleIdentificationVar, err := dao.Role.Ctx(ctx).WherePri(roleId).Value(dao.Role.Columns.Identification)

		if err != nil {
			return err
		}

		_, err = Enforcer.DeletePermissionsForUser(ctx, roleIdentificationVar.String())

		if err != nil {
			return err
		}
		var menus []*entity.AdminMenu

		err = dao.AdminMenu.Ctx(ctx).WherePri(menuIds).Scan(&menus)
		if err != nil {
			return err
		}
		for _, menu := range menus {
			if menu.Identification == "" {
				continue
			}
			_, err = Enforcer.AddPermissionForUserOrRole(ctx, roleIdentificationVar.String(), gstr.ToLower(menu.Identification), gstr.ToLower(menu.Method))
			if err != nil {
				return err
			}
		}

		m := g.Slice{}
		_, err = dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns.RoleId, roleId).Delete()
		if err != nil {
			return err
		}
		for _, menuId := range menuIds {
			m = append(m, g.Map{
				dao.RoleMenu.Columns.RoleId: roleId,
				dao.RoleMenu.Columns.MenuId: menuId,
			})
		}
		_, err = dao.RoleMenu.Ctx(ctx).Insert(m)

		return err
	})
	return
}

func (r *role) All(ctx context.Context) (roles []*entity.Role, err error) {
	roles = []*entity.Role{}
	err = dao.Role.Ctx(ctx).Scan(&roles)
	return

}

func (r *role) GetByAdministratorId(ctx context.Context, administratorId uint) (roles []*entity.Role, err error) {
	roleIds, err := dao.AdministratorRole.Ctx(ctx).Where(dao.AdministratorRole.Columns.AdministratorId, administratorId).Array(dao.AdministratorRole.Columns.RoleId)
	if err != nil {
		return
	}
	roles = []*entity.Role{}
	err = dao.Role.Ctx(ctx).WherePri(roleIds).Scan(&roles)
	return
}
