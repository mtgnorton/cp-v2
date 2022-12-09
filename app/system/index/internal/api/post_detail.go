package api

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/define"
	"gf-admin/app/system/index/internal/service"
	"gf-admin/utility/response"
)

var PostDetail = postDetail{}

type postDetail struct {
}

func (p *postDetail) Detail(ctx context.Context, req *define.AppletPostsDetailReq) (res *define.AppletPostsDetailRes, err error) {
	res = &define.AppletPostsDetailRes{}
	//判断主题是否存在
	exist, err := shared.Post.Exist(ctx, req.PostId)
	if err != nil {
		return
	}
	if !exist {
		service.View().Render404(ctx, define.View{Error: "主题不存在"})
		return
	}
	// 判断主题是否未审核
	exist, err = shared.Post.Exist(ctx, req.PostId, model.POST_STATUS_NORMAL)
	if err != nil {
		return
	}
	if !exist {
		service.View().Render404(ctx, define.View{Error: "主题正在审核中"})
		return
	}
	res.PostWithComments, err = service.Post.DetailWithNodeAndComments(ctx, model.PostWithNodeAndCommentsReq{
		Id: req.PostId,
	})

	if err != nil {
		return res, response.WrapError(err, "系统错误")
	}
	return

}
