package shared

import (
	"context"
	"gf-admin/app/model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
)

var WsService = wsService{
	wsUsersMap: make(map[string]map[string]*WsUser),
	timeout:    10,
}

type wsService struct {
	wsUsersMap map[string]map[string]*WsUser //第一维为用户id，第二维为设备号
	timeout    int64                         //超时时间，连接多久没有进行交互后移除
	sync.RWMutex
}

//判断连接是否仍然存活
func init() {
	gtimer.AddSingleton(gctx.New(), time.Second*5, func(ctx context.Context) {

		//fmt.Println("clear开始运行")
		for _, users := range WsService.wsUsersMap {
			for deviceNumber, user := range users {
				g.Log("ws").Debugf(ctx, "%#v", user)

				if gtime.Timestamp()-user.lastSendTime > WsService.timeout || user.close {
					WsService.RemoveUser(user.UniqueId, deviceNumber)
					g.Log("ws").Debugf(ctx, "移除unique id为：%s，deviceNumber为：%s的用户", user.UniqueId, deviceNumber)
				}
			}
		}
	})
}

// WsUser WebSocket 用户标示
type WsUser struct {
	UniqueId          string
	Conn              *ghttp.WebSocket
	writeMessagesChan chan interface{} //ws消息队列，当多个进程同时向同一个conn发送消息，并且消息量大时，可能会有并发问题，使用chan改为串行
	lastSendTime      int64            //最后一次推送消息的时间,超过10s移除用户
	close             bool             //清理失活用户连接有两种可能性，第一种通过定时任务主动检测到用户断开，可以进行主动清理，第二种没有检测到，此时用户的conn已经不可用，但是后端感知不到，在WriteLoop函数抛错后才能感知到,此时将用户标记为close，等待定时任务清理
	deviceNumber      string           //用户设备号
	data              interface{}      //附加数据,存储用户信息
}

// 广播信息
func (ws *wsService) Broadcast(message *WsMessage) {
	for _, users := range ws.wsUsersMap {
		for _, user := range users {
			user.Write(message)
		}
	}
}

func (ws *wsService) GetStatus() (allCount uint, loginCount uint, usernames []string) {
	usernames = make([]string, 0)
	for _, users := range ws.wsUsersMap {

		for _, user := range users {
			userSummary, ok := user.GetData().(*model.UserSummary)
			if !ok {
				break
			}
			if userSummary.Username != "" {
				usernames = append(usernames, userSummary.Username)
				loginCount++
				break
			}
		}
		allCount++
	}
	return
}

//获取所有用户
func (ws *wsService) GetUsers() map[string]map[string]*WsUser {
	return ws.wsUsersMap
}

//判断某个用户是否存在
func (ws *wsService) ExistUser(uniqueId string) bool {
	return len(ws.wsUsersMap[uniqueId]) > 0
}

// ExistUserDeviceNumber 判断某个用户的某个设备是否存在
func (ws *wsService) ExistUserDeviceNumber(uniqueId string, deviceNumber string) bool {
	if ws.wsUsersMap[uniqueId] == nil {
		return false
	}
	return ws.wsUsersMap[uniqueId][deviceNumber] != nil
}

//获取某个用户
func (ws *wsService) GetUser(uniqueId string) map[string]*WsUser {
	return ws.wsUsersMap[uniqueId]
}

// AddUser 添加用户
func (ws *wsService) AddUser(user *WsUser) {

	ws.Lock()
	defer ws.Unlock()
	uniqueId := user.UniqueId

	if len(ws.wsUsersMap[uniqueId]) == 0 {
		ws.wsUsersMap[uniqueId] = make(map[string]*WsUser)
	}

	g.Log().Infof(gctx.New(), "添加用户:%s,设备号:%s", user.UniqueId, user.deviceNumber)

	ws.wsUsersMap[uniqueId][user.deviceNumber] = user

	go user.WriteLoop()

}

//移除用户
func (ws *wsService) RemoveUser(uniqueId string, deviceNumbers ...string) bool {
	ws.Lock()
	defer ws.Unlock()

	users := ws.GetUser(uniqueId)

	if len(users) == 0 {
		return true
	}
	var deviceNumber string

	if len(deviceNumbers) > 0 {
		deviceNumber = deviceNumbers[0]
	}

	if deviceNumber != "" {
		if _, ok := users[deviceNumber]; ok {
			users[deviceNumber].Clear()
			delete(users, deviceNumber)

			if len(users) == 0 {
				delete(ws.wsUsersMap, uniqueId)
			}
		}
	} else {
		for _, user := range users {
			user.Clear()
			g.Log().Infof(gctx.New(), "删除用户:%s,设备号:%s", user.UniqueId, user.deviceNumber)
		}
		delete(ws.wsUsersMap, uniqueId)
	}

	return true

}

func NewWsUser(conn *ghttp.WebSocket, uniqueId string, data interface{}, deviceNumbers ...string) *WsUser {
	var deviceNumber string
	if len(deviceNumbers) > 0 {
		uniqueId = deviceNumbers[0]
	}
	user := &WsUser{
		UniqueId:          uniqueId,
		Conn:              conn,
		deviceNumber:      deviceNumber,
		writeMessagesChan: make(chan interface{}, 500),
		data:              data,
	}
	user.UpdateLastSendTime()
	return user
}

func (user *WsUser) SetUniqueId(id string) {
	user.UniqueId = id
}

func (user *WsUser) SetData(data interface{}) {
	user.data = data
}
func (user *WsUser) GetData() interface{} {
	return user.data
}
func (user *WsUser) UpdateLastSendTime() {
	user.lastSendTime = gtime.Timestamp()
}

func (user *WsUser) WriteLoop() {
	ctx := gctx.New()
	for msg := range user.writeMessagesChan {
		if user.close {
			return
		}

		g.Log("ws").Debugf(ctx, "发送消息:%s", gconv.String(msg))
		err := user.Conn.WriteJSON(msg)
		if err != nil {
			g.Log("ws").Debugf(gctx.New(), "用户id%v连接已关闭", user.UniqueId)
			user.close = true //等待定时任务进行清理
			return
		}
	}
}

func (user *WsUser) Write(data *WsMessage) {
	if user.close {
		return
	}
	user.writeMessagesChan <- data
}

func (user *WsUser) Read() (messageType int, p []byte, err error) {
	return user.Conn.ReadMessage()
}

func (user *WsUser) Clear() {
	defer func() {
		err := recover()

		if err != nil {
			g.Log("ws").Debugf(gctx.New(), "recover error,%s", err)
		}

	}()
	user.close = true

	err := user.Conn.Close()
	if err != nil {
		g.Log("ws").Debugf(gctx.New(), "close ws  conn error,%s", err)

	}

	close(user.writeMessagesChan)
}

const (
	WsMessageTypeHeart = "heart"
	WsMessageTypeInit  = "init"
	WsMessageTypeInc   = "inc"
	WsMessageTypeDec   = "dec"
)

//接受和发送的消息格式
type WsMessage struct {
	Type    string      `json:"type"`              // 可选值：heart(心跳包), system(系统信息),error(错误),debug(调试)
	Message string      `json:"message,omitempty"` // 提示信息
	Data    interface{} `json:"data,omitempty"`    // 返回数据(业务接口定义具体数据结构)
}

func TransferWsMessage(p []byte) (m *WsMessage, err error) {
	m = &WsMessage{}
	err = gconv.Scan(p, m)

	if err != nil {
		return m, gerror.New("输入消息转换失败")
	}
	return
}
