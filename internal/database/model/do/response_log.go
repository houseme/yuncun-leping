// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ResponseLog is the golang structure of table response_log for DAO operations like Where/Data.
type ResponseLog struct {
	g.Meta            `orm:"table:response_log, do:true"`
	Id                interface{} // ID
	AppNo             interface{} // 应用 ID
	YearTime          interface{} // 年份
	MonthTime         interface{} // 月份
	DayTime           interface{} // 日期
	RequestIp         interface{} // 请求 IP
	RequestTime       *gtime.Time // 请求时间
	ResponseTime      *gtime.Time // 响应时间
	SongId            interface{} // 歌曲 ID
	CommentId         interface{} // 评论 ID
	Title             interface{} // 歌曲标题
	CommentUserId     interface{} // 评论用户 ID
	CommentContent    interface{} // 评论内容
	CommentLikedCount interface{} // 评论点赞数
	CounterValue      interface{} // 访问统计值
	CreateTime        *gtime.Time // 创建时间
	ModifyTime        *gtime.Time // 修改时间
}
