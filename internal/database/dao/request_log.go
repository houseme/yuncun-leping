// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/houseme/yuncun-leping/internal/database/dao/internal"
)

// internalRequestLogDao is internal type for wrapping internal DAO implements.
type internalRequestLogDao = *internal.RequestLogDao

// requestLogDao is the data access object for table request_log.
// You can define custom methods on it to extend its functionality as you wish.
type requestLogDao struct {
	internalRequestLogDao
}

var (
	// RequestLog is globally public accessible object for table request_log operations.
	RequestLog = requestLogDao{
		internal.NewRequestLogDao(),
	}
)

// Fill with you ideas below.
