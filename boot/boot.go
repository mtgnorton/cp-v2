package boot

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model/entity"
	_ "gf-admin/packed"
	"gf-admin/utility"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

var EnvName = ENV_LOCAL

const (
	ENV_DEV   = "dev"
	ENV_PROD  = "prod"
	ENV_LOCAL = "local"
)

func init() {
	ctx := context.TODO()
	loadConfigFile(ctx)
	newDefaultAdministrator(ctx)
	InitSensitiveWords(ctx)
}

// 根据环境变量或命令行参数家在对应的配置文件，命令行参数优先级高于环境变量,默认使用prod环境
func loadConfigFile(ctx context.Context) {
	// 获取环境变量
	tempEnvName := gcmd.GetOptWithEnv("v2.env").String()

	if tempEnvName != "" {
		EnvName = tempEnvName
	}
	switch EnvName {
	case ENV_DEV:
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-dev.toml")
	case ENV_PROD:
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-prod.toml")
	case ENV_LOCAL:
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config-local.toml")
	}
	g.Log().Infof(ctx, "从命令行或环境变量中读取的环境参数为：%s", tempEnvName)
	g.Log().Infof(ctx, "当前配置环境：%s", EnvName)
}

//系统启动时，判断是否存在管理员，如果不存在，则创建一个管理员
func newDefaultAdministrator(ctx context.Context) {
	administrator, err := dao.Administrator.Ctx(ctx).Where("username", "admin").One()
	if err != nil {
		g.Log().Fatalf(ctx, "query administrator error ,error msg following: %s", err)
	}
	if administrator.IsEmpty() {
		g.Log().Debug(ctx, "开始创建admin管理员")

		_, err = dao.Administrator.Ctx(ctx).Insert(entity.Administrator{
			Username: "admin",
			Password: utility.EncryptPassword("admin", "admin"),
			Nickname: "admin",
			Avatar:   "",
			Status:   "",
			Remark:   "",
		})
		if err != nil {
			g.Log().Fatalf(ctx, "init admin error,%s", err)
		}
	} else {
		g.Log().Debug(ctx, "admin管理员已经存在")
	}
}

func InitSensitiveWords(ctx context.Context) {
	wordVars, err := dao.SensitiveWords.Ctx(ctx).Array(dao.SensitiveWords.Columns().Word)
	if err != nil {
		g.Log().Fatalf(ctx, "load sensitive wordVars error,%s", err)
	}
	words := []string{}
	for _, word := range wordVars {
		words = append(words, word.String())
	}
	utility.SensitiveInspector.AddSensitiveWords(words)
	g.Log().Debugf(ctx, "敏感词加载完成,敏感词数量为:%d", len(words))
}
