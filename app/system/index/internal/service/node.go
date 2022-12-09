package service

import (
	"context"
	"database/sql"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/utility/response"
)

var Node = node{}

type node struct {
}

// GetIndexFirst 获取首页的第一个节点
func (n *node) GetIndexFirst(ctx context.Context) (node entity.Nodes, err error) {
	err = dao.Nodes.Ctx(ctx).Where(dao.Nodes.Columns().IsIndex, 1).Order(dao.Nodes.Columns().Sort + " asc").
		OrderAsc(dao.Nodes.Columns().Id).
		Scan(&node)

	if err == sql.ErrNoRows {
		return node, nil
	}

	return
}

//ToggleCollect 用户收藏/取消节点
func (n *node) ToggleCollect(ctx context.Context, nodeId uint, userId uint, username string) (err error) {

	d := dao.Nodes.Ctx(ctx)
	var node entity.Nodes
	err = d.WherePri(nodeId).Scan(&node)
	if err != nil {
		return
	}
	if node.Id == 0 {
		return response.NewError("节点不存在")
	}

	if node.IsVirtual == 1 {
		return response.NewError("节点不存在")
	}
	relationId, err := shared.User.WhetherCollectNode(ctx, userId, nodeId)

	if err != nil {
		return
	}
	if relationId == 0 {
		_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
			UserId:   userId,
			Username: username,
			TargetId: nodeId,
			Type:     model.AssociationTypeCollectNode,
		})

	} else {
		_, err = dao.Association.Ctx(ctx).WherePri(relationId).Delete()
	}

	return
}
