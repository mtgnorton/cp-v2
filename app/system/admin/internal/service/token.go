package service

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var AdminTokenInstance = adminTokenHandle{ //如果直接在这里获取配置，会出错，因为这里的执行顺序可能早于boot.go加载配置文件，所以改为使用LoadConfig方法
}

type adminTokenHandle struct {
	shared.TokenHandler
	administrator *model.AdministratorSummary
	loadConfig    bool
}

//加载配置文件里面token的相关配置,应该最先调用该方法
func (a *adminTokenHandle) LoadConfig() *adminTokenHandle {
	ctx := context.TODO()
	AdminTokenInstance.TokenHandler = shared.TokenHandler{
		CacheMode:  shared.CacheModeRedis,
		Timeout:    g.Cfg().MustGet(ctx, "token.timeout").Int(),
		MaxRefresh: g.Cfg().MustGet(ctx, "token.maxRefresh").Int(),
		CacheKey:   g.Cfg().MustGet(ctx, "token.cacheKey").String(),
		EncryptKey: g.Cfg().MustGet(ctx, "token.encryptKey").Bytes(),
		MultiLogin: g.Cfg().MustGet(ctx, "token.multiLogin").Bool(),
	}
	a.loadConfig = true
	return a
}

func (a *adminTokenHandle) GetAdministrator(ctx context.Context) (administrator *model.AdministratorSummary, err error) {

	if !a.loadConfig {
		a.LoadConfig()
	}

	if a.administrator != nil && a.administrator.Id != 0 {
		return a.administrator, nil
	}
	data := shared.Context.GetUser(ctx)
	administrator = &model.AdministratorSummary{}
	err = gconv.Scan(data, &administrator)
	a.administrator = administrator
	return
}

func (a *adminTokenHandle) GetAdministratorId(ctx context.Context) (administratorId uint, err error) {
	administrator, err := a.GetAdministrator(ctx)
	if err != nil {
		return 0, err
	}

	return administrator.Id, err
}
