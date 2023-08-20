// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HotComments is the golang structure for table hot_comments.
type HotComments struct {
	Id            uint64      `json:"id"            description:""`
	SongId        uint64      `json:"songId"        description:""`
	UserId        uint64      `json:"userId"        description:""`
	Nickname      string      `json:"nickname"      description:""`
	AvatarUrl     string      `json:"avatarUrl"     description:""`
	CommentId     uint64      `json:"commentId"     description:""`
	LikedCount    uint        `json:"likedCount"    description:""`
	Content       string      `json:"content"       description:""`
	PublishedDate *gtime.Time `json:"publishedDate" description:""`
	CheckoutDate  *gtime.Time `json:"checkoutDate"  description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
}
