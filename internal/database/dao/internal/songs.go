// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SongsDao is the data access object for table songs.
type SongsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns SongsColumns // columns contains all the column names of Table for convenient usage.
}

// SongsColumns defines and stores column names for table songs.
type SongsColumns struct {
	Id            string //
	SongId        string //
	Title         string //
	Images        string //
	Author        string //
	Album         string //
	Description   string //
	PublishedDate string //
	CreatedAt     string //
	UpdatedAt     string //
}

// songsColumns holds the columns for table songs.
var songsColumns = SongsColumns{
	Id:            "id",
	SongId:        "song_id",
	Title:         "title",
	Images:        "images",
	Author:        "author",
	Album:         "album",
	Description:   "description",
	PublishedDate: "published_date",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSongsDao creates and returns a new DAO object for table data access.
func NewSongsDao() *SongsDao {
	return &SongsDao{
		group:   "default",
		table:   "songs",
		columns: songsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SongsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SongsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SongsDao) Columns() SongsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SongsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SongsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SongsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
