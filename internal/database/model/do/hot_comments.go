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
	Id            interface{} //
	SongId        interface{} //
	UserId        interface{} //
	Nickname      interface{} //
	AvatarUrl     interface{} //
	CommentId     interface{} //
	LikedCount    interface{} //
	Content       interface{} //
	PublishedDate *gtime.Time //
	CheckoutDate  *gtime.Time //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
