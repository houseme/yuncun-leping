// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppInfo is the golang structure of table app_info for DAO operations like Where/Data.
type AppInfo struct {
	g.Meta          `orm:"table:app_info, do:true"`
	Id              interface{} // ID
	AppNo           interface{} // app 唯一标识
	AppName         interface{} // 应用名称
	AppId           interface{} // app id
	AppSecret       interface{} // 密钥
	AppKey          interface{} // 授权密钥
	WhitelistIp     interface{} // 白名单 IP
	WhitelistDomain interface{} // 白名单域名
	LimitAppKey     interface{} // 限速规则
	LimitAppId      interface{} // 限速 app id
	State           interface{} // 状态 0 默认 100 正常 200 禁用
	Remark          interface{} // 应用备注
	CreateTime      *gtime.Time // 创建时间
	ModifyTime      *gtime.Time // 修改时间
}
