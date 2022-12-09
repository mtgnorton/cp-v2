package service

import (
	"context"
	"database/sql"
	"fmt"
	"gf-admin/app/dao"
	"gf-admin/app/model"
	"gf-admin/app/model/entity"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/utility"
	"gf-admin/utility/response"
	"time"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

var UserNoNeedLogin = userNoNeedLogin{}

type userNoNeedLogin struct {
}

func (a *userNoNeedLogin) Register(ctx context.Context, in *define.RegisterInput) (err error) {
	if !shared.Captcha.Verify(ctx, in.Code, in.CaptchaId) {
		return response.NewError("验证码错误")
	}
	in.IsNeedEmail = true
	in.Status = model.USER_STATUS_NO_ACTIVE
	_, err = shared.User.Register(ctx, in.UserRegisterInput)

	if err != nil {
		return err
	}

	return response.NewSuccess("注册邮件已经发送，请您激活")
}

func (a *userNoNeedLogin) Login(ctx context.Context, in *define.LoginInput) (out *define.LoginOutput, err error) {
	out = &define.LoginOutput{}
	if !shared.Captcha.Verify(ctx, in.Code, in.CaptchaId) {
		return out, response.NewError("验证码错误")
	}

	var user entity.Users
	err = dao.Users.Ctx(ctx).Where(g.Map{
		dao.Users.Columns().Username: in.Username,
		dao.Users.Columns().Password: utility.EncryptPassword(in.Username, in.Password),
	}).Scan(&user)

	if err != nil {
		return out, response.WrapError(err, "用户名或密码错误")
	}

	if user.Id == 0 {
		return out, response.NewError("用户名或密码错误")
	}

	canActive, err := shared.User.CanActive(ctx, user.Id)
	if err != nil {
		return out, response.WrapError(err, "登录失败")
	}
	if canActive {
		return out, response.NewError("请先激活账号")
	}
	// 用户是否禁止登录
	canLogin, err := shared.User.CanLogin(ctx, user.Id)
	if err != nil {
		return out, response.WrapError(err, "登录失败")
	}
	if !canLogin {
		return out, response.NewError("您的账号已被禁止登录")
	}
	token, err := FrontTokenInstance.LoadConfig().TokenHandler.GenerateAndSaveData(ctx, in.Username, user)
	if err != nil {
		return
	}
	out.Token = token
	shared.Context.SetUser(ctx, user)
	return out, response.NewSuccess("登录成功")

}

// ResendActiveEmail 用户重新发送进货邮件
func (a *userNoNeedLogin) ResendActiveEmail(ctx context.Context, req *define.ResendActiveEmailReq) (err error) {
	var user entity.Users

	if !shared.Captcha.Verify(ctx, req.Code, req.CaptchaId) {
		return response.NewError("验证码错误")
	}

	d := dao.Users.Ctx(ctx)

	err = d.Where(dao.Users.Columns().Email, req.Email).Scan(&user)
	if err == sql.ErrNoRows {
		return response.NewError("该邮箱尚未注册")
	}
	if err != nil {
		return err
	}

	// 检查是否已经激活
	active, err := shared.User.CanActive(ctx, user.Id)

	if !active {
		return response.NewError("该用户已经激活，请您直接登录")
	}

	//判断发送时间间隔是否超过后台设置的值 model.REGISTER_SEND_EMAIL_DIFF_HOUR
	diffHour, err := shared.Config.Get(ctx, model.CONFIG_MODULE_FORUM, model.CONFIG_REGISTER_SEND_EMAIL_DIFF_HOUR)

	if err != nil {
		return err
	}

	lastSendTime, err := shared.EmailRecord.GetUserLastEmailTime(ctx, req.Email, model.EMAIL_TYPE_REGISTER)
	if diffHour.Int() > 0 && time.Now().Unix()-lastSendTime < int64(diffHour.Int())*3600 {

		return response.NewError(fmt.Sprintf("发送邮件间隔不能超过%d小时，请稍后再试", diffHour.Int()), g.Map{
			"diffHour":     diffHour.Int(),
			"lastSendTime": lastSendTime,
			"now":          time.Now().Unix(),
		})
	}
	go func() {
		err = shared.User.SendActiveEmail(ctx, gconv.Int64(user.Id))
		g.Log().Error(ctx, err)
	}()
	return response.NewSuccess("注册邮件已经发送，请您激活")

}

// ForgetPassword 用户忘记密码，发送重置邮件
func (a *userNoNeedLogin) ForgetPassword(ctx context.Context, req *define.ForgetPasswordReq) (err error) {
	if !shared.Captcha.Verify(ctx, req.Code, req.CaptchaId) {
		return response.NewError("验证码错误")
	}
	// 判断用户和邮箱是否存在
	var user entity.Users
	d := dao.Users.Ctx(ctx)
	err = d.Where(g.Map{
		dao.Users.Columns().Username: req.Username,
		dao.Users.Columns().Email:    req.Email,
	}).Scan(&user)
	if err == sql.ErrNoRows {
		return response.NewError("用户不存在")
	}
	if err != nil {
		return err
	}
	// 发送重置邮件

	go func() {
		err = shared.User.SendForgetPasswordEmail(ctx, req.Username, req.Email)
		g.Log().Error(ctx, err)
	}()
	return
}

// ResetPassword 重置密码
func (a *userNoNeedLogin) ResetPassword(ctx context.Context, req *define.ResetPasswordReq) (err error) {

	// 如果req.Time超过8小时，则激活失败
	if time.Now().Unix()-req.Time > 28800 {
		return response.NewError("超过8小时，链接失效")
	}
	gvar, err := gcache.Get(ctx, "user:reset-password:"+req.Proof)

	if err != nil {
		return response.NewError("激活失败")
	}
	if gvar.Int() > 0 {
		return response.NewError("密码已经重置，请勿重复操作")
	}
	// 生成激活凭证
	proof, err := shared.User.GenerateActiveProof(ctx, req.Username, req.Email, gconv.String(req.Time))

	if err != nil {
		return response.WrapError(err, "激活失败")
	}
	if proof != req.Proof {
		return response.NewError("激活凭证错误")
	}

	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, req.Username).Update(g.Map{
		dao.Users.Columns().Password: utility.EncryptPassword(req.Username, req.Password),
	})
	if err != nil {
		return response.WrapError(err, "重置密码失败")
	}
	err = gcache.Set(ctx, "user:reset-password:"+req.Proof, 1, 8*time.Hour)

	if err != nil {
		return
	}

	return response.NewSuccess("重置密码成功")

}
