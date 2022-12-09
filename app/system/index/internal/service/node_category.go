package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/index/internal/define"

	"github.com/gogf/gf/v2/database/gdb"
)

var NodeCategory = nodeCategory{}

type nodeCategory struct {
}

// List 获取某些节点分类下的所有节点
func (n *nodeCategory) List(ctx context.Context, in *define.NodeCategoryListInput) (out *define.NodeCategoryListOutput, err error) {
	out = &define.NodeCategoryListOutput{}
	out.NodeTotal, err = dao.Nodes.Ctx(ctx).Count()
	if err != nil {
		return
	}
	d := dao.NodeCategories.Ctx(ctx)
	columns := dao.NodeCategories.Columns()
	if in.IsIndexNavigation != "" {
		d = d.Where(columns.IsIndexNavigation, in.IsIndexNavigation)
	}
	if in.Ids != nil {
		d = d.Where(columns.Id, in.Ids)
	}
	err = d.ScanList(&out.List, "NodeCategories")
	if err != nil {
		return
	}

	err = dao.Nodes.Ctx(ctx).
		Where(dao.Nodes.Columns().CategoryId, gdb.ListItemValuesUnique(out.List, "Id")).
		ScanList(&out.List, "Nodes", "NodeCategories", "category_id:Id")

	return
}
