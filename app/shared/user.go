package shared

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/utility"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/crypto/gaes"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
)

var User = user{}

type user struct {
}

// Info 获取用户信息
func (u *user) Info(ctx context.Context, in *model.UserInfoInput) (user *entity.Users, err error) {
	user = &entity.Users{}
	d := dao.Users.Ctx(ctx)
	if in.Username != "" {
		d = d.Where(dao.Users.Columns().Username, in.Username)
	}
	if in.Id != 0 {
		d = d.Where(dao.Users.Columns().Id, in.Id)
	}
	err = d.Scan(&user)

	if err == sql.ErrNoRows {
		return user, nil
	}

	return
}

// CanLogin 是否能够登录
func (u *user) CanLogin(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if status.Int()&model.USER_STATUS_DISABLE_LOGIN != 0 {
		return false, nil
	}

	return true, nil
}

// CanPost 是否能够发帖
func (u *user) CanPost(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if status.Int()&model.USER_STATUS_DISABLE_POST != 0 {
		return false, nil
	}
	return true, nil
}

// CanReply 是否能够回复
func (u *user) CanReply(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if status.Int()&model.USER_STATUS_DISABLE_REPLY != 0 {
		return false, nil
	}
	return true, nil
}

// CanActive 用户是否能够激活,返回true表示用户未激活
func (u *user) CanActive(ctx context.Context, userId uint) (bool, error) {
	status, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Status)
	if err != nil {
		return false, err
	}
	if status.Int()&model.USER_STATUS_NO_ACTIVE != 0 {
		return true, nil
	}
	return false, nil
}

// List 获取用户列表
func (u *user) List(ctx context.Context, in *model.UserListInput) (out *model.UserListOutput, err error) {
	out = &model.UserListOutput{}

	d := dao.Users.Ctx(ctx)
	if in.Username != "" {
		d = d.WhereLike(dao.Users.Columns().Username, fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Status != "" {
		// 正常
		if in.Status == "0" {
			d = d.Where(dao.Users.Columns().Status, 0)
		} else {
			d = d.Where(`status  & ? > 0`, in.Status)
		}
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()
	if err != nil {
		return out, response.NewError(err.Error())
	}
	d = d.Page(in.Page, in.Size)
	if in.OrderField != "" && in.OrderDirection != "" {
		d = d.Order(in.OrderField, in.OrderDirection)
	}

	err = d.Scan(&out.List)

	return
}

// Register 注册
func (u *user) Register(ctx context.Context, in *model.UserRegisterInput) (userId int64, err error) {
	userId = 0
	d := dao.Users.Ctx(ctx)
	//判断用户名是否存在
	exist, err := d.Where(dao.Users.Columns().Username, in.Username).Count()
	if err != nil {
		return userId, response.NewError(err.Error())
	}
	if exist > 0 {
		return userId, response.NewError("用户名已存在")
	}
	//判断邮箱是否存在
	exist, err = d.Where(dao.Users.Columns().Email, in.Email).Count()
	if err != nil {
		return userId, response.NewError(err.Error())
	}
	if exist > 0 {
		return userId, response.NewError("邮箱已存在")
	}
	var user *entity.Users

	if err = gconv.Struct(in, &user); err != nil {
		return userId, err
	}
	user.Password = utility.EncryptPassword(user.Username, user.Password)

	// 设置默认头像
	user.Avatar, err = Config.GetString(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REGISTER_DEFAULT_AVATAR)

	err = dao.Users.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		userId, err = dao.Users.Ctx(ctx).OmitEmptyData().InsertAndGetId(user)
		if err != nil {
			return err
		}
		settingRegisterGiveAmount, err := Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_TOKEN_REGISTER_GIVE)
		if err != nil {
			return err
		}
		return u.ChangeBalance(ctx, &model.UserChangeBalanceInput{
			UserId:     uint(userId),
			Amount:     settingRegisterGiveAmount.Int(),
			ChangeType: model.BALANCE_CHANGE_TYPE_REGSITER,
		})
	})

	if in.IsNeedEmail {
		go func() {
			err = u.SendActiveEmail(ctx, userId)
			g.Log().Error(ctx, err)
		}()
	}

	return userId, err
}

