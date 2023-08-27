// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RequestLog is the golang structure for table request_log.
type RequestLog struct {
	Id          uint64      `json:"id"          description:"ID"`
	AppNo       uint64      `json:"appNo"       description:"应用 ID"`
	YearTime    uint        `json:"yearTime"    description:"年份"`
	MonthTime   uint        `json:"monthTime"   description:"月份"`
	DayTime     uint        `json:"dayTime"     description:"日期"`
	RequestIp   string      `json:"requestIp"   description:"请求 IP"`
	RequestTime *gtime.Time `json:"requestTime" description:"请求时间"`
	UserAgent   string      `json:"userAgent"   description:"请求 user_agent"`
	Referer     string      `json:"referer"     description:"referer"`
	Origin      string      `json:"origin"      description:"origin"`
	Path        string      `json:"path"        description:"链接域名之后的路径"`
	RequestUri  string      `json:"requestUri"  description:"请求的链接（路径和参数）"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	ModifyTime  *gtime.Time `json:"modifyTime"  description:"修改时间"`
}
