// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RequestLogDao is the data access object for table request_log.
type RequestLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns RequestLogColumns // columns contains all the column names of Table for convenient usage.
}

// RequestLogColumns defines and stores column names for table request_log.
type RequestLogColumns struct {
	Id          string // ID
	AppNo       string // 应用 ID
	YearTime    string // 年份
	MonthTime   string // 月份
	DayTime     string // 日期
	RequestIp   string // 请求 IP
	RequestTime string // 请求时间
	UserAgent   string // 请求 user_agent
	Referer     string // referer
	Origin      string // origin
	Path        string // 链接域名之后的路径
	RequestUri  string // 请求的链接（路径和参数）
	CreateTime  string // 创建时间
	ModifyTime  string // 修改时间
}

// requestLogColumns holds the columns for table request_log.
var requestLogColumns = RequestLogColumns{
	Id:          "id",
	AppNo:       "app_no",
	YearTime:    "year_time",
	MonthTime:   "month_time",
	DayTime:     "day_time",
	RequestIp:   "request_ip",
	RequestTime: "request_time",
	UserAgent:   "user_agent",
	Referer:     "referer",
	Origin:      "origin",
	Path:        "path",
	RequestUri:  "request_uri",
	CreateTime:  "create_time",
	ModifyTime:  "modify_time",
}

// NewRequestLogDao creates and returns a new DAO object for table data access.
func NewRequestLogDao() *RequestLogDao {
	return &RequestLogDao{
		group:   "default",
		table:   "request_log",
		columns: requestLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RequestLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RequestLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RequestLogDao) Columns() RequestLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RequestLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RequestLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RequestLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
