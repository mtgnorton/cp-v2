package service

import (
	"context"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

var Post = post{}

type post struct {
}

// ChooseListByKeyword 根据关键字获取对应的主题列表
func (p *post) ChooseListByKeyword(ctx context.Context, keyword string, page ...int) (res *model.PostListOutput, err error) {
	var pa = 1
	if len(page) > 0 {
		pa = page[0]
	}
	switch keyword {
	case "all":
		res, err = p.AllPostList(ctx, model.PageSizeInput{
			Page: 1,
			Size: 40,
		})
	case "today-hot":
		res, err = p.HotList(ctx, model.Today, model.PageSizeInput{
			Page: 1,
			Size: 40,
		})
	case "day7-hot":
		res, err = p.HotList(ctx, model.Day7, model.PageSizeInput{
			Page: 1,
			Size: 40,
		})
	case "follow-user":
		res, err = p.FollowUserList(ctx, 1, model.PageSizeInput{
			Page: 1,
			Size: 40,
		})
	case "follow-node":
		res, err = p.FollowNodeList(ctx, 1, model.PageSizeInput{
			Page: 1,
			Size: 40,
		})
	default:
		res, err = p.IndexNodeList(ctx, keyword, pa)
	}
	return
}

// IndexNodeList 获取首页的主题列表
func (p *post) IndexNodeList(ctx context.Context, NodeKeyword string, page ...int) (res *model.PostListOutput, err error) {
	var pa = 1
	if len(page) == 1 {
		pa = page[0]
	}
	return shared.Post.List(ctx, &model.PostListInput{
		Period:      model.DayAll,
		NodeKeyword: NodeKeyword,
		Status:      gconv.String(model.POST_STATUS_NORMAL),
		PageSizeInput: model.PageSizeInput{
			Page: pa,
			Size: 6,
		},
		OrderFunc: func(d *gdb.Model) *gdb.Model {
			d = d.Order(gdb.Raw("if(top_end_time>now(),1,0) desc"))
			return d
		},
	})
}

// FollowNodeList 关注节点的主题列表
func (p *post) FollowNodeList(ctx context.Context, userId uint, pageSizeInput model.PageSizeInput,
) (res *model.PostListOutput, err error) {

	nodeIds, err := shared.User.GetCollectNodeIds(ctx, userId)
	if err != nil {
		return
	}
	return shared.Post.List(ctx, &model.PostListInput{
		PageSizeInput: pageSizeInput,
		NodeIds:       nodeIds,
		Status:        gconv.String(model.POST_STATUS_NORMAL),

		OrderFunc: func(d *gdb.Model) *gdb.Model {
			return d.OrderDesc(dao.Posts.Columns().CreatedAt)
		},
	})
}

// FollowUserList 关注用户的主题列表
func (p *post) FollowUserList(ctx context.Context, userId uint, pageSizeInput model.PageSizeInput,
) (res *model.PostListOutput, err error) {

	userIds, err := shared.User.GetFollowUserIds(ctx, userId)
	if err != nil {
		return
	}
	return shared.Post.List(ctx, &model.PostListInput{
		PageSizeInput: pageSizeInput,
		UserIds:       userIds,
		Status:        gconv.String(model.POST_STATUS_NORMAL),

		OrderFunc: func(d *gdb.Model) *gdb.Model {
			return d.OrderDesc(dao.Posts.Columns().CreatedAt)
		},
	})
}

// HotList 获取某个周期的热门主题列表
func (p *post) HotList(ctx context.Context, period model.TimePeriod, pageSizeInput model.PageSizeInput) (res *model.PostListOutput, err error) {

	return shared.Post.List(ctx, &model.PostListInput{
		Period:        period,
		Status:        gconv.String(model.POST_STATUS_NORMAL),
		PageSizeInput: pageSizeInput,
		OrderFunc: func(d *gdb.Model) *gdb.Model {
			d = d.Order(gdb.Raw("if(top_end_time>now(),1,0) desc")).
				OrderDesc(dao.Posts.Columns().Weight).
				OrderDesc(dao.Posts.Columns().ReplyAmount)
			return d
		},
	})
}

// AllPostList 获取所有主题列表
func (p *post) AllPostList(ctx context.Context, pageSizeInput model.PageSizeInput,
) (res *model.PostListOutput, err error) {

	return shared.Post.List(ctx, &model.PostListInput{
		PageSizeInput: pageSizeInput,
		Status:        gconv.String(model.POST_STATUS_NORMAL),

		OrderFunc: func(d *gdb.Model) *gdb.Model {
			d = d.Order(gdb.Raw("if(top_end_time>now(),1,0) desc")).
				OrderDesc(dao.Posts.Columns().CreatedAt)
			return d
		},
	})
}

// IndexNodeList 获取节点的主题列表
func (p *post) NodeList(ctx context.Context, keyword string, pageSizeInput model.PageSizeInput) (res *model.PostListOutput, err error) {
	return shared.Post.List(ctx, &model.PostListInput{
		NodeKeyword:   keyword,
		Status:        gconv.String(model.POST_STATUS_NORMAL),
		PageSizeInput: pageSizeInput,
	})
}

// HistoryList 获取用户的历史主题列表
func (p *post) HistoryList(ctx context.Context, userId uint) (res *model.PostListOutput, err error) {
	postIds, err := shared.History.GetPostIds(ctx, userId)
	return shared.Post.List(ctx, &model.PostListInput{
		PostIds: postIds,
		Status:  gconv.String(model.POST_STATUS_NORMAL),
	})
}

// Store 创建主题
func (p *post) Store(ctx context.Context, req *define.PostsStoreReq) (res *define.PostsStoreRes, err error) {

	res = &define.PostsStoreRes{}

	// 节点是否存在
	nodeCount, err := dao.Nodes.Ctx(ctx).
		Where(dao.Nodes.Columns().Id, req.NodeId).
		Where(dao.Nodes.Columns().IsVirtual, 0).
		Count()
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	if nodeCount == 0 {
		return res, response.NewError("节点不存在")
	}
	// 标题字符长度小于255
	if gstr.LenRune(req.Title) > 255 {
		return res, response.NewError("标题过长")
	}

	gjson, err := gjson.LoadContent(req.Content)
	if err != nil {
		return
	}
	var quillDelta *define.PostQuill
	err = gjson.Scan(&quillDelta)

	if err != nil {
		return res, response.NewError("内容格式错误")
	}

	quillContent := ""
	// 拼接quillDelta里所有的insert
	for _, v := range quillDelta.Ops {
		quillContent += v.Insert
	}
	g.Dump("quillContent", quillContent)

	// 内容长度小于后台设定的长度
	contentLength := gstr.LenRune(quillContent)
	settingPostsCharacterMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CHARACTER_MAX)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if gconv.Int(settingPostsCharacterMax) > 0 && contentLength > gconv.Int(settingPostsCharacterMax) {
		return res, response.NewError("内容过长")
	}

	// 用户是否被禁止发帖
	user, err := FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if ok, _ := shared.User.CanPost(ctx, user.Id); !ok {
		return res, response.NewError("您已被禁止发帖")
	}

	// 用户今天发帖数量没有超过后台设定的数量
	settingPostsDayMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_EVERY_DAY_MAX)

	hasPostsAmount, err := dao.Posts.Ctx(ctx).
		Where(dao.Posts.Columns().UserId, user.Id).
		WhereBetween(dao.Posts.Columns().CreatedAt, gconv.String(gtime.Now().Format("Y-m-d 00:00:00")), gconv.String(gtime.Now().Format("Y-m-d 23:59:59"))).Count()
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if settingPostsDayMax.Int() > 0 && hasPostsAmount > settingPostsDayMax.Int() {
		return res, response.NewError("今日创建主题数量已达上限")
	}

	// 验证主题内容是否涉及敏感词
	sensitiveWords, _ := utility.SensitiveInspector.MatchAndReplace(quillContent)
	if len(sensitiveWords) > 0 {
		return res, response.NewError("内容包含下列敏感词:" + gstr.Join(sensitiveWords, ","))
	}

	// 获取配置，发帖是否需要审核
	settingPostNeedAudit, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_IS_NEED_AUDIT)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	status := model.POST_STATUS_NORMAL
	if settingPostNeedAudit.Int() == 1 {
		status = model.POST_STATUS_NO_AUDIT
	}

	entity := &entity.Posts{}

	err = gconv.Scan(req, entity)

	entity.UserId = user.Id
	entity.Username = user.Username
	entity.CharacterAmount = uint(contentLength)
	entity.ContentType = string(model.POST_CONTENT_TYPE_QUILL) // todo 后期增加其它格式时再进行调整

	entity.LastChangeTime = gtime.Now()
	entity.Status = status
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}

	id, err := dao.Posts.Ctx(ctx).InsertAndGetId(entity)
	if err != nil {
		return res, err
	}

	res.Id = gconv.Uint(id)
	if settingPostNeedAudit.Int() == 0 {
		err = shared.Post.OfficialPublishHook(ctx, id)
		if err != nil {
			return
		}
	} else {
		return res, response.NewSuccess("正在审核中")
	}
	return res, response.NewSuccess("创建主题成功")

}

