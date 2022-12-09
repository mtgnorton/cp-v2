package model

type MenuBase struct {
	Name               string `json:"name"                 ` // 菜单名称
	Path               string `json:"path"                 ` // 请求前端路径，可以是外链
	Method             string `json:"method"               `
	ParentId           uint   `json:"parent_id"            ` // 父id
	Type               string `json:"type"`
	LinkType           string `json:"link_type"`
	Identification     string `json:"identification"       ` // 权限标识符
	FrontComponentPath string `json:"front_component_path" ` // 前端组件路径
	Icon               string `json:"icon"                 ` // 菜单图标
	Sort               int    `json:"sort"                 ` // 显示顺序，越小越靠前
	Status             string `json:"status"               ` // 状态 normal 正常 disabled 禁用
}

type MenuShow struct {
	Id uint `json:"id"`
	MenuBase
}

type MenuTree struct {
	MenuShow
	Children []*MenuTree `json:"children"`
}

type MenuMiniTree struct {
	Id       uint            `json:"id"                 `
	Name     string          `json:"name"                 ` // 菜单名称
	Children []*MenuMiniTree `json:"children"`
}

type Meta struct {
	Title string `json:"title"` // 标题
	Icon  string `json:"icon"`  // 图标
}

// 前台动态路由
type FrontRoute struct {
	AlwaysShow bool         `json:"alwaysShow,omitempty"` // 总是显示
	Component  string       `json:"component,omitempty"`  // 组件路径
	Hidden     bool         `json:"hidden"`               // 是否隐藏
	Meta       Meta         `json:"meta"`                 // meta
	Name       string       `json:"name"`                 // 名称
	Path       string       `json:"path"`                 // 地址
	Redirect   string       `json:"redirect,omitempty"`   // 跳转链接
	Children   []FrontRoute `json:"children,omitempty"`   // 子菜单

}
