// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HotCommentsDao is the data access object for table hot_comments.
type HotCommentsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns HotCommentsColumns // columns contains all the column names of Table for convenient usage.
}

// HotCommentsColumns defines and stores column names for table hot_comments.
type HotCommentsColumns struct {
	Id            string // ID
	SongId        string // 歌曲 ID
	UserId        string // 用户 ID
	Nickname      string // 昵称
	AvatarUrl     string // 头像地址
	CommentId     string // 评论 ID
	LikedCount    string // 喜欢数
	Content       string // 内容
	PublishedDate string // 发布时间
	CheckoutDate  string // 获取时间
	CreateTime    string // 创建时间
	ModifyTime    string // 更新时间
}

// hotCommentsColumns holds the columns for table hot_comments.
var hotCommentsColumns = HotCommentsColumns{
	Id:            "id",
	SongId:        "song_id",
	UserId:        "user_id",
	Nickname:      "nickname",
	AvatarUrl:     "avatar_url",
	CommentId:     "comment_id",
	LikedCount:    "liked_count",
	Content:       "content",
	PublishedDate: "published_date",
	CheckoutDate:  "checkout_date",
	CreateTime:    "create_time",
	ModifyTime:    "modify_time",
}

// NewHotCommentsDao creates and returns a new DAO object for table data access.
func NewHotCommentsDao() *HotCommentsDao {
	return &HotCommentsDao{
		group:   "default",
		table:   "hot_comments",
		columns: hotCommentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HotCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HotCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HotCommentsDao) Columns() HotCommentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HotCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HotCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HotCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