// UploadImage 上传图片
func (p *post) UploadImage(ctx context.Context, userId uint, req *define.PostsUploadImageReq) (res *define.PostsUploadImageRes, err error) {
	res = &define.PostsUploadImageRes{}
	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.Image,
		UploadPosition: model.UPLOAD_POSITION_FRONTED,
		Dir:            "post/" + gconv.String(userId),
		UploadType:     model.UPLOAD_TYPE_IMAGE,
	})
	res.Url = out.RelativePath
	return
}

/**
更新主题
更新主题的前提条件
	- 主题创建之后的 CONFIG_POSTS_CAN_UPDATE_TIME 分钟内，如果主题获得的回复数量小于 CONFIG_POSTS_CAN_UPDATE_REPLY_AMOUNT
	- 每次编辑消耗后台设定CONFIG_TOKEN_UPDATE_POSTS_DEDUCT
*/
func (p *post) Update(ctx context.Context, req *define.PostsUpdateReq) (res *define.PostsUpdateRes, err error) {
	res = &define.PostsUpdateRes{}

	posts := entity.Posts{}

	err = dao.Posts.Ctx(ctx).WherePri(req.Id).Scan(&posts)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}
	if posts.Id == 0 {
		return res, response.NewError("主题不存在")
	}
	settingCanUpdateTime, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CAN_UPDATE_TIME)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}

	if gtime.Now().Sub(posts.CreatedAt) > time.Duration(settingCanUpdateTime.Int())*time.Minute {
		return res, response.NewError(fmt.Sprintf("主题创建之后的 %d 分钟内才能编辑", settingCanUpdateTime.Int()))
	}
	settingCanUpdateReplyAmount, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_POSTS_CAN_UPDATE_REPLY_AMOUNT)
	if err != nil {
		return res, response.WrapError(err, "更新主题失败")
	}
	replyAmount, err := dao.Replies.Ctx(ctx).Where(dao.Replies.Columns().PostsId, posts.Id).Count()
	if replyAmount > settingCanUpdateReplyAmount.Int() {
		return res, response.NewError(fmt.Sprintf("主题获得的回复数量大于 %d 之后不能编辑", settingCanUpdateReplyAmount.Int()))
	}
	return
}

