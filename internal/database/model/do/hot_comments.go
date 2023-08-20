// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HotComments is the golang structure of table hot_comments for DAO operations like Where/Data.
type HotComments struct {
	g.Meta        `orm:"table:hot_comments, do:true"`
	Id            interface{} // ID
	SongId        interface{} // 歌曲 ID
	UserId        interface{} // 用户 ID
	Nickname      interface{} // 昵称
	AvatarUrl     interface{} // 头像地址
	CommentId     interface{} // 评论 ID
	LikedCount    interface{} // 喜欢数
	Content       interface{} // 内容
	PublishedDate *gtime.Time // 发布时间
	CheckoutDate  *gtime.Time // 获取时间
	CreateTime    *gtime.Time // 创建时间
	ModifyTime    *gtime.Time // 更新时间
}
