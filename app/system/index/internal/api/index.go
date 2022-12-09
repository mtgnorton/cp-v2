package api

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/util/gconv"
)

var Index = index{}

type index struct {
}

func (i *index) Index(ctx context.Context, req *define.AppletIndexReq) (res *define.AppletIndexRes, err error) {
	res = &define.AppletIndexRes{}
	if req.Keyword == "" {
		// 取出在首页的节点中，顺序第一(sort值最小)的节点
		firstNode, err := service.Node.GetIndexFirst(ctx)
		if err != nil {
			return res, response.WrapError(err, "系统错误")
		}
		req.Keyword = firstNode.Keyword
	}
	// 判断节点是否存在
	if exist, _ := shared.Node.ExistByKeyword(ctx, req.Keyword); !exist {

		return res, response.NewError("节点不存在")
	}

	res.NodeList, err = shared.Node.List(ctx, &model.NodeListInput{
		IsIndex: "1",
	})
	if err != nil {
		return
	}

	res.PostList, err = service.Post.ChooseListByKeyword(ctx, req.Keyword, gconv.Int(req.Page))
	if err != nil {
		return
	}
	res.CurrentKeyword = req.Keyword
	return
}