func (p *post) DetailWithNodeAndComments(ctx context.Context, req model.PostWithNodeAndCommentsReq) (res *model.PostWithNodeAndCommentsRes, err error) {
	res = &model.PostWithNodeAndCommentsRes{}
	err = dao.Posts.Ctx(ctx).WherePri(req.Id).With(model.PostWithNodeAndCommentsRes{}.Node, model.PostWithNodeAndCommentsRes{}.User).Scan(res)

	out, err := shared.Reply.List(ctx, &model.ReplyListInput{
		PostId:    req.Id,
		SeeUserId: req.SeeUserId,
		Status:    gconv.String(model.REPLY_STATUS_NORMAL),
		OrderFieldDirectionInput: model.OrderFieldDirectionInput{
			OrderField:     "id",
			OrderDirection: "asc",
		},
	})
	//res.Replies.Page = req.Page
	//res.Replies.Size = req.Size
	res.Replies.Total = out.Total
	res.Replies.List = out.List
	return
}

// GetSeoDesc 获取seo描述
func (p *post) GetSeoDesc(ctx context.Context, entityPost *entity.Posts) (desc string, err error) {
	desc = "作者 @" + entityPost.Username + "  "

	gjson, err := gjson.LoadContent(entityPost.Content)
	if err != nil {
		return
	}
	var quillDelta *define.PostQuill
	err = gjson.Scan(&quillDelta)

	if err != nil {
		return "", response.NewError("内容格式错误")
	}
	for _, item := range quillDelta.Ops {
		if item.Insert != "" {
			desc += gstr.Trim(item.Insert, "\n")
		}
		if gstr.LenRune(desc) > 100 {
			break
		}
	}
	desc += "..."
	return
}

// 访问主题后，进行的相关操作
func (p *post) Visit(ctx context.Context, postId uint, user *model.UserSummary) (err error) {
	d := dao.Posts.Ctx(ctx)
	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = d.WherePri(postId).Increment(dao.Posts.Columns().VisitAmount, 1)
		if err != nil {
			return
		}
		if user.Id > 0 {
			_, err = dao.UserPostsHistories.Ctx(ctx).Insert(&entity.UserPostsHistories{
				UserId:   user.Id,
				Username: user.Username,
				PostsId:  postId,
			})
		}

		return
	})
	return
}

// 用户收藏/取消主题
func (p *post) ToggleCollect(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := shared.User.WhetherCollectPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId == 0 {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Increment(dao.Posts.Columns().CollectionAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
				UserId:         userId,
				Username:       username,
				TargetId:       postId,
				TargetUserId:   post.UserId,
				TargetUsername: post.Username,
				Type:           model.AssociationTypeCollectPost,
			})
			return
		})
	} else {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Decrement(dao.Posts.Columns().CollectionAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}

	return
}

