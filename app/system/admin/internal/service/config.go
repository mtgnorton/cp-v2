package service

import (
	"context"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/define"
	"github.com/gogf/gf/v2/container/gvar"
)

var Config = config{}

type config struct {
}

// 获取多个模块的配置
func (c *config) GetModules(ctx context.Context, in *define.ConfigListInput) (out *define.ConfigListOutput, err error) {
	out = &define.ConfigListOutput{}
	out.Data = make(map[string]map[string]*gvar.Var)

	for _, module := range in.Modules {
		m, err := shared.Config.Gets(ctx, module)
		out.Data[module] = m

		if err != nil {
			return nil, err
		}
	}
	return
}

// 更新某个模块的配置
func (c *config) Update(ctx context.Context, in *define.ConfigUpdateInput) (err error) {

	return shared.Config.Sets(ctx, in.Module, in.KeyValueMap)
}
