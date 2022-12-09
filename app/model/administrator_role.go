package model

type RoleBase struct {
	Id             uint   `json:"id"`
	Name           string `v:"required#角色名必须" dc:"角色名称" json:"name"`
	Identification string `v:"required#角色标识必须" dc:"角色标志性符" json:"identification"`
	Sort           int    `v:"integer#排序字段必须是整数" dc:"越小越靠前" json:"sort"`
	Status         string `v:"in:normal,disabled#状态错误" dc:"状态，normal正常，disabled 禁用" d:"normal" json:"status"`
	MenuIds        []uint `dc:"角色具有的所有菜单权限id集合" json:"menu_ids"`
}

type RoleShow struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Identification string `json:"identification"`
	Status         string `json:"status"`
}
