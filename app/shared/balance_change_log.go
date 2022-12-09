package shared

import (
	"context"
	"gf-admin/app/dao"
	"gf-admin/app/model"
)

var BalanceChangeLog = balanceChangeLog{}

type balanceChangeLog struct {
}

// List 钱包变动日志列表
func (b *balanceChangeLog) List(ctx context.Context, in *model.BalanceChangeLogListInput) (out *model.BalanceChangeLogListOutput, err error) {
	out = &model.BalanceChangeLogListOutput{}
	d := dao.BalanceChangeLog.Ctx(ctx)
	c := dao.BalanceChangeLog.Columns()
	if in.UserId > 0 {
		d = d.Where(c.UserId, in.UserId)
	}
	if in.Username != "" {
		d = d.WhereLike(c.Username, "%"+in.Username+"%")
	}
	if in.Type != "" {
		d = d.Where(c.Type, in.Type)
	}
	out.Page = in.Page
	out.Size = in.Size
	out.Total, err = d.Count()

	if err != nil {
		return
	}
	d = d.OrderDesc(dao.BalanceChangeLog.Columns().Id)
	err = d.Page(in.Page, in.Size).Scan(&out.List)

	if err != nil {
		return
	}
	// 将type映射为可读的文字
	for _, item := range out.List {
		item.TypeShow = b.GetTypeShow(model.BalanceChangeType(item.Type))
	}
	return
}

func (b *balanceChangeLog) GetTypeShow(t model.BalanceChangeType) (show string) {
	switch t {
	case model.BALANCE_CHANGE_TYPE_LOGIN:
		show = "登录奖励"
	case model.BALANCE_CHANGE_TYPE_REGSITER:
		show = "注册奖励"
	case model.BALANCE_CHANGE_TYPE_ESTABLISH_POST_DEDUCT:
		show = "发布主题扣除"
	case model.BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_DEDUCT:
		show = "发布回复扣除"
	case model.BALANCE_CHANGE_TYPE_ESTABLISH_REPLY_REWARD:
		show = "发布回复奖励创建主题者"
	case model.BALANCE_CHANGE_TYPE_THANK_REPLY_DEDUCT:
		show = "感谢回复扣除"
	case model.BALANCE_CHANGE_TYPE_THANK_REPLY_REWARD:
		show = "感谢回复奖励创建回复者"
	case model.BALANCE_CHANGE_TYPE_THANK_POST_DEDUCT:
		show = "感谢主题扣除"
	case model.BALANCE_CHANGE_TYPE_THANK_POST_REWARD:
		show = "感谢主题奖励创建主题者"
	case model.BALANCE_CHANGE_TYPE_ACTIVITY:
		show = "活跃度奖励"
	}
	return
}
