// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Initializer is a middleware handler for ghttp.Request.
		Initializer(r *ghttp.Request)
		// ClientIP sets the client ip to the context.
		ClientIP(r *ghttp.Request)
		// AuthLogin sets the auth info to the context.
		AuthLogin(r *ghttp.Request)
		// Logger Middleware Log
		Logger(r *ghttp.Request)
		// HandlerResponse is a middleware handler for ghttp.Request.
		HandlerResponse(r *ghttp.Request)
		// AuthorizationForAPIKey is a middleware handler for ghttp.Request.
		AuthorizationForAPIKey(r *ghttp.Request)
		// AuthorizationForPassword is a middleware handler for ghttp.Request.
		AuthorizationForPassword(r *ghttp.Request)
		// RequestLog is a middleware handler for ghttp.Request.
		RequestLog(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