// SendActiveEmail 发送激活邮件
func (u *user) SendActiveEmail(ctx context.Context, userId int64) (err error) {
	user, err := u.Info(ctx, &model.UserInfoInput{
		Id: gconv.Uint(userId),
	})
	if err != nil {
		return
	}
	if user.Id == 0 {
		return response.NewError("用户不存在")
	}

	can, err := u.CanActive(ctx, gconv.Uint(userId))

	if !can {
		return response.NewError("用户已激活")
	}

	siteName, logo, err := Common.GetSiteNameAndLogo(ctx)

	if err != nil {
		return
	}
	now := gtime.Timestamp()

	proof, err := u.GenerateActiveProof(ctx, user.Username, gconv.String(now))

	if err != nil {
		return
	}
	url, err := Common.CompleteUrl(ctx, fmt.Sprintf("/register/verify/%s/%s/%s", user.Username, gconv.String(now), proof))

	// 获取邮件模板 模板路径 app/system/index/internal/template/email.html
	c, err := g.View().Parse(ctx, "email/register.html", g.Map{
		"siteName": siteName,
		"username": user.Username,
		"url":      url,
		"logo":     logo,
	})

	if err != nil {
		return
	}

	err = Email.SendSync(ctx, &model.SendEmailInput{
		Subject: siteName,
		Body:    c,
		To:      user.Email,
		IsHtml:  true,
		Type:    model.EMAIL_TYPE_REGISTER,
	})
	return
}

// SendChangeEmail 发送邮箱变更邮件
func (u *user) SendChangeEmail(ctx context.Context, username string, newEmail string) (err error) {

	user, err := u.Info(ctx, &model.UserInfoInput{
		Username: username,
	})

	if err != nil {
		return
	}

	if user.Id == 0 {
		return response.NewError("用户不存在")
	}

	siteName, logo, err := Common.GetSiteNameAndLogo(ctx)
	if err != nil {
		return
	}
	now := gtime.Timestamp()

	proof, err := u.GenerateActiveProof(ctx, user.Username, newEmail, gconv.String(now))

	if err != nil {
		return
	}
	url, err := Common.CompleteUrl(ctx, fmt.Sprintf("/user/update/email-active/%s/%s/%s/%s", user.Username, newEmail, gconv.String(now), proof))

	// 获取邮件模板 模板路径 app/system/index/internal/template/email.html
	c, err := g.View().Parse(ctx, "email/change-email.html", g.Map{
		"siteName": siteName,
		"username": user.Username,
		"url":      url,
		"logo":     logo,
	})

	if err != nil {
		return
	}

	err = Email.SendSync(ctx, &model.SendEmailInput{
		Subject: siteName,
		Body:    c,
		To:      newEmail,
		IsHtml:  true,
		Type:    model.EMAIL_TYPE_CHANGE_EMAIL,
	})
	return
}

// SendForgetPasswordEmail 忘记密码发送重置密码邮件
func (u *user) SendForgetPasswordEmail(ctx context.Context, username string, email string) (err error) {

	user, err := u.Info(ctx, &model.UserInfoInput{
		Username: username,
	})

	if err != nil {
		return
	}
	if user.Id == 0 {
		return response.NewError("用户不存在")
	}
	siteName, logo, err := Common.GetSiteNameAndLogo(ctx)
	if err != nil {
		return
	}
	now := gtime.Timestamp()

	proof, err := u.GenerateActiveProof(ctx, user.Username, email, gconv.String(now))

	if err != nil {
		return
	}
	url, err := Common.CompleteUrl(ctx, fmt.Sprintf("/user/reset-password/%s/%s/%s/%s", user.Username, email, gconv.String(now), proof))

	// 获取邮件模板 模板路径 app/system/index/internal/template/email.html
	c, err := g.View().Parse(ctx, "email/forget-password.html", g.Map{
		"siteName": siteName,
		"username": user.Username,
		"url":      url,
		"logo":     logo,
	})

	if err != nil {
		return
	}

	err = Email.SendSync(ctx, &model.SendEmailInput{
		Subject: siteName,
		Body:    c,
		To:      email,
		IsHtml:  true,
		Type:    model.EMAIL_TYPE_CHANGE_EMAIL,
	})
	return
}

