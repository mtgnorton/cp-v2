package model

type AdministratorBase struct {
	Id       uint   `json:"id"`
	Username string `v:"required|passport#请输入管理员用户名|管理员用户名只能字母开头，只能包含字母、数字和下划线，长度在6~18之间" dc:"admin" d:"admin" json:"username"`
	Password string `v:"required#请输入密码" dc:"123456" d:"123456" json:"password"`
	Status   string `d:"normal" dc:"管理员状态，可选值：normal,disabled"  json:"status"`
	Nickname string `dc:"昵称" d:"nickname" json:"nickname"`
	Remark   string `dc:"管理员备注" d:"备注测试" json:"remark"`
	Avatar   string `dc:"管理员头像地址" d:"http://www.baidu.com" json:"avatar"`
	RoleIds  []uint `dc:"角色列表id" json:"role_ids"`
}

type AdministratorSummary struct {
	Id       uint        `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Status   string      `json:"status"`
	Nickname string      `json:"nickname"`
	Remark   string      `json:"remark"`
	Avatar   string      `json:"avatar"`
	Roles    []*RoleShow `json:"roles"`
	Menus    []*MenuShow `json:"menus"`
}
