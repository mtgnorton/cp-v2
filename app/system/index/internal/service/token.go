package service

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var FrontTokenInstance = frontTokenHandle{ //如果直接在这里获取配置，会出错，因为这里的执行顺序可能早于boot.go加载配置文件，所以改为使用LoadConfig方法
}

type frontTokenHandle struct {
	shared.TokenHandler
	loadConfig bool
}

//加载配置文件里面token的相关配置,应该最先调用该方法
func (a *frontTokenHandle) LoadConfig() *frontTokenHandle {
	ctx := context.TODO()
	FrontTokenInstance.TokenHandler = shared.TokenHandler{
		CacheMode:  shared.CacheModeRedis,
		Timeout:    g.Cfg().MustGet(ctx, "front_token.timeout").Int(),
		MaxRefresh: g.Cfg().MustGet(ctx, "front_token.maxRefresh").Int(),
		CacheKey:   g.Cfg().MustGet(ctx, "front_token.cacheKey").String(),
		EncryptKey: g.Cfg().MustGet(ctx, "front_token.encryptKey").Bytes(),
		MultiLogin: g.Cfg().MustGet(ctx, "front_token.multiLogin").Bool(),
	}
	a.loadConfig = true
	return a
}

func (a *frontTokenHandle) GetUser(ctx context.Context) (user *model.UserSummary, err error) {

	if !a.loadConfig {
		a.LoadConfig()
	}
	data := shared.Context.GetUser(ctx)
	user = &model.UserSummary{}
	err = gconv.Scan(data, &user)
	return
}

func (a *frontTokenHandle) GetUserId(ctx context.Context) (userId uint, err error) {
	user, err := a.GetUser(ctx)
	if err != nil {
		return 0, err
	}

	return user.Id, err
}

func (a *frontTokenHandle) Remove(ctx context.Context, token string) (err error) {
	return a.TokenHandler.Remove(ctx, token)
}
