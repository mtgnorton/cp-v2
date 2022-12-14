// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceChangeLogDao is the data access object for table forum_balance_change_log.
type BalanceChangeLogDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns BalanceChangeLogColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceChangeLogColumns defines and stores column names for table forum_balance_change_log.
type BalanceChangeLogColumns struct {
	Id         string //
	UserId     string // 用户id
	Username   string // 用户名
	Type       string // 每日登录奖励:login, 每日活跃度奖励: activity, 感谢主题: thanks_posts,感谢回复: thanks_relpy,创建主题: create_posts,创建回复: create_reply,初始奖励: register
	Amount     string // 金额
	Before     string // 变动前余额
	After      string // 变动后余额
	RelationId string // 关联主题id或关联回复id
	Remark     string // 备注
	CreatedAt  string // 创建时间
}

//  BalanceChangeLogColumns holds the columns for table forum_balance_change_log.
var balanceChangeLogColumns = BalanceChangeLogColumns{
	Id:         "id",
	UserId:     "user_id",
	Username:   "username",
	Type:       "type",
	Amount:     "amount",
	Before:     "before",
	After:      "after",
	RelationId: "relation_id",
	Remark:     "remark",
	CreatedAt:  "created_at",
}

// NewBalanceChangeLogDao creates and returns a new DAO object for table data access.
func NewBalanceChangeLogDao() *BalanceChangeLogDao {
	return &BalanceChangeLogDao{
		group:   "default",
		table:   "forum_balance_change_log",
		columns: balanceChangeLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceChangeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceChangeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceChangeLogDao) Columns() BalanceChangeLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceChangeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceChangeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceChangeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
