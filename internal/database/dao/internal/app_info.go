// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppInfoDao is the data access object for table app_info.
type AppInfoDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AppInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AppInfoColumns defines and stores column names for table app_info.
type AppInfoColumns struct {
	Id              string // ID
	AppNo           string // app 唯一标识
	AppName         string // 应用名称
	AppId           string // app id
	AppSecret       string // 密钥
	AppKey          string // 授权密钥
	WhitelistIp     string // 白名单 IP
	WhitelistDomain string // 白名单域名
	LimitAppKey     string // 限速规则
	LimitAppId      string // 限速 app id
	State           string // 状态 0 默认 100 正常 200 禁用
	Remark          string // 应用备注
	CreateTime      string // 创建时间
	ModifyTime      string // 修改时间
}

// appInfoColumns holds the columns for table app_info.
var appInfoColumns = AppInfoColumns{
	Id:              "id",
	AppNo:           "app_no",
	AppName:         "app_name",
	AppId:           "app_id",
	AppSecret:       "app_secret",
	AppKey:          "app_key",
	WhitelistIp:     "whitelist_ip",
	WhitelistDomain: "whitelist_domain",
	LimitAppKey:     "limit_app_key",
	LimitAppId:      "limit_app_id",
	State:           "state",
	Remark:          "remark",
	CreateTime:      "create_time",
	ModifyTime:      "modify_time",
}

// NewAppInfoDao creates and returns a new DAO object for table data access.
func NewAppInfoDao() *AppInfoDao {
	return &AppInfoDao{
		group:   "default",
		table:   "app_info",
		columns: appInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppInfoDao) Columns() AppInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
