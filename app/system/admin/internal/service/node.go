package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Node = node{}

type node struct {
}

func (n *node) Store(ctx context.Context, in *define.NodeStoreInput) (err error) {
	d := dao.Nodes.Ctx(ctx)
	exist, err := d.Where(dao.Nodes.Columns().Name, in.Name).Count()
	if err != nil {
		return err
	}
	if exist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点名称已存在")
	}
	exist, err = d.Where(dao.Nodes.Columns().Keyword, in.Keyword).Count()
	if err != nil {
		return err
	}
	if exist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点关键字已存在")
	}

	_, err = d.Insert(in)
	return err
}

func (n *node) Update(ctx context.Context, in *define.NodeUpdateInput) (err error) {
	d := dao.Nodes.Ctx(ctx)
	var oldNode *entity.Nodes
	err = d.WherePri(in.Id).Scan(&oldNode)
	if err != nil {
		return
	}
	if oldNode.Id == 0 {
		return response.NewError("节点不存在")
	}
	nameExist, err := d.WhereNot(dao.Nodes.Columns().Id, in.Id).Where(dao.Nodes.Columns().Name, in.Name).Count()
	if err != nil {
		return err
	}
	if nameExist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点名称已存在")
	}
	keywordExist, err := d.WhereNot(dao.Nodes.Columns().Id, in.Id).Where(dao.Nodes.Columns().Keyword, in.Keyword).Count()
	if err != nil {
		return err
	}
	if keywordExist > 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "节点关键词已存在")
	}
	_, err = d.WherePri(in.Id).Update(in)

	// 递归的将子节点的分类也更新为当前节点的分类
	if oldNode.CategoryId != gconv.Int(in.CategoryId) {
		var waitNodes []*entity.Nodes
		err = d.Where(dao.Nodes.Columns().ParentId, in.Id).Scan(&waitNodes)
		if err != nil {
			return
		}

		for i := len(waitNodes); i > 0; i-- {

			_, err = dao.Nodes.Ctx(ctx).WherePri(waitNodes[i-1].Id).Update(g.Map{
				dao.Nodes.Columns().CategoryId: in.CategoryId,
			})
			if err != nil {
				return
			}
			// 获取该节点的子节点
			var childNodes []*entity.Nodes
			err = d.Where(dao.Nodes.Columns().ParentId, waitNodes[i-1].Id).Scan(&childNodes)

			if err != nil {
				return
			}
			waitNodes = append(waitNodes, childNodes...)
			// 删除该节点
			waitNodes = append(waitNodes[:i-1], waitNodes[i:]...)

		}

	}
	return err
}

func (n *node) Destroy(ctx context.Context, Id uint) (err error) {
	err = shared.Node.ExistById(ctx, Id)
	if err != nil {
		return err
	}
	d := dao.Nodes.Ctx(ctx)
	_, err = d.WherePri(Id).Delete()
	return err
}
