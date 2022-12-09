package shared

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/response"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

var Reply = reply{}

type reply struct {
}

// List 获取回复列表
func (r *reply) List(ctx context.Context, in *model.ReplyListInput) (out *model.ReplyListOutput, err error) {
	out = &model.ReplyListOutput{}
	d := dao.Replies.Ctx(ctx)
	if in.PostId != 0 {
		d = d.Where(dao.Replies.Columns().PostsId, in.PostId)
	}
	if in.UserId != 0 {
		d = d.Where(dao.Replies.Columns().UserId, in.UserId)
	}
	if in.Username != "" {
		d = d.WhereLike(dao.Replies.Columns().Username, in.Username)
	}
	var shieldReplyIds []gdb.Value
	if in.SeeUserId != 0 {
		t := dao.Association
		shieldReplyIds, err = t.Ctx(ctx).
			Where(t.Columns().UserId, in.SeeUserId).
			Where(t.Columns().Type, model.AssociationTypeShieldReply).
			Array(t.Columns().TargetId)
		if err != nil {
			return
		}
	}

	if len(shieldReplyIds) > 0 {
		d = d.WhereNotIn(dao.Replies.Columns().Id, shieldReplyIds)
	}

	if in.Status != "" {
		d = d.Where(dao.Replies.Columns().Status, in.Status)
	}

	if in.Keyword != "" {
		d = d.WhereLike(dao.Replies.Columns().Content, "%"+in.Keyword+"%")
	}

	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return
	}
	if in.WithPost {
		// 关联主题，获取主题信息
		d = d.LeftJoin(dao.Posts.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Replies.Table(), dao.Replies.Columns().PostsId, dao.Posts.Table(), dao.Posts.Columns().Id))

		// 关联节点，获取节点信息
		d = d.LeftJoin(dao.Nodes.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Posts.Table(), dao.Posts.Columns().NodeId, dao.Nodes.Table(), dao.Nodes.Columns().Id))

		d = d.FieldsPrefix(dao.Replies.Table(), "*").Fields("forum_posts.title as post_title,forum_posts.username as post_username,forum_posts.id as post_id ,forum_nodes.name as node_name,forum_nodes.keyword as node_keyword")
	}
	// 关联用户，获取头像信息
	d = d.LeftJoin(dao.Users.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Replies.Table(), dao.Replies.Columns().UserId, dao.Users.Table(), dao.Users.Columns().Id))

	d = d.FieldsPrefix(dao.Replies.Table(), "*").Fields("forum_users.avatar as user_avatar")

	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField + " " + in.OrderDirection)
	} else {
		d = d.OrderDesc(dao.Replies.Columns().Id)
	}

	if in.Page > 0 && in.Size > 0 {
		d = d.Page(in.Page, in.Size)
	}
	err = d.Scan(&out.List)
	return
}

// OfficialPublishHook 当回复正式发布时的钩子，如扣除发帖token，修改主题最后回复相关信息
func (r *reply) OfficialPublishHook(ctx context.Context, replyId int64) (err error) {

	var reply entity.Replies

	err = dao.Replies.Ctx(ctx).WherePri(replyId).Scan(&reply)

	if err != nil {
		return
	}
	var post entity.Posts

	err = dao.Posts.Ctx(ctx).WherePri(reply.PostsId).Scan(&post)
	if err != nil {
		return
	}
	// 获取需要扣除的余额
	settingNeedToken, err := Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_ESTABLISH_REPLY_DEDUCT)

	if err != nil {
		return response.WrapError(err, "创建主题失败")
	}

	// 扣除余额,插入回复并修改主题最后回复相关信息
	err = dao.Replies.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {

		// token扣除
		err = User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     reply.UserId,
			Amount:     -gconv.Int(settingNeedToken),
			ChangeType: model.BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_DEDUCT,
			RelationId: reply.PostsId,
			Remark:     fmt.Sprintf(`发表了长度为%d字符的回复 > <a href="%s">%s</a>`, reply.CharacterAmount, "/post/"+gconv.String(post.Id), post.Title),
		})
		if err != nil {
			return err
		}
		if reply.UserId != post.UserId {
			// 奖励创建主题者
			err = User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
				UserId:     post.UserId,
				Amount:     gconv.Int(settingNeedToken),
				ChangeType: model.BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_REWARD,
				RelationId: reply.PostsId,
				Remark:     fmt.Sprintf(`%s 回复了主题<a href="%s">%s</a>`, reply.Username, "/post/"+gconv.String(post.Id), post.Title),
			})
			if err != nil {
				return err
			}
			// 对创建主题者进行消息提醒
			_, err = dao.Messages.Ctx(ctx).Insert(entity.Messages{
				UserId:          reply.UserId,
				Username:        reply.Username,
				RepliedUserId:   post.UserId,
				RepliedUsername: post.Username,
				PostId:          post.Id,
				ReplyId:         reply.Id,
				IsRead:          0,
				Type:            gconv.String(model.MESSAGE_TYPE_POST_OWNER),
			})
			if err != nil {
				return
			}
		}

		// 对@到的用户进行提醒
		if reply.RelationUserIds != "" {
			ids := gstr.Split(reply.RelationUserIds, ",")
			for _, id := range ids {
				usernameVar, err := dao.Users.Ctx(ctx).WherePri(id).Value(dao.Users.Columns().Username)
				_, err = dao.Messages.Ctx(ctx).Insert(entity.Messages{
					UserId:          reply.UserId,
					Username:        reply.Username,
					RepliedUserId:   gconv.Uint(id),
					RepliedUsername: usernameVar.String(),
					PostId:          post.Id,
					ReplyId:         reply.Id,
					Type:            gconv.String(model.MESSAGE_TYPE_REPLY),
				})
				if err != nil {
					return err
				}
			}
		}

		// 修改主题最后回复相关信息
		_, err = dao.Posts.Ctx(ctx).WherePri(reply.PostsId).Update(g.Map{
			dao.Posts.Columns().LastChangeTime:    gtime.Now().Format("Y-m-d H:i:s"),
			dao.Posts.Columns().ReplyLastUserId:   reply.UserId,
			dao.Posts.Columns().ReplyLastUsername: reply.Username,
			dao.Posts.Columns().ReplyAmount:       gdb.Raw(dao.Posts.Columns().ReplyAmount + " + 1"),
		})

		if err != nil {
			return
		}
		// 用户回复次数+1
		_, err = dao.Users.Ctx(ctx).WherePri(reply.UserId).Update(g.Map{
			dao.Users.Columns().ReplyAmount: gdb.Raw(dao.Users.Columns().ReplyAmount + " + 1"),
		})
		return err

	})
	return
}

// Destroy 删除回复
func (r *reply) Destroy(ctx context.Context, id uint) (err error) {
	// 判断回复是否存在
	d := dao.Replies.Ctx(ctx)
	c, err := d.WherePri(id).Count()
	if c == 0 {
		return response.NewError("回复不存在")
	}
	_, err = d.WherePri(id).Delete()
	// todo 是否需要进行逆操作

	return
}
