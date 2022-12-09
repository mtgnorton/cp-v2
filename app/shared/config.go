package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var Config = config{
	BACKEND: "backend",
}

type config struct {
	BACKEND string
}

//根据传递的module和key获取对应的配置值
//module 可以为空
func (c *config) Get(ctx context.Context, module, key string) (value *gvar.Var, err error) {

	value, err = gcache.GetOrSetFunc(ctx, model.CONFIG_CONFIGS_PREFIX+module+key, func(ctx context.Context) (interface{}, error) {
		condition := g.Map{
			dao.Config.Columns.Key: key,
		}
		if module != "" {
			condition[dao.Config.Columns.Module] = module
		}
		record, err := dao.Config.Ctx(ctx).Where(condition).One()
		if err != nil {
			return nil, err
		}
		return record[dao.Config.Columns.Value], nil

	}, 0)

	return
}

func (c *config) GetString(ctx context.Context, module, key string) (string, error) {
	v, err := c.Get(ctx, module, key)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

//根据传递的module和keys批量获取配置值
//module 可以为空,keys可以为空
func (c *config) Gets(ctx context.Context, module string, keys ...string) (values map[string]*gvar.Var, err error) {

	values = make(map[string]*gvar.Var)

	if len(keys) == 0 {
		keysVar, err := dao.Config.Ctx(ctx).Where(dao.Config.Columns.Module, module).Array(dao.Config.Columns.Key)
		if err != nil {
			return nil, err
		}
		for _, keyVar := range keysVar {
			keys = append(keys, keyVar.String())
		}
	}
	for _, key := range keys {
		v, innerErr := c.Get(ctx, module, key)
		if innerErr != nil {
			return values, innerErr
		}
		values[key] = v
	}

	return values, nil

}

// 根据传递的module和key设置对应的配置值,存在则更新，不存在则插入
func (c *config) Set(ctx context.Context, module, key string, value interface{}) (err error) {
	_, err = gcache.Remove(ctx, model.CONFIG_CONFIGS_PREFIX+module+key)
	if err != nil {
		return
	}
	data := g.Map{
		dao.Config.Columns.Module: module,
		dao.Config.Columns.Key:    key,
		dao.Config.Columns.Value:  value,
	}
	_, err = dao.Config.Ctx(ctx).Save(data)
	return
}

// 根据传递的module和mapping批量设置对应的配置值,存在则更新，不存在则插入
func (c *config) Sets(ctx context.Context, module string, mapping map[string]interface{}) (err error) {
	if len(mapping) == 0 {
		return
	}
	return dao.Config.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for key, value := range mapping {
			err := c.Set(ctx, module, key, value)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