// GenerateActiveProof 生成激活凭证
func (u *user) GenerateActiveProof(ctx context.Context, fields ...string) (proof string, err error) {
	var encryptKey string

	encryptKeyVar := g.Cfg().MustGet(ctx, "front.registerEncryptKey")

	if encryptKeyVar.String() == "" {
		encryptKey = "noworldcanexpressmywholehearted1"
	} else {
		encryptKey = encryptKeyVar.String()
	}

	bytes := gconv.Bytes("")
	for _, field := range fields {
		bytes = append(bytes, gconv.Bytes(field)...)
	}
	aesValue, err := gaes.Encrypt([]byte(bytes), []byte(encryptKey))
	if err != nil {

		err = response.WrapError(err, "", g.Map{
			"encryptKey":    encryptKey,
			"encryptKeyVar": encryptKeyVar.String(),
			"fields":        fields,
			"aesValue":      aesValue,
		})
		return
	}
	proof = hex.EncodeToString(aesValue)
	return
}

// Balance 查询用户余额
func (u *user) Balance(ctx context.Context, userId uint) (uint, error) {
	balanceVar, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Value(dao.Users.Columns().Balance)

	return balanceVar.Uint(), err
}

// ChangeBalance 修改余额
func (u *user) ChangeBalance(ctx context.Context, in *model.UserChangeBalanceInput) error {

	if in.Amount == 0 {
		return nil
	}
	return dao.Users.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		balanceVar, err := dao.Users.Ctx(ctx).WherePri(in.UserId).LockUpdate().Value(dao.Users.Columns().Balance)
		balance := balanceVar.Uint()

		if in.Amount < 0 && balance < uint(-in.Amount) {
			return response.NewError("积分不足")
		}

		_, err = dao.Users.Ctx(ctx).WherePri(in.UserId).Increment(dao.Users.Columns().Balance, in.Amount)
		if err != nil {
			return err
		}
		username, err := dao.Users.Ctx(ctx).WherePri(in.UserId).Value(dao.Users.Columns().Username)
		if err != nil {
			return err
		}
		var log = entity.BalanceChangeLog{
			UserId:     in.UserId,
			Username:   username.String(),
			Type:       string(in.ChangeType),
			Amount:     in.Amount,
			Before:     uint(balance),
			After:      uint(int(balance) + in.Amount),
			RelationId: in.RelationId,
			Remark:     in.Remark,
		}
		_, err = dao.BalanceChangeLog.Ctx(ctx).Insert(log)
		return err
	})
}

// 获取用户关注用户的id数组
func (u *user) GetFollowUserIds(ctx context.Context, userId uint) ([]uint, error) {
	return u.GetRelateToContentIds(ctx, userId, 0, model.AssociationTypeFollowUser)
}

// GetCollectNodeIds 获取用户收藏节点的id数组
func (u *user) GetCollectNodeIds(ctx context.Context, userId uint) (targetIds []uint, err error) {
	return u.GetRelateToContentIds(ctx, userId, 0, model.AssociationTypeCollectNode)
}

// GetThanksReplyIds 获取用户感谢回复的id数组
func (u *user) GetThanksReplyIds(ctx context.Context, userId, postId uint) (targetIds []uint, err error) {
	return u.GetRelateToContentIds(ctx, userId, postId, model.AssociationTypeThanksReply)
}

// GetCollectPostIds  获取用户收藏主题的id数组
func (u *user) GetCollectPostIds(ctx context.Context, userId uint) (targetIds []uint, err error) {
	return u.GetRelateToContentIds(ctx, userId, 0, model.AssociationTypeCollectPost)
}

//  GetRelateToContentIds 获取用户 感谢｜屏蔽|收藏|关注 --> 主题｜回复 |节点 |用户 的id数组
func (u *user) GetRelateToContentIds(ctx context.Context, userId, additionalId uint, relationType string) (targetIds []uint, err error) {
	if userId == 0 {
		return
	}
	condition := g.Map{
		dao.Association.Columns().UserId: userId,
		dao.Association.Columns().Type:   relationType,
	}
	if additionalId > 0 {
		condition[dao.Association.Columns().AdditionalId] = additionalId
	}

	v, err := dao.Association.Ctx(ctx).Where(condition).Array(dao.Association.Columns().TargetId)

	// 将[]gdb.Value 转为 []uint
	targetIds = make([]uint, 0)
	for _, item := range v {
		targetIds = append(targetIds, item.Uint())
	}
	return targetIds, err
}

