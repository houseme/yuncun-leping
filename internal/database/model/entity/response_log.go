// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ResponseLog is the golang structure for table response_log.
type ResponseLog struct {
	Id                uint64      `json:"id"                description:"ID"`
	AppNo             uint64      `json:"appNo"             description:"应用 ID"`
	YearTime          uint        `json:"yearTime"          description:"年份"`
	MonthTime         uint        `json:"monthTime"         description:"月份"`
	DayTime           uint        `json:"dayTime"           description:"日期"`
	RequestId         uint64      `json:"requestId"         description:"对应的请求 ID"`
	RequestIp         string      `json:"requestIp"         description:"请求 IP"`
	RequestTime       *gtime.Time `json:"requestTime"       description:"请求时间"`
	ResponseTime      *gtime.Time `json:"responseTime"      description:"响应时间"`
	SongId            uint64      `json:"songId"            description:"歌曲 ID"`
	CommentId         uint64      `json:"commentId"         description:"评论 ID"`
	Title             string      `json:"title"             description:"歌曲标题"`
	CommentUserId     uint64      `json:"commentUserId"     description:"评论用户 ID"`
	CommentContent    string      `json:"commentContent"    description:"评论内容"`
	CommentLikedCount uint        `json:"commentLikedCount" description:"评论点赞数"`
	CounterValue      uint64      `json:"counterValue"      description:"访问统计值"`
	CreateTime        *gtime.Time `json:"createTime"        description:"创建时间"`
	ModifyTime        *gtime.Time `json:"modifyTime"        description:"修改时间"`
}
