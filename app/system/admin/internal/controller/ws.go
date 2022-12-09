package controller

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/service"
	"gf-admin/utility/custom_log"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

var Ws = ws{}

type ws struct {
}

func (w *ws) Ws(r *ghttp.Request) {

	ws, err := r.WebSocket()
	if err != nil {
		g.Log("ws").Error(r.Context(), err)
		r.Exit()
	}

	userAgent := r.Header.Get("User-Agent")

	var administrator *model.AdministratorSummary

	if administrator, err = w.auth(r); err != nil {

		custom_log.Log(r, err)
		r.Exit()
	}

	wsUser := shared.NewWsUser(ws, gconv.String(administrator.Id), nil, userAgent)

	wsUser.SetUniqueId(gconv.String(administrator.Id))
	shared.WsService.AddUser(wsUser)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			custom_log.Log(r, err)
			return
		}
		wm, err := shared.TransferWsMessage(msg)

		wsUser.UpdateLastSendTime()

		if wm.Type == shared.WsMessageTypeHeart {
			wsUser.Write(&shared.WsMessage{
				Type: shared.WsMessageTypeHeart,
				Data: g.Map{
					"pong": gtime.Timestamp(),
				},
			})
			continue
		}

		if err != nil {
			wsUser.Write(&shared.WsMessage{
				Type:    "error",
				Message: err.Error(),
			})
			continue
		}
		wsUser.Write(wm)
	}
}

func (w *ws) auth(r *ghttp.Request) (administrator *model.AdministratorSummary, err error) {
	g.Log("auth").Debug(r.Context(), "ws是否登录验证开始执行")

	customCtx := &model.Context{
		Data: make(g.Map),
	}
	shared.Context.Init(r, customCtx)

	administrator = &model.AdministratorSummary{}

	err = service.AdminTokenInstance.LoadConfig().InitUser(r)

	if err != nil {
		return administrator, err
	}

	administrator, err = service.AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		return
	}
	if administrator.Id == 0 {
		return administrator, gerror.New("未登录或会话已过期，请您登录后再继续")
	}

	return
}

func (ws *ws) MonitorSystem(ctx context.Context) {
	gtimer.AddSingleton(ctx, 1*time.Second, func(ctx context.Context) {
		//
		//connAmount := shared.WsService.ConnCount()
		//if connAmount > 0 {
		//	memoryInfo, err := utility.GetMemoryInfo()
		//	if err != nil {
		//		return
		//	}
		//	cpuInfo, err := utility.GetCpuInfo()
		//	if err != nil {
		//		return
		//	}
		//
		//	shared.WsService.Broadcast(&shared.WsMessage{
		//		Type: shared.WsMessageTypeInit,
		//		Data: g.Map{
		//			"cpu":                 cpuInfo,
		//			"memory":              memoryInfo,
		//			"administratorAmount": shared.WsService.UserCount(),
		//		},
		//	})
		//}
	})
}
