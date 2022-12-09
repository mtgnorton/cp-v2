package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/boot"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var Menu = menu{
	rootName: "根菜单",
	rootId:   0,
}

type menu struct {
	rootName string
	rootId   uint
}

const (
	LINK_EXTERNAL       = "external" // 外链
	LINK_INTERNAL       = "internal" //内部跳转
	MENU_TYPE_DIR       = "directory"
	MENU_TYPE_LINK      = "link"
	MENU_TYPE_OPERATION = "operation"

	COMPONENT_LAYOUT      = "layout"
	COMPONENT_PARENT_VIEW = "parent_view"

	BREAD_CRUMB_REDIRECT_YES = "redirect_yes"
	BREAD_CRUMB_REDIRECT_NO  = "redirect_no"
)

func (m *menu) GetRootId(ctx context.Context) uint {
	return m.rootId
}
func (m *menu) List(ctx context.Context, in *define.MenuListInput) (output *define.MenuListOutput, err error) {
	output = &define.MenuListOutput{}

	d := dao.AdminMenu.Ctx(ctx)
	if in.Name != "" {
		d = d.WhereLike(dao.AdminMenu.Columns.Name, fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Status != "" {
		d = d.Where(dao.AdminMenu.Columns.Status, in.Status)
	}
	err = d.Scan(&output.List)

	return
}

func (m *menu) Store(ctx context.Context, in *define.MenuStoreInput) (err error) {

	d := dao.AdminMenu.Ctx(ctx)
	exist, err := d.Where(dao.AdminMenu.Columns.Name, in.Name).
		Where(dao.AdminMenu.Columns.ParentId, in.ParentId).
		Count()

	if err != nil {
		return err
	}
	if exist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "菜单名称已存在")
	}
	if in.Identification != "" {
		exist, err = d.Where(dao.AdminMenu.Columns.Identification, in.Identification).Count()
		if err != nil {
			return err
		}
		if exist > 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "权限标识已存在")
		}
	}
	_, err = d.Insert(in)
	return err
}

func (m *menu) Info(ctx context.Context, Id uint) (output *define.MenuInfoOutput, err error) {
	output = &define.MenuInfoOutput{}
	err = dao.AdminMenu.Ctx(ctx).WherePri(Id).Scan(&output)

	return
}
func (m *menu) Update(ctx context.Context, in *define.MenuUpdateInput) (err error) {

	d := dao.AdminMenu.Ctx(ctx)

	old := &entity.AdminMenu{}
	err = d.WherePri(in.Id).Scan(&old)

	if err != nil {
		return
	}

	if old == nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "菜单不存在")
	}

	exist, err := d.WhereNot(dao.AdminMenu.Columns.Id, in.Id).
		Where("? = ? && ? = ?", dao.AdminMenu.Columns.Name, in.Name, dao.AdminMenu.Columns.ParentId, in.ParentId).
		Count()

	if err != nil {
		return err
	}
	if exist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "菜单名称已存在")
	}

	if in.Identification != "" {
		exist, err = d.WhereNot(dao.AdminMenu.Columns.Id, in.Id).Where(dao.AdminMenu.Columns.Identification, in.Identification).Count()
		if err != nil {
			return err
		}
		if exist > 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "权限标识已存在")
		}
	}

	if in.Identification != "" && in.Method != "" && (in.Identification != old.Identification || in.Method != old.Method) {
		_, err = Enforcer.UpdatePolicy([]string{old.Identification, old.Method}, []string{in.Identification, in.Method})
		if err != nil {
			return
		}
	}
	_, err = d.Where(dao.AdminMenu.Columns.Id, in.Id).Update(in)
	return err
}

func (m *menu) Destroy(ctx context.Context, Id uint) (err error) {
	if err = m.ExistById(ctx, Id); err != nil {
		return err
	}
	_, err = dao.AdminMenu.Ctx(ctx).WherePri(Id).Delete()
	return
}

/*获取指定菜单的树形结构，参数用户id（adminIdArgs）是可选的，当指定时，获取该用户具有权限的菜单的树形结构*/
func (m *menu) Tree(ctx context.Context, parentId uint, adminIdArgs ...uint) (tree []*model.MenuTree, err error) {
	var (
		menus   []*entity.AdminMenu
		adminId uint
	)
	tree = make([]*model.MenuTree, 0)

	// 如果是本地开发环境，则直接获取所有的路由和菜单
	if boot.EnvName == "local" {
		adminIdArgs = []uint{}
	}

	if len(adminIdArgs) > 0 {
		adminId = adminIdArgs[0]
		menus, err = m.GetByAdministratorId(ctx, adminId)
	} else {
		err = dao.AdminMenu.Ctx(ctx).Where(dao.AdminMenu.Columns.Status, model.STATUS_NORMAL).OrderAsc(dao.AdminMenu.Columns.Sort).Scan(&menus)
	}
	if err != nil {
		return tree, err
	}
	err = m.tree(ctx, menus, &tree, parentId)
	return tree, err
}

// 获取只包含id，name,children的菜单树
func (m *menu) MiniTree(ctx context.Context) (miniTree []*model.MenuMiniTree, err error) {
	tree, err := m.Tree(ctx, m.GetRootId(ctx))
	if err != nil {
		return miniTree, err
	}
	return m.miniTree(ctx, tree)
}