// WhetherShieldUser 是否屏蔽了用户，如果id为0则表示未屏蔽
func (u *user) WhetherShieldUser(ctx context.Context, userId, shieldUserId uint) (uint, error) {
	return u.WhetherRelateToContent(ctx, userId, shieldUserId, model.AssociationTypeShieldUser)
}

// WhetherFollowUser 是否关注用户,如果id为0则表示未关注
func (u *user) WhetherFollowUser(ctx context.Context, userId, followUserId uint) (uint, error) {
	return u.WhetherRelateToContent(ctx, userId, followUserId, model.AssociationTypeFollowUser)
}

// WhetherCollectPost 用户是否收藏主题,如果id为0则表示未收藏
func (u *user) WhetherCollectPost(ctx context.Context, userId, postId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, postId, model.AssociationTypeCollectPost)
}

// WhetherShieldPost 用户是否屏蔽主题，如果id为0则表示未屏蔽
func (u *user) WhetherShieldPost(ctx context.Context, userId, postId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, postId, model.AssociationTypeShieldPost)
}

// WhetherThanksPost 用户是否感谢主题，如果id为0则表示未感谢
func (u *user) WhetherThanksPost(ctx context.Context, userId, postId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, postId, model.AssociationTypeThanksPost)
}

// WhetherShieldReply 用户是否屏蔽回复，如果id为0则表示未屏蔽
func (u *user) WhetherShieldReply(ctx context.Context, userId, postId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, postId, model.AssociationTypeShieldReply)
}

// WhetherThanksReply 用户是否感谢回复，如果id为0则表示未感谢
func (u *user) WhetherThanksReply(ctx context.Context, userId, postId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, postId, model.AssociationTypeThanksReply)
}

// WhetherCollectNode 用户是否收藏节点，如果id为0则表示未收藏
func (u *user) WhetherCollectNode(ctx context.Context, userId, nodeId uint) (id uint, err error) {
	// 判断是否已经收藏
	return u.WhetherRelateToContent(ctx, userId, nodeId, model.AssociationTypeCollectNode)
}

// WhetherRelateToContent 判断用户是否 感谢｜屏蔽|收藏|关注 --> 主题｜回复 |节点｜用户
func (u *user) WhetherRelateToContent(ctx context.Context, userId, targetId uint, relationType string) (id uint, err error) {

	v, err := dao.Association.Ctx(ctx).Where(g.Map{
		dao.Association.Columns().UserId:   userId,
		dao.Association.Columns().TargetId: targetId,
		dao.Association.Columns().Type:     relationType,
	}).Value(dao.Association.Columns().Id)

	return v.Uint(), err

}

// GetHistoryPostList  获取用户最近浏览的主题
func (u *user) GetHistoryPostList(ctx context.Context, userId uint) (list *model.PostListOutput, err error) {
	list = &model.PostListOutput{}
	if userId == 0 {
		return
	}
	historyPostIds, err := History.GetPostIds(ctx, userId, 5)

	return Post.List(ctx, &model.PostListInput{
		Status:  gconv.String(model.POST_STATUS_NORMAL),
		PostIds: historyPostIds,
	})

}

// UpdatePassword 更新用户密码，oldPassword可选，当传递时，则进行验证
func (u *user) UpdatePassword(ctx context.Context, username, newPassword string, oldPasswords ...string) (err error) {
	var user entity.Users
	d := dao.Users.Ctx(ctx)
	err = d.Where(g.Map{
		dao.Users.Columns().Username: username,
	}).Scan(&user)
	if err != nil {
		return
	}
	var oldPassword = ""
	if len(oldPasswords) > 0 {
		oldPassword = oldPasswords[0]
	}
	if oldPassword != "" && utility.EncryptPassword(username, oldPassword) != user.Password {
		return response.NewError("旧密码错误")
	}
	_, err = d.Where(dao.Users.Columns().Username, username).Update(g.Map{
		dao.Users.Columns().Password: utility.EncryptPassword(username, newPassword),
	})
	if err != nil {
		return
	}
	return response.NewSuccess("密码修改成功")

}
