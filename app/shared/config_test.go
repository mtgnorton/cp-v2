package shared

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/stretchr/testify/assert"
)

// sqlite 不支持save操作, 所以这里不测试
//func TestConfig_Sets(t *testing.T) {
//	err := Config.Sets(ctx, "global_sets", g.Map{
//		"name": "gf",
//		"age":  18,
//	})
//	assert.Nil(t, err)
//}

// sqlite 不支持save操作, 所以这里不测试
//func TestConfig_Set(t *testing.T) {
//
//	err := Config.Set(ctx, "global", "k1", "v1")
//	assert.Nil(t, err)
//
//	err = Config.Set(ctx, "global", "k2", g.Map{
//		"k2": "v2",
//	})
//	assert.Nil(t, err)
//	err = Config.Set(ctx, "global", "k3", g.Map{
//		"k3": g.Map{
//			"k4": "v4",
//		},
//	})
//	assert.Nil(t, err)
//	err = Config.Set(ctx, "global", "k5", g.Slice{1, 2, 3, 4})
//	assert.Nil(t, err)
//
//	err = Config.Set(ctx, "global", "k6", g.Slice{
//		g.Map{
//			"k6": "v6",
//		},
//		g.Map{"k7": "v7"},
//	})
//	assert.Nil(t, err)
//}

func TestConfig_Get(t *testing.T) {

	v, err := Config.Get(ctx, "global", "k1")
	assert.Nil(t, err)
	g.DumpWithType(v)

	v, err = Config.Get(ctx, "global", "k2")
	assert.Nil(t, err)
	g.DumpWithType(v.MapStrStr())

	v, err = Config.Get(ctx, "global", "k3")
	assert.Nil(t, err)
	g.DumpWithType(v.Slice())

}

func TestConfig_Gets(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		values, err := Config.Gets(ctx, "global", "k1", "k2")
		assert.Nil(t, err)
		g.DumpWithType(values)
	})
	t.Run("module_is_empty", func(t *testing.T) {
		values, err := Config.Gets(ctx, "", "k1", "k2")
		assert.Nil(t, err)
		g.DumpWithType(values)
	})

	t.Run("module_and_key_is_empty", func(t *testing.T) {
		values, err := Config.Gets(ctx, "")
		assert.Nil(t, err)
		g.DumpWithType(values)
	})
	t.Run("json", func(t *testing.T) {
		values, err := Config.Gets(ctx, "", "k3")
		assert.Nil(t, err)
		g.DumpWithType(values)
	})

}
