// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppInfo is the golang structure for table app_info.
type AppInfo struct {
	Id              uint        `json:"id"              description:"ID"`
	AppNo           uint        `json:"appNo"           description:"app 唯一标识"`
	AppName         string      `json:"appName"         description:"应用名称"`
	AppId           string      `json:"appId"           description:"app id"`
	AppSecret       string      `json:"appSecret"       description:"密钥"`
	AppKey          string      `json:"appKey"          description:"授权密钥"`
	WhitelistIp     string      `json:"whitelistIp"     description:"白名单 IP"`
	WhitelistDomain string      `json:"whitelistDomain" description:"白名单域名"`
	LimitAppKey     uint        `json:"limitAppKey"     description:"限速规则"`
	LimitAppId      uint        `json:"limitAppId"      description:"限速 app id"`
	State           uint        `json:"state"           description:"状态 0 默认 100 正常 200 禁用"`
	Remark          string      `json:"remark"          description:"应用备注"`
	CreateTime      *gtime.Time `json:"createTime"      description:"创建时间"`
	ModifyTime      *gtime.Time `json:"modifyTime"      description:"修改时间"`
}
