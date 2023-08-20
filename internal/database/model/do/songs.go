// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Songs is the golang structure of table songs for DAO operations like Where/Data.
type Songs struct {
	g.Meta        `orm:"table:songs, do:true"`
	Id            interface{} //
	SongId        interface{} //
	Title         interface{} //
	Images        interface{} //
	Author        interface{} //
	Album         interface{} //
	Description   interface{} //
	PublishedDate *gtime.Time //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
