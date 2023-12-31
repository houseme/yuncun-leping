// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/yuncun-leping/internal/database/dao/internal"
)

// internalAppInfoDao is internal type for wrapping internal DAO implements.
type internalAppInfoDao = *internal.AppInfoDao

// appInfoDao is the data access object for table app_info.
// You can define custom methods on it to extend its functionality as you wish.
type appInfoDao struct {
	internalAppInfoDao
}

var (
	// AppInfo is globally public accessible object for table app_info operations.
	AppInfo = appInfoDao{
		internal.NewAppInfoDao(),
	}
)

// Fill with you ideas below.
