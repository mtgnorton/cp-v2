package service

import (
	"context"
	"database/sql"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

var User = user{}

type user struct {
}

// GetUserInfo 获取用户显示在右侧侧边栏的信息
func (u *user) GetTemplateShow(ctx context.Context) (user *model.UserSummary, err error) {
	user, err = FrontTokenInstance.GetUser(ctx)

	followUserIds, err := shared.User.GetFollowUserIds(ctx, user.Id)
	if err != nil {
		return
	}
	user.FollowUserCount = gconv.Uint(len(followUserIds))

	collectNodeIds, err := shared.User.GetCollectNodeIds(ctx, user.Id)
	if err != nil {
		return
	}
	user.CollectNodeCount = gconv.Uint(len(collectNodeIds))

	collectPostIds, err := shared.User.GetCollectPostIds(ctx, user.Id)
	if err != nil {
		return
	}
	user.CollectPostCount = gconv.Uint(len(collectPostIds))

	balance, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, user.Id).Value(dao.Users.Columns().Balance)
	if err != nil {
		return
	}
	user.Balance = balance.Int()

	// 获取未读的消息
	user.NoReadCount, err = shared.Message.GetUserNoReadCount(ctx, user.Id)

	return
}

// ToggleFollowUser 关注｜取消 用户 userId:当前用户id, targetId:被关注用户id
func (u *user) ToggleFollowUser(ctx context.Context, userId uint, username string, targetId uint) (err error) {
	if userId == targetId {
		return response.NewError("不能关注自己")
	}
	d := dao.Users.Ctx(ctx)
	// 判断targetId 是否存在
	var targetUser entity.Users
	err = d.WherePri(targetId).Scan(&targetUser)
	if err != nil {
		return
	}
	if targetUser.Id == 0 {
		return response.NewError("关注用户不存在")
	}

	relationId, err := shared.User.WhetherFollowUser(ctx, userId, targetId)
	if err != nil {
		return err
	}

	g.Dump(relationId, "relationId")
	if relationId == 0 {
		err = dao.Association.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(targetId).Increment(dao.Users.Columns().FollowByOtherAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
				UserId:         userId,
				Username:       username,
				TargetId:       targetId,
				TargetUserId:   targetId,
				TargetUsername: targetUser.Username,
				Type:           model.AssociationTypeFollowUser,
			})

			return
		})
	} else {
		// 取消关注
		err = dao.Association.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(targetId).Decrement(dao.Users.Columns().FollowByOtherAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}
	return
}

// ToggleShieldUser 屏蔽｜取消屏蔽 用户 userId:当前用户id, targetId:被屏蔽用户id
func (u *user) ToggleShieldUser(ctx context.Context, userId uint, username string, targetId uint) (err error) {
	if userId == targetId {
		return response.NewError("不能屏蔽自己")
	}
	d := dao.Users.Ctx(ctx)
	// 判断targetId 是否存在
	var targetUser entity.Users
	err = d.WherePri(targetId).Scan(&targetUser)
	if err != nil {
		return
	}
	if targetUser.Id == 0 {
		return response.NewError("屏蔽用户不存在")
	}

	relationId, err := shared.User.WhetherShieldUser(ctx, userId, targetId)
	if err != nil {
		return err
	}
	if relationId == 0 {
		err = dao.Association.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(targetId).Increment(dao.Users.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).Insert(&entity.Association{
				UserId:         userId,
				Username:       username,
				TargetId:       targetId,
				TargetUserId:   targetId,
				TargetUsername: targetUser.Username,
				Type:           model.AssociationTypeShieldUser,
			})

			return
		})
	} else {
		// 取消关注
		err = dao.Association.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
			_, err = d.WherePri(targetId).Decrement(dao.Users.Columns().ShieldedAmount, 1)
			if err != nil {
				return
			}
			_, err = dao.Association.Ctx(ctx).WherePri(relationId).Delete()
			return
		})
	}
	return
}

// GetCollectNodes 获取用户收藏的节点
func (u *user) GetCollectNodes(ctx context.Context, userId uint) (out []*define.CollectNodeItem, err error) {
	collectNodeIds, err := shared.User.GetCollectNodeIds(ctx, userId)
	if err != nil || len(collectNodeIds) == 0 {
		return
	}

	err = dao.Nodes.Ctx(ctx).Where(dao.Nodes.Columns().Id, collectNodeIds).Scan(&out)
	if err != nil {
		return
	}
	// 填充每项的收藏人数
	for _, item := range out {
		item.CollectAmount, err = shared.Node.GetCollectAmount(ctx, item.Id)
		if err != nil {
			return
		}
	}
	return
}

// GetCollectPosts 获取用户收藏的主题 todo 分页
func (u *user) GetCollectPosts(ctx context.Context, userId uint) (out *model.PostListOutput, err error) {
	out = &model.PostListOutput{}

	collectPostIds, err := shared.User.GetCollectPostIds(ctx, userId)
	if err != nil || len(collectPostIds) == 0 {
		return
	}
	out, err = shared.Post.List(ctx, &model.PostListInput{
		PostIds: collectPostIds,
	})
	return
}

// GetFollowUserPosts 获取用户关注的用户发表的所有主题 todo 分页
func (u *user) GetFollowUserPosts(ctx context.Context, userId uint) (out *model.PostListOutput, err error) {
	out = &model.PostListOutput{}

	followUserIds, err := shared.User.GetFollowUserIds(ctx, userId)
	if err != nil || len(followUserIds) == 0 {
		return
	}
	out, err = shared.Post.List(ctx, &model.PostListInput{
		UserIds: followUserIds,
	})
	return
}

