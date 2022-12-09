package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
)

var Association = association{}

type association struct {
}

// List 关联列表
func (a *association) List(ctx context.Context, in *model.AssociationListInput) (out *model.AssociationListOutput, err error) {

	d := dao.Association.Ctx(ctx)
	c := dao.Association.Columns()
	out = &model.AssociationListOutput{}
	if in.Username != "" {
		d = d.Where(c.Username, in.Username)
	}
	if in.Type != "" {
		d = d.Where(c.Type, in.Type)
	}
	if in.TargetId != 0 {
		d = d.Where(c.TargetId, in.TargetId)
	}
	if in.TargetUsername != "" {
		d = d.Where(c.TargetUsername, in.TargetUsername)
	}
	if in.AdditionalId != 0 {
		d = d.Where(c.AdditionalId, in.AdditionalId)
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return
	}
	err = d.
		Page(in.Page, out.Size).
		OrderDesc(dao.BalanceChangeLog.Columns().Id).Scan(&out.List)
	// 根据关联的类型，获取关联的显示内容即TargetContent
	for _, item := range out.List {
		switch item.Type {
		case model.AssociationTypeThanksPost, model.AssociationTypeShieldPost, model.AssociationTypeCollectPost:
			postTitle, err1 := dao.Posts.Ctx(ctx).WherePri(item.TargetId).Value(dao.Posts.Columns().Title)
			if err1 != nil {
				err = err1
				return
			}
			item.TargetContent = postTitle.String()
		case model.AssociationTypeThanksReply, model.AssociationTypeShieldReply:
			replyContent, err1 := dao.Replies.Ctx(ctx).WherePri(item.TargetId).Value(dao.Replies.Columns().Content)
			if err1 != nil {
				err = err1
				return
			}
			item.TargetContent = replyContent.String()

		case model.AssociationTypeCollectNode:
			nodeName, err1 := dao.Nodes.Ctx(ctx).WherePri(item.TargetId).Value(dao.Nodes.Columns().Name)
			if err1 != nil {
				err = err1
				return
			}
			item.TargetContent = nodeName.String()
		case model.AssociationTypeFollowUser, model.AssociationTypeShieldUser:
			item.TargetContent = item.TargetUsername
		}

	}
	return
}
