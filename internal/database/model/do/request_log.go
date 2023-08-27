// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RequestLog is the golang structure of table request_log for DAO operations like Where/Data.
type RequestLog struct {
	g.Meta      `orm:"table:request_log, do:true"`
	Id          interface{} // ID
	AppNo       interface{} // 应用 ID
	YearTime    interface{} // 年份
	MonthTime   interface{} // 月份
	DayTime     interface{} // 日期
	RequestIp   interface{} // 请求 IP
	RequestTime *gtime.Time // 请求时间
	UserAgent   interface{} // 请求 user_agent
	Referer     interface{} // referer
	Origin      interface{} // origin
	Path        interface{} // 链接域名之后的路径
	RequestUri  interface{} // 请求的链接（路径和参数）
	CreateTime  *gtime.Time // 创建时间
	ModifyTime  *gtime.Time // 修改时间
}
