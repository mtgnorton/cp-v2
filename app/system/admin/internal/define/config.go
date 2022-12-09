package define

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ConfigListInput struct {
	Modules []string `json:"module"`
}
type ConfigListOutput struct {
	Data map[string]map[string]*gvar.Var `json:"data"`
}
type ConfigListReq struct {
	g.Meta `path:"/config-list" method:"get" summary:"配置管理" tags:"配置管理"`
	*ConfigListInput
}

type ConfigListRes struct {
	*ConfigListOutput
}

type ConfigUpdateInput struct {
	Module      string                 `json:"module"`
	KeyValueMap map[string]interface{} `json:"key_value_map"`
}

type ConfigUpdateReq struct {
	g.Meta `path:"/config-update" method:"put" summary:"保存配置" tags:"配置管理"`
	*ConfigUpdateInput
}

type ConfigUpdateRes struct {
}

type UploadLogoReq struct {
	g.Meta `path:"/upload-logo" method:"post" summary:"上传logo图片" tags:"配置管理"`
	Logo   *ghttp.UploadFile `v:"required#上传图片不能为空" p:"logo"`
}
type UploadLogoRes struct {
	Url string `json:"url"`
}

type UploadFaviconReq struct {
	g.Meta  `path:"/upload-favicon" method:"post" summary:"上传favicon图片" tags:"配置管理"`
	Favicon *ghttp.UploadFile `v:"required#上传图片不能为空" p:"logo"`
}
type UploadFaviconRes struct {
	Url string `json:"url"`
}

type UploadDefaultAvatarReq struct {
	g.Meta `path:"/upload-default-avatar" method:"post" summary:"上传默认头像" tags:"配置管理"`
	Avatar *ghttp.UploadFile `v:"required#上传头像不能为空"`
}
type UploadDefaultAvatarRes struct {
	Url string `json:"url"`
}
