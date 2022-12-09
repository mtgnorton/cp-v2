package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/text/gregex"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/text/gstr"
)

var Reply = reply{}

type reply struct {
}

func (c *reply) Store(ctx context.Context, req *define.ReplyStoreReq) (res *define.ReplyStoreRes, err error) {
	res = &define.ReplyStoreRes{}

	// 判断回复字符长度
	contentLength := gstr.LenRune(req.Content)
	settingReplyCharacterMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_CHARACTER_MAX)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	if gconv.Int(settingReplyCharacterMax) > 0 && contentLength > gconv.Int(settingReplyCharacterMax) {
		return res, response.NewError("内容过长")
	}

	// 判断主题是否存在
	var post entity.Posts
	err = dao.Posts.Ctx(ctx).Where(dao.Posts.Columns().Id, req.PostId).Scan(&post)
	if err != nil {
		return res, response.WrapError(err, "主题不存在")
	}

	// 判断用户是否被禁止发帖
	user, err := FrontTokenInstance.GetUser(ctx)
	if err != nil {
		return res, response.WrapError(err, "创建主题失败")
	}
	if ok, _ := shared.User.CanReply(ctx, user.Id); !ok {
		return res, response.NewError("您已被禁止回复")
	}

	// 判断是否达到最大发帖数量
	settingDayReplyMax, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_EVERY_DAY_MAX)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}

	// 判断用户今日发帖的数量是否超过设定数量
	replyCount, err := dao.Replies.Ctx(ctx).
		Where(dao.Replies.Columns().UserId, user.Id).
		WhereGTE(dao.Replies.Columns().CreatedAt, gtime.Now().Format("Y-m-d")+" 00:00:00").
		Count()
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}

	if gconv.Int(settingDayReplyMax) > 0 && replyCount > gconv.Int(settingDayReplyMax) {
		return res, response.NewError("超过今日发帖数量")
	}

	sensitiveBeginTime := gtime.TimestampMilli()

	// 敏感词检查
	sensitiveWords, _ := utility.SensitiveInspector.MatchAndReplace(req.Content)

	g.Log().Infof(ctx, "敏感词检查耗时：%d", gtime.TimestampMilli()-sensitiveBeginTime)

	if len(sensitiveWords) > 0 {
		return res, response.NewError("内容包含以下敏感词:" + gstr.Join(sensitiveWords, ","))
	}
	// 解析回复内容中涉及到的所有其它用户
	matches, err := gregex.MatchAllString(`@(\w+)[^\w]?`, req.Content)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	relationUserIds := ""

	g.Dump("matches", matches, req.Content)
	for _, match := range matches {
		// 获取用户
		var user entity.Users
		err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, match[1]).Scan(&user)
		if err != nil {
			continue
		}

		relationUserIds += gconv.String(user.Id) + ","
	}
	relationUserIds = gstr.Trim(relationUserIds, ",")

	//判断是否需要审核
	settingReplyNeedAudit, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REPLY_IS_NEED_AUDIT)
	if err != nil {
		return res, response.WrapError(err, "回复失败")
	}
	status := model.REPLY_STATUS_NORMAL
	if settingReplyNeedAudit.Int() == 1 {
		status = model.REPLY_STATUS_NO_AUDIT
	}

	req.Content, err = utility.ReplaceWarp(req.Content)
	if err != nil {
		return
	}
	replyId, err := dao.Replies.Ctx(ctx).InsertAndGetId(&entity.Replies{
		PostsId:         req.PostId,
		UserId:          user.Id,
		Username:        user.Username,
		Content:         req.Content,
		CharacterAmount: gconv.Uint(contentLength),
		RelationUserIds: relationUserIds,
		Status:          status,
	})

	if err != nil {
		return
	}
	if settingReplyNeedAudit.Int() == 0 {
		err = shared.Reply.OfficialPublishHook(ctx, replyId)
		if err != nil {
			return
		}
	} else {
		return res, response.NewSuccess("正在审核中")

	}

	return res, response.NewSuccess("回复成功")
}
