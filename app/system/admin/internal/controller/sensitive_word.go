package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"
)

var SensitiveWord = sensitiveWord{}

type sensitiveWord struct {
}

// List 敏感词列表
func (s *sensitiveWord) List(ctx context.Context, in *define.SensitiveWordListReq) (res *define.SensitiveWordListRes, err error) {
	res, err = service.SensitiveWord.List(ctx, in)
	return
}

// Store 保存敏感词
func (s *sensitiveWord) Store(ctx context.Context, in *define.SensitiveWordStoreReq) (res *define.SensitiveWordStoreRes, err error) {
	res, err = service.SensitiveWord.Store(ctx, in)
	return
}

// Destroy 删除敏感词
func (s *sensitiveWord) Destroy(ctx context.Context, in *define.SensitiveWordDestroyReq) (res *define.SensitiveWordDestroyRes, err error) {
	res, err = service.SensitiveWord.Destroy(ctx, in)
	return
}
