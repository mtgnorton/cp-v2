package service

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/utility/response"

	"github.com/gogf/gf/v2/frame/g"
)

var SensitiveWord = sensitiveWord{}

type sensitiveWord struct {
}

// List 获取敏感词列表
func (s *sensitiveWord) List(ctx context.Context, in *define.SensitiveWordListReq) (res *define.SensitiveWordListRes, err error) {
	res = &define.SensitiveWordListRes{}
	d := dao.SensitiveWords.Ctx(ctx)
	d = d.Where(dao.SensitiveWords.Columns().Word+" LIKE ?", "%"+in.Keyword+"%")
	d = d.Page(in.Page, in.Size).
		OrderDesc(dao.SensitiveWords.Columns().Id)

	res.Page = in.Page
	res.Size = in.Size
	res.Total, err = d.Count()

	if err != nil {
		return
	}
	err = d.Scan(&res.List)
	return
}

// Store 保存敏感词
func (s *sensitiveWord) Store(ctx context.Context, in *define.SensitiveWordStoreReq) (res *define.SensitiveWordStoreRes, err error) {
	res = &define.SensitiveWordStoreRes{}
	d := dao.SensitiveWords.Ctx(ctx)

	count, err := d.Where(dao.SensitiveWords.Columns().Word, in.Word).Count()
	if count > 0 {
		return nil, response.NewError("敏感词已存在")
	}
	_, err = d.Insert(g.Map{
		dao.SensitiveWords.Columns().Word: in.Word,
		dao.SensitiveWords.Columns().Type: "",
	})
	return
}

// Destroy 删除敏感词
func (s *sensitiveWord) Destroy(ctx context.Context, in *define.SensitiveWordDestroyReq) (res *define.SensitiveWordDestroyRes, err error) {
	res = &define.SensitiveWordDestroyRes{}
	d := dao.SensitiveWords.Ctx(ctx)
	_, err = d.WherePri(in.Ids).Delete()
	return
}
