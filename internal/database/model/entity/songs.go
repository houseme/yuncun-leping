// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Songs is the golang structure for table songs.
type Songs struct {
	Id            uint64      `json:"id"            description:""`
	SongId        uint64      `json:"songId"        description:""`
	Title         string      `json:"title"         description:""`
	Images        string      `json:"images"        description:""`
	Author        string      `json:"author"        description:""`
	Album         string      `json:"album"         description:""`
	Description   string      `json:"description"   description:""`
	PublishedDate *gtime.Time `json:"publishedDate" description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
}