// 用户忽略/取消主题
func (p *post) ToggleShield(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := shared.User.WhetherShieldPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId == 0 {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Increment(dao.Posts.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
				UserId:         userId,
				Username:       username,
				TargetId:       postId,
				TargetUserId:   post.UserId,
				TargetUsername: post.Username,
				Type:           model.AssociationTypeShieldPost,
			})
			return
		})
	} else {
		err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(postId).Decrement(dao.Posts.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}

	return
}

// 用户感谢主题
func (p *post) ThanksPost(ctx context.Context, postId uint, userId uint, username string) (err error) {

	d := dao.Posts.Ctx(ctx)
	var post entity.Posts
	err = d.WherePri(postId).Scan(&post)
	if err != nil {
		return
	}
	if post.Id == 0 {
		return response.NewError("主题不存在")
	}

	relationId, err := shared.User.WhetherThanksPost(ctx, userId, postId)

	if err != nil {
		return
	}

	if relationId > 0 {
		return response.NewError("您已经感谢过该主题")
	}

	settingNeedToken, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_THANKS_POSTS_DEDUCT)

	if err != nil {
		return response.WrapError(err, "感谢主题失败")
	}

	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = d.WherePri(postId).Increment(dao.Posts.Columns().ThanksAmount, 1)
		if err != nil {
			return
		}
		_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
			UserId:         userId,
			Username:       username,
			TargetId:       postId,
			TargetUserId:   post.UserId,
			TargetUsername: post.Username,
			Type:           model.AssociationTypeThanksPost,
		})
		// token扣除
		err = shared.User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     userId,
			Amount:     -settingNeedToken.Int(),
			ChangeType: model.BALANCE_CHANGE_TYPE_THANK_POST_DEDUCT,
			RelationId: uint(postId),
			Remark:     fmt.Sprintf("感谢主题 %s 创建的 %s", post.Username, post.Title),
		})

		if err != nil {
			return err
		}

		// token奖励
		err = shared.User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     post.UserId,
			Amount:     settingNeedToken.Int(),
			ChangeType: model.BALANCE_CHANGE_TYPE_THANK_POST_REWARD,
			RelationId: uint(postId),
			Remark:     fmt.Sprintf("用户 %s 感谢主题 > %s", username, post.Title),
		})
		return err
	})

	return
}

// ShieldReply 屏蔽回复
func (p *post) ShieldReply(ctx context.Context, replyId uint, userId uint, username string) (err error) {
	var reply entity.Replies
	d := dao.Replies.Ctx(ctx)
	err = d.WherePri(replyId).Scan(&reply)

	if err != nil {
		return
	}
	if reply.Id == 0 {
		return response.NewError("回复不存在")
	}

	relationId, err := shared.User.WhetherShieldReply(ctx, userId, replyId)
	if relationId > 0 {
		return response.NewError("您已经屏蔽过该回复")
	}
	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = d.WherePri(replyId).Increment(dao.Replies.Columns().ShieldedAmount, 1)
		if err != nil {
			return
		}
		_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
			UserId:         userId,
			Username:       username,
			TargetId:       replyId,
			AdditionalId:   reply.PostsId,
			TargetUserId:   reply.UserId,
			TargetUsername: reply.Username,
			Type:           model.AssociationTypeShieldReply,
		})

		return
	})
	return
}

func (p *post) ThanksReply(ctx context.Context, replyId uint, userId uint, username string) (err error) {

	var reply entity.Replies
	d := dao.Replies.Ctx(ctx)
	err = d.WherePri(replyId).Scan(&reply)

	if err != nil {
		return
	}
	if reply.Id == 0 {
		return response.NewError("回复不存在")
	}

	relationId, err := shared.User.WhetherThanksReply(ctx, userId, replyId)
	if relationId > 0 {
		return response.NewError("您已经感谢过该回复")
	}

	settingNeedToken, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_THANKS_REPLY_DEDUCT)

	if err != nil {
		return response.WrapError(err, "感谢回复失败")
	}

	err = d.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = d.WherePri(replyId).Increment(dao.Replies.Columns().ThanksAmount, 1)
		if err != nil {
			return
		}
		_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
			UserId:         userId,
			Username:       username,
			TargetId:       replyId,
			AdditionalId:   reply.PostsId,
			TargetUserId:   reply.UserId,
			TargetUsername: reply.Username,
			Type:           model.AssociationTypeThanksReply,
		})

		if err != nil {
			return
		}
		err = shared.User.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     userId,
			Amount:     -settingNeedToken.Int(),
			ChangeType: "",
			RelationId: 0,
			Remark:     "",
		})
		return
	})
	return
}
