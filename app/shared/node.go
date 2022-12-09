package shared

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Node = node{}

type node struct {
}

// List 获取节点列表
func (n *node) List(ctx context.Context, in *model.NodeListInput) (out *model.NodeListOutput, err error) {
	out = &model.NodeListOutput{}
	d := dao.Nodes.Ctx(ctx)
	var nodes []*entity.Nodes
	if in.Name != "" {
		d = d.WhereLike(dao.Nodes.Columns().Name, fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.IsIndex != "" {
		d = d.Where(dao.Nodes.Columns().IsIndex, in.IsIndex)
	}
	if in.IsVirtual != "" {
		d = d.Where(dao.Nodes.Columns().IsVirtual, in.IsVirtual)
	}

	d = d.OrderAsc(dao.Nodes.Columns().Sort).
		OrderAsc(dao.Nodes.Columns().Id)

	err = d.Scan(&nodes)

	out.List = make([]*model.NodeTree, 0)

	if in.NeedChildren {

		var allNodes []*entity.Nodes

		err = dao.Nodes.Ctx(ctx).Scan(&allNodes)
		if err != nil {
			return
		}
		out = &model.NodeListOutput{}

		for _, node := range nodes {
			t := &model.NodeTree{}
			err = gconv.Scan(node, t)
			if err != nil {
				return
			}

			err = n.tree(ctx, allNodes, t)
			if err != nil {
				return
			}
			out.List = append(out.List, t)
		}

	} else {
		err = gconv.Scan(&nodes, &out.List)
	}

	err = dao.NodeCategories.Ctx(ctx).WherePri(gdb.ListItemValuesUnique(out.List, "CategoryId")).ScanList(&out.List, "NodeCategory", "Nodes", "id:CategoryId")

	return
}

// tree 获取节点的树形结构
func (n *node) tree(ctx context.Context, allNodes []*entity.Nodes, parentNode *model.NodeTree) (err error) {
	for _, node := range allNodes {
		if node.ParentId == parentNode.Id {
			t := &model.NodeTree{}
			err = gconv.Scan(node, t)
			if err != nil {
				return
			}
			t.Children = make([]*model.NodeTree, 0)
			err = n.tree(ctx, allNodes, t)

			if parentNode.Children == nil {
				parentNode.Children = make([]*model.NodeTree, 0)
			}

			parentNode.Children = append(parentNode.Children, t)
		}
	}
	return
}

// Detail 获取某个节点详情
func (n *node) Detail(ctx context.Context, req model.NodeDetailInput) (node entity.Nodes, err error) {
	node = entity.Nodes{}
	d := dao.Nodes.Ctx(ctx)
	if req.Id > 0 {
		d = d.WherePri(req.Id)
		err = d.Scan(&node)
	}
	if req.Keyword != "" {
		d = d.Where(dao.Nodes.Columns().Keyword, req.Keyword)
		err = d.Scan(&node)
	}

	return
}

// GetCollectAmount 获取节点收藏数量
func (n *node) GetCollectAmount(ctx context.Context, nodeId uint) (amount int, err error) {

	t := dao.Association
	amount, err = t.Ctx(ctx).
		Where(t.Columns().TargetId, nodeId).
		Where(t.Columns().Type, model.AssociationTypeCollectNode).
		Count()
	return
}

// ExistById 判断节点是否存在
func (n *node) ExistById(ctx context.Context, Id uint) (err error) {
	d := dao.Nodes.Ctx(ctx)
	exist, err := d.WherePri(Id).Count()
	if err != nil {
		return
	}
	if exist == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点不存在")
	}
	return
}

// ExistByKeyword 判断节点是否存在
func (n *node) ExistByKeyword(ctx context.Context, keyword string) (exist bool, err error) {
	d := dao.Nodes.Ctx(ctx)
	count, err := d.Where(dao.Nodes.Columns().Keyword, keyword).Count()

	return count > 0, err
}
