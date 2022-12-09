package define

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UserSettingPageReq struct {
	g.Meta `path:"/user/setting/*type" method:"get" tags:"个人中心" summary:"个人设置页面"`
	Type   string
}
type UserSettingPageRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UserUpdateInfoReq struct {
	g.Meta       `path:"/user/update" method:"post" tags:"个人中心" summary:"更新个人信息"`
	Site         string `v:"max-length:50#个人网站长度不能超过50字符"`  // 个人站点
	Company      string `v:"max-length:50#所在公司长度不能超过50字符"`  // 所在公司
	Job          string `v:"max-length:50#工作职位长度不能超过50字符"`  // 工作职位
	Location     string `v:"max-length:50#所在地长度不能超过50字符"`   // 所在地
	Signature    string `v:"max-length:255#签名长度不能超过50字符"`   // 个人签名
	Introduction string `v:"max-length:255#个人简介长度不能超过50字符"` // 个人简介

}

type UserUpdateInfoRes struct {
}

type UserUpdatePasswordReq struct {
	g.Meta      `path:"/user/update/password" method:"post" tags:"个人中心" summary:"更新密码"`
	OldPassword string `v:"required#旧密码不能为空"`
	Password    string `v:"required|password#请输入密码|密码长度需要在长度在6~18之间" dc:"密码" d:"password" json:"password"`
	Password2   string `v:"required|same:password#请输入确认密码|两次密码不一致" dc:"确认密码" d:"password2" json:"password2"`
}

type UserUpdatePasswordRes struct {
}

type UserUpdateEmailReq struct {
	g.Meta   `path:"/user/update/email" method:"post" tags:"个人中心" summary:"更新邮箱"`
	NewEmail string `v:"required|email#请输入邮箱|邮箱格式不正确"`
	Password string `v:"required#请输入密码"`
}

type UserUpdateEmailRes struct {
}

type UserUploadAvatarReq struct {
	g.Meta `path:"/user/upload-avatar" method:"post" tags:"个人中心" summary:"上传头像"`
	Avatar *ghttp.UploadFile `v:"required#请选择头像"`
}

type UserUploadAvatarRes struct {
}
