package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"/user-list" method:"get" summary:"用户列表" tags:"用户管理"`
	*model.UserListInput
}

type UserListRes struct {
	*model.UserListOutput
}

type UserStoreReq struct {
	g.Meta `path:"/user-store" method:"post" summary:"添加用户" tags:"用户管理"`
	model.UserRegisterInput
}

type UserStoreRes struct {
}

type UserUpdateReq struct {
	g.Meta      `path:"/user-update" method:"put" summary:"更新用户" tags:"用户管理"`
	Id          uint   `json:"id"                    `   // 用户ID
	Username    string `json:"username"               `  // 用户名
	Email       string `json:"email"                  `  // email
	Description string `json:"description"            `  // 简介
	Password    string `json:"password"               `  // MD5密码
	Password2   string `json:"password2"               ` // MD5密码
	Avatar      string `json:"avatar"                 `  // 头像地址
	Status      string `json:"status"                 `  // 状态：disable_login | disable_posts | disable_reply
	Remark      string `json:"remark"                 `  // 备注
}

type UserUpdateRes struct {
}

type UserInfoReq struct {
	g.Meta `path:"/user-info" method:"get" summary:"用户详情" tags:"用户管理"`
	Id     uint `json:"id" v:"min:1#请选择需要查看的用户"`
}
type UserInfoRes struct {
	entity.Users
}

type UserDestroyReq struct {
	g.Meta `path:"/user-destroy" method:"delete" summary:"删除用户" tags:"用户管理"`
	Id     uint `json:"id"                    ` // 用户ID
}

type UserDestroyRes struct {
}
