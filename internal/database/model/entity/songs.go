// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Songs is the golang structure for table songs.
type Songs struct {
	Id            uint64      `json:"id"            description:"ID"`
	SongId        uint64      `json:"songId"        description:"歌曲 ID"`
	Title         string      `json:"title"         description:"歌曲标题"`
	Images        string      `json:"images"        description:"图片"`
	Author        string      `json:"author"        description:"作者"`
	Album         string      `json:"album"         description:"相册"`
	Description   string      `json:"description"   description:"描述"`
	PublishedDate *gtime.Time `json:"publishedDate" description:"发布时间"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
	ModifyTime    *gtime.Time `json:"modifyTime"    description:"修改时间"`
}
