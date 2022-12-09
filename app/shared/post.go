package shared

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility/response"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var Post = post{}

type post struct {
}

//List 获取主题列表的通用方法 默认排序字段为LastChangeTime
func (p *post) List(ctx context.Context, in *model.PostListInput) (out *model.PostListOutput, err error) {
	out = &model.PostListOutput{
		List: make([]*model.PostWithoutContent, 0),
	}
	d := dao.Posts.Ctx(ctx)

	if in.Title != "" {
		d = d.WhereLike(dao.Posts.Columns().Title, fmt.Sprintf("%%%s%%", in.Title))
	}

	if in.Period == -1 {
		// 代表今日0点到现在
		d = d.WhereGTE(dao.Posts.Columns().CreatedAt, gtime.NewFromTime(time.Now()).Format("Y-m-d 00:00:00"))

	} else if in.Period > 0 {
		d = d.WhereGTE(dao.Posts.Columns().CreatedAt, gtime.Now().Add(-time.Duration(in.Period)*time.Hour*24))
	}

	if in.NodeKeyword != "" {
		d = d.Where("`forum_nodes`.`keyword`='" + in.NodeKeyword + "'")
	}
	if in.NodeName != "" {
		d = d.Where("`forum_nodes`.`name`='" + in.NodeName + "'")
	}
	if in.NodeIds != nil {
		d = d.Where(dao.Posts.Columns().NodeId, in.NodeIds)
	}
	if in.UserIds != nil {
		d = d.Where(dao.Posts.Columns().UserId, in.UserIds)
	}
	if in.Usernames != nil {
		d = d.Where(dao.Posts.Columns().Username, in.Usernames)
	}
	if in.Status != "" {
		d = d.Where(dao.Posts.Columns().Status, in.Status)
	}
	if in.PostIds != nil {
		d = d.Where(dao.Posts.Columns().Id, in.PostIds)
	}

	if in.FilterKeyword != "" {
		d = d.Where(d.Builder().
			WhereLike("forum_posts.title", "%"+in.FilterKeyword+"%").
			WhereOrLike("forum_posts.content", "%"+in.FilterKeyword+"%"))
	}

	// nodes.id=post.node_id
	d = d.LeftJoin(dao.Nodes.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Posts.Table(), dao.Posts.Columns().NodeId, dao.Nodes.Table(), dao.Nodes.Columns().Id))
	// users.id=post.user_id
	d = d.LeftJoin(dao.Users.Table(), fmt.Sprintf("%s.%s=%s.%s", dao.Posts.Table(), dao.Posts.Columns().UserId, dao.Users.Table(), dao.Users.Columns().Id))

	if in.Page > 0 && in.Size > 0 {
		out.Page = in.Page
		out.Size = in.Size
		out.Total, err = d.Count()
	}

	if in.OrderFunc != nil {
		d = in.OrderFunc(d)
	}

	d = d.OrderDesc(dao.Posts.Columns().LastChangeTime)

	if in.Page > 0 && in.Size > 0 {
		d = d.Page(in.Page, in.Size)
	}

	err = d.FieldsPrefix(dao.Posts.Table(), "*").Fields("forum_nodes.name as node_name,forum_nodes.keyword as node_keyword,forum_users.avatar as user_avatar,if(top_end_time>now(),1,0) is_top").Scan(&out.List)

	return

}

// OfficialPublishHook 主题正式发布后执行的钩子
func (p *post) OfficialPublishHook(ctx context.Context, postId int64) (err error) {

	var post entity.Posts
	d := dao.Posts.Ctx(ctx)
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	settingNeedToken, err := Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_ESTABLISH_POSTS_DEDUCT)

	if err != nil {
		return response.WrapError(err, "")
	}

	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 增加用户发帖数量
		_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, post.UserId).Update(g.Map{
			dao.Users.Columns().PostsAmount: gdb.Raw(dao.Users.Columns().PostsAmount + " + 1"),
		})
		if err != nil {
			return err
		}
		// token扣除
		err = User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     post.UserId,
			Amount:     -settingNeedToken.Int(),
			ChangeType: model.BALANCE_CHANGE_TYPE_ESTABLISH_POST_DEDUCT,
			RelationId: uint(postId),
		})
		return err

	})
	return
}

// Destroy 删除回复
func (r *post) Destroy(ctx context.Context, id uint) (err error) {
	// 判断回复是否存在
	d := dao.Posts.Ctx(ctx)
	c, err := d.WherePri(id).Count()
	if c == 0 {
		return response.NewError("主题不存在")
	}
	_, err = d.WherePri(id).Delete()
	// todo 是否需要进行逆操作

	return
}

// Exist 判断主题是否存在
func (p *post) Exist(ctx context.Context, id uint, status ...int) (exist bool, err error) {
	d := dao.Posts.Ctx(ctx).WherePri(id)
	if len(status) > 0 {
		d = d.Where(dao.Posts.Columns().Status, status[0])
	}
	c, err := d.Count()
	if err != nil {
		return
	}
	exist = c > 0
	return
}

// Info 获取主题详情
func (p *post) Info(ctx context.Context, id uint) (post entity.Posts, err error) {
	err = dao.Posts.Ctx(ctx).WherePri(id).Scan(&post)
	return
}
