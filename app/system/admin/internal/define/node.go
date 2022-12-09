package define

import (
	"gf-admin/app/model"
	"gf-admin/app/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
)

type NodeListReq struct {
	g.Meta `path:"/node-list" method:"get" summary:"节点列表" tags:"节点管理"`
	model.NodeListInput
}

type NodeListRes struct {
	*model.NodeListOutput
}

type NodeInfoReq struct {
	g.Meta `path:"/node-info" method:"get" summary:"节点详情" tags:"节点管理"`
	Id     uint `json:"id" dc:"id" v:"min:1#请选择节点id"`
}

type NodeInfoRes struct {
	entity.Nodes
}

type NodeStoreInput struct {
	Name        string `json:"name" v:"required#节点名称不能为空"`
	Keyword     string `json:"keyword" v:"required#节点关键词不能为空"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Img         string `json:"img"`
	ParentId    uint   `json:"parent_id"`
	CategoryId  uint   `json:"category_id"`
	IsIndex     int    `json:"is_index"`
	Sort        int    `json:"sort"`
}

type NodeStoreReq struct {
	g.Meta `path:"/node-store" method:"post" summary:"添加节点" tags:"节点管理"`
	*NodeStoreInput
}

type NodeStoreRes struct {
}

type NodeUpdateReq struct {
	g.Meta `path:"/node-update" method:"put" summary:"更新节点" tags:"节点管理"`
	*NodeUpdateInput
}

type NodeUpdateRes struct {
}
type NodeUpdateInput struct {
	Id          uint   `json:"id" v:"required#节点id不能为空"`
	Name        string `json:"name" v:"required#节点名称不能为空"`
	Keyword     string `json:"keyword" v:"required#节点关键字不能为空"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Img         string `json:"img"`
	ParentId    uint   `json:"parent_id"`
	CategoryId  uint   `json:"category_id"`
	IsIndex     int    `json:"is_index"`
	Sort        int    `json:"sort"`
}

type NodeDestroyReq struct {
	g.Meta `path:"/node-destroy" method:"delete" summary:"删除节点" tags:"节点管理"`
	Id     uint `json:"id" v:"required#节点id不能为空"`
}

type NodeDestroyRes struct {
}

type NodeUploadImgReq struct {
	g.Meta    `path:"/node-upload-img" method:"post" summary:"上传节点图片" tags:"节点管理"`
	NodeImage *ghttp.UploadFile `v:"required#上传图片不能为空" p:"node_img"`
}
type NodeUploadImgRes struct {
	Url string `json:"url"`
}

type NodeDownloadImgReq struct {
	g.Meta   `path:"/node-download-img" method:"post" summary:"下载节点图片" tags:"节点管理"`
	ImageUrl string `v:"required#图片地址不能为空" `
}
type NodeDownloadImgRes struct {
	Url string `json:"url"`
}