// Active 激活用户
func (u *user) Active(ctx context.Context, req *define.ActivePageReq) (err error) {
	if req.Proof == "" || req.Time == 0 || req.Username == "" {
		return response.NewError("激活失败")
	}

	// 如果req.Time超过24小时，则激活失败
	if time.Now().Unix()-req.Time > 86400 {
		return response.NewError("超过一天，激活失败")
	}

	var user entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, req.Username).Scan(&user)
	if err == sql.ErrNoRows {
		return response.NewError("用户不存在")
	}
	if err != nil {
		return response.WrapError(err, "激活失败")
	}

	// 生成激活凭证
	proof, err := shared.User.GenerateActiveProof(ctx, req.Username, gconv.String(req.Time))

	if err != nil {
		return response.WrapError(err, "激活失败")
	}
	if proof != req.Proof {
		return response.NewError("激活凭证错误")
	}

	status := user.Status &^ model.USER_STATUS_NO_ACTIVE
	_, err = dao.Users.Ctx(ctx).WherePri(user.Id).Update(g.Map{
		dao.Users.Columns().Status: status,
	})
	return
}

// UpdateInfo 保存用户信息
func (u *user) UpdateInfo(ctx context.Context, userId uint, req *define.UserUpdateInfoReq) (err error) {
	d := dao.Users.Ctx(ctx)

	// 保存用户信息
	_, err = d.WherePri(userId).Update(req)
	if err != nil {
		return
	}
	return response.NewSuccess("用户信息修改成功")
}

// UpdateEmail 更新用户邮箱
func (u *user) UpdateEmail(ctx context.Context, username string, req *define.UserUpdateEmailReq) (err error) {
	d := dao.Users.Ctx(ctx)
	var user entity.Users

	err = d.Where(dao.Users.Columns().Username, username).Scan(&user)

	if err == sql.ErrNoRows {
		return response.NewError("用户不存在")
	}
	if err != nil {
		return response.WrapError(err, "更新邮箱失败")
	}

	// 判断密码是否正确
	if utility.EncryptPassword(username, req.Password) != user.Password {
		return response.NewError("密码错误")
	}

	// 判断邮箱是否已经被使用
	count, err := d.Where(dao.Users.Columns().Email, req.NewEmail).Count()
	if err != nil {
		return
	}
	if count > 0 {
		return response.NewError("邮箱已经被使用")
	}

	go func() {
		// 发送激活邮件
		err = shared.User.SendChangeEmail(ctx, username, req.NewEmail)
		if err != nil {
			g.Log().Error(ctx, err, g.Map{
				"username": username,
				"email":    req.NewEmail,
			})
		}
	}()

	return response.NewSuccess("激活邮件已发送，请注意查收")
}

// UpdateEmailActive 更新邮箱后，用户点击邮箱里的邮件后，将新的邮箱保存到数据库
func (u *user) UpdateEmailActive(ctx context.Context, req *define.UserUpdateEmailActivePageReq) (err error) {
	if req.Proof == "" || req.Time == 0 || req.Username == "" || req.Email == "" {
		return response.NewError("激活失败")
	}

	// 如果req.Time超过8小时，则激活失败
	if time.Now().Unix()-req.Time > 28800 {
		return response.NewError("超过8小时，激活失败,请重新修改邮箱")
	}

	var user entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, req.Username).Scan(&user)
	if err == sql.ErrNoRows {
		return response.NewError("用户不存在")
	}
	if err != nil {
		return response.WrapError(err, "激活失败")
	}

	// 生成激活凭证
	proof, err := shared.User.GenerateActiveProof(ctx, req.Username, req.Email, gconv.String(req.Time))

	if err != nil {
		return response.WrapError(err, "激活失败")
	}
	if proof != req.Proof {
		return response.NewError("激活凭证错误")
	}

	_, err = dao.Users.Ctx(ctx).WherePri(user.Id).Update(g.Map{
		dao.Users.Columns().Email: req.Email,
	})
	return
}

// UploadAvatar 上传头像
func (u *user) UploadAvatar(ctx context.Context, username string, req *define.UserUploadAvatarReq) (err error) {
	var user entity.Users
	d := dao.Users.Ctx(ctx)

	out, err := shared.Upload.Single(ctx, &model.UploadInput{
		File:           req.Avatar,
		UploadPosition: model.UPLOAD_POSITION_FRONTED,
		Dir:            "avatar/" + username,
		UploadType:     model.UPLOAD_TYPE_IMAGE,
		LimitSize:      "2mb",
	})
	if err != nil {
		return
	}
	// 保存用户信息
	_, err = d.Where(dao.Users.Columns().Username, username).
		Update(g.Map{
			dao.Users.Columns().Avatar: out.RelativePath,
		})
	if err != nil {
		return
	}

	err = d.Where(dao.Users.Columns().Username, username).Scan(&user)
	if err != nil {
		return
	}
	err = FrontTokenInstance.UpdateData(ctx, username, user)
	if err != nil {
		return
	}
	return response.NewSuccess("头像上传成功")
}

func (a *user) Info(ctx context.Context) (out *define.AuthInfoOutput, err error) {
	out = &define.AuthInfoOutput{}

	userId, err := FrontTokenInstance.GetUserId(ctx)
	if err != nil {
		return
	}

	err = dao.Users.Ctx(ctx).WherePri(userId).Scan(&out)

	if err != nil {
		return out, response.WrapError(err, "获取失败")
	}
	return out, response.NewSuccess("获取成功")
}

func (a *user) Logout(ctx context.Context) (err error) {
	token, err := FrontTokenInstance.TokenHandler.GetTokenFromRequest(ctx, g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	err = FrontTokenInstance.Remove(ctx, token)
	if err != nil {
		return err
	}
	return response.NewSuccess("退出成功")

}