func (m *menu) miniTree(ctx context.Context, tree []*model.MenuTree) (miniTree []*model.MenuMiniTree, err error) {
	miniTree = []*model.MenuMiniTree{}
	for _, menu := range tree {
		t := model.MenuMiniTree{
			Id:   menu.Id,
			Name: menu.Name,
		}
		t.Children, err = m.miniTree(ctx, menu.Children)
		miniTree = append(miniTree, &t)
	}
	return miniTree, err
}

/*tree 通过指针传递，直接直接修改原变量*/
func (m *menu) tree(ctx context.Context, menus []*entity.AdminMenu, tree *[]*model.MenuTree, parentId uint) (err error) {
	for _, menu := range menus {
		if parentId == menu.ParentId {
			t := &model.MenuTree{}
			err = gconv.Scan(menu, t)
			if err != nil {
				return
			}
			t.Children = make([]*model.MenuTree, 0)

			err = m.tree(ctx, menus, &t.Children, menu.Id)
			*tree = append(*tree, t)
		}
	}
	return
}

func (m *menu) ExistById(ctx context.Context, Id uint) (err error) {

	exist, err := dao.AdminMenu.Ctx(ctx).WherePri(Id).Count()
	if err != nil {
		return err
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "菜单不存在")
	}
	return nil
}

func (m *menu) GetByAdministratorId(ctx context.Context, administratorId uint) (menus []*entity.AdminMenu, err error) {
	menus = []*entity.AdminMenu{}
	roles, err := Role.GetByAdministratorId(ctx, administratorId)
	if err != nil {
		return menus, err
	}
	menuIds, err := dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns.RoleId, gdb.ListItemValuesUnique(roles, "Id")).Array(dao.RoleMenu.Columns.MenuId)
	if err != nil {
		return
	}
	err = dao.AdminMenu.Ctx(ctx).WherePri(menuIds).Where(dao.AdminMenu.Columns.Status, model.STATUS_NORMAL).OrderAsc(dao.AdminMenu.Columns.Sort).Scan(&menus)
	return
}

// 将生成菜单树封装成前端路由数组
func (m *menu) FrontRoutes(ctx context.Context, adminId uint) (routes []model.FrontRoute, err error) {
	tree, err := m.Tree(ctx, m.rootId, adminId)
	if err != nil {
		return
	}
	routes, err = m.frontRoutes(ctx, tree)
	return
}

func (m *menu) frontRoutes(ctx context.Context, tree []*model.MenuTree) (routes []model.FrontRoute, err error) {
	routes = []model.FrontRoute{}
	for _, menu := range tree {

		if menu.Type == MENU_TYPE_OPERATION {
			continue
		}
		fr := model.FrontRoute{
			Component: m.getComponentPath(menu),
			Hidden:    false,
			Meta: model.Meta{
				Title: menu.Name,
				Icon:  menu.Icon,
			},
			Name:     m.getRouteName(menu),
			Path:     m.getRoutePath(menu),
			Children: []model.FrontRoute{},
		}
		if len(menu.Children) > 0 && menu.Type == MENU_TYPE_DIR {
			if menu.Type == MENU_TYPE_DIR {
				fr.Redirect = BREAD_CRUMB_REDIRECT_NO
				fr.AlwaysShow = true
			}
			childrenRoutes, errSub := m.frontRoutes(ctx, menu.Children)
			if errSub != nil {
				err = errSub
				return
			}
			fr.Children = childrenRoutes
		} else if m.isTopInternalLink(menu) {
			fr.Redirect = menu.Path
			child := model.FrontRoute{
				Path:      menu.Path,
				Component: menu.FrontComponentPath,
				Name:      menu.Name,
				Meta: model.Meta{
					Title: menu.Name,
					Icon:  menu.Icon,
				},
			}
			fr.Children = append(fr.Children, child)
		}
		routes = append(routes, fr)
	}
	return
}

// 获取组件信息
func (m *menu) getComponentPath(menu *model.MenuTree) string {
	component := COMPONENT_LAYOUT
	if menu.FrontComponentPath != "" && !m.isTopInternalLink(menu) {
		component = menu.FrontComponentPath
	} else if menu.FrontComponentPath == "" && m.isSubNestDir(menu) {
		component = COMPONENT_PARENT_VIEW
	} else if menu.LinkType == LINK_EXTERNAL {
		component = ""
	}
	return component
}

//是否是顶级内部链接
func (m *menu) isTopInternalLink(menu *model.MenuTree) bool {
	return menu.ParentId == m.rootId && menu.Type == MENU_TYPE_LINK && menu.LinkType == LINK_INTERNAL
}

//是否是子级嵌套菜单
func (m *menu) isSubNestDir(menu *model.MenuTree) bool {
	return menu.ParentId != m.rootId && menu.Type == MENU_TYPE_DIR
}

// 获取路由名称
func (m *menu) getRouteName(menu *model.MenuTree) (routeName string) {
	routeName = gstr.UcFirst(menu.Path)
	//if m.isTopInternalLink(menu) {
	//	return ""
	//}
	return
}

// 获取路由地址
func (m *menu) getRoutePath(menu *model.MenuTree) string {
	routerPath := menu.Path
	// 非外链并且是一级目录
	if m.rootId == menu.ParentId && menu.Type == MENU_TYPE_DIR && menu.LinkType == LINK_INTERNAL {
		routerPath = "/" + menu.Path
	} else if m.isTopInternalLink(menu) {
		routerPath = "/"
	}
	return routerPath
}
