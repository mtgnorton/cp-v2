package controller

import (
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/index/internal/service"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var Ws = ws{}

type ws struct {
}

func (w *ws) Ws(r *ghttp.Request) {

	defer func() {
		if err := recover(); err != nil {
			g.Log("ws").Error(r.Context(), "err1111", err)
		}
	}()
	ws, err := r.WebSocket()
	if err != nil {
		g.Log("ws").Error(r.Context(), "err2222", err)
		r.Exit()
	}

	uniqueId := r.Cookie.Get("cp-v2-unique-id").String()
	token := r.Cookie.Get("cp-v2-token").String()

	g.Log("ws").Infof(r.Context(), "uniqueId 为%s的用户连接成功", uniqueId)

	user := &model.UserSummary{}

	wsUser := shared.NewWsUser(ws, uniqueId, user, uniqueId)

	if token != "" {
		user, err = w.getUserByToken(r)
		if err != nil {
			g.Log("ws").Error(r.Context(), "err3333", err)
		} else {
			wsUser = shared.NewWsUser(ws, token, user, token)
		}

	}

	shared.WsService.AddUser(wsUser)

	allCount, loginCount, usernames := shared.WsService.GetStatus()

	// 发送初始化消息
	wsUser.Write(&shared.WsMessage{
		Type: shared.WsMessageTypeInit,
		Data: g.Map{
			"allCount":   allCount,
			"loginCount": loginCount,
			"users":      usernames,
		},
	})

	// 新加入的用户进行广播
	shared.WsService.Broadcast(&shared.WsMessage{
		Type: shared.WsMessageTypeInc,
		Data: g.Map{
			"allCount":   allCount,
			"loginCount": loginCount,
			"username":   user.Username,
		},
	})

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			errMsg := err.Error()
			if errMsg == "websocket: close 1001 (going away)" || gstr.Contains(errMsg, "use of closed network connection") {

				shared.WsService.RemoveUser(uniqueId)

				allCount, loginCount, _ := shared.WsService.GetStatus()

				// 用户离开广播
				shared.WsService.Broadcast(&shared.WsMessage{
					Type: shared.WsMessageTypeDec,
					Data: g.Map{
						"allCount":   allCount,
						"loginCount": loginCount,
						"username":   user.Username,
					},
				})

				g.Log("ws").Infof(r.Context(), "uniqueId 为%s的用户断开连接", uniqueId)
			} else {
				g.Log("ws").Error(r.Context(), "err4444", err)
			}
			return
		}
		wm, err := shared.TransferWsMessage(msg)
		g.Log("ws").Infof(r.Context(), "uniqueId 为%s的用户发送消息为%#v", uniqueId, wm)
		if err != nil {
			g.Log("ws").Error(r.Context(), "err5555", err)

			wsUser.Write(&shared.WsMessage{
				Type:    "error",
				Message: err.Error(),
			})
			continue
		}

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

		wsUser.Write(wm)
	}
}

func (w *ws) getUserByToken(r *ghttp.Request) (user *model.UserSummary, err error) {
	g.Log("auth").Debug(r.Context(), "ws是否登录验证开始执行")

	customCtx := &model.Context{
		Data: make(g.Map),
	}
	shared.Context.Init(r, customCtx)

	user = &model.UserSummary{}

	err = service.FrontTokenInstance.LoadConfig().InitUser(r)

	if err != nil {
		return user, err
	}

	user, err = service.FrontTokenInstance.GetUser(r.Context())

	return
}
