package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
)

var Prompt = prompt{}

type prompt struct {
}

// GetContent 获取指定位置的可用提示
func (p *prompt) GetContent(ctx context.Context, position model.PromptPosition) (res string, err error) {
	d := dao.Prompts.Ctx(ctx)

	gvar, err := d.Where(dao.Prompts.Columns().Position, position).
		Where(dao.Prompts.Columns().IsDisabled, 0).
		Value(dao.Prompts.Columns().Content)

	if err != nil {
		return
	}
	return gvar.String(), nil
}
