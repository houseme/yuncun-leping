// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HotComments is the golang structure for table hot_comments.
type HotComments struct {
	Id            uint64      `json:"id"            description:"ID"`
	SongId        uint64      `json:"songId"        description:"歌曲 ID"`
	UserId        uint64      `json:"userId"        description:"用户 ID"`
	Nickname      string      `json:"nickname"      description:"昵称"`
	AvatarUrl     string      `json:"avatarUrl"     description:"头像地址"`
	CommentId     uint64      `json:"commentId"     description:"评论 ID"`
	LikedCount    uint        `json:"likedCount"    description:"喜欢数"`
	Content       string      `json:"content"       description:"内容"`
	PublishedDate *gtime.Time `json:"publishedDate" description:"发布时间"`
	CheckoutDate  *gtime.Time `json:"checkoutDate"  description:"获取时间"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
	ModifyTime    *gtime.Time `json:"modifyTime"    description:"更新时间"`
}
