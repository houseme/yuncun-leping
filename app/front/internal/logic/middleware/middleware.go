// Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// You can obtain one at https://github.com/houseme/yuncun-leping.

// Package middleware for middleware.
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/houseme/yuncun-leping/app/front/internal/consts"
	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/internal/tracing"
	"github.com/houseme/yuncun-leping/utility/cache"
	"github.com/houseme/yuncun-leping/utility/helper"
)

type sMiddleware struct {
}

// init
func init() {
	service.RegisterMiddleware(&sMiddleware{})
}

// Initializer is a middleware handler for ghttp.Request.
func (s *sMiddleware) Initializer(r *ghttp.Request) {
	r.SetCtxVar("logger", consts.DefaultLogger)
	r.Middleware.Next()
}

// ClientIP sets the client ip to the context.
func (s *sMiddleware) ClientIP(r *ghttp.Request) {
	r.SetParam("clientIP", r.GetClientIp())
	r.SetParam("userAgent", r.UserAgent())
	r.SetParam("header", r.Header)
	r.Middleware.Next()
}

// AuthLogin sets the auth info to the context.
func (s *sMiddleware) AuthLogin(r *ghttp.Request) {
	r.SetParam("userAgent", r.UserAgent())
	r.SetParam("header", r.Header)
	r.Middleware.Next()
}

// Logger Middleware Log
func (s *sMiddleware) Logger(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := "success"
	if err := r.GetError(); err != nil {
		errStr = err.Error()
		g.Log(r.GetCtxVar("logger").String()).Errorf(r.GetCtx(), "Server logger Error:%+v", err)
	}
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "status: ", r.Response.Status, "path: ", r.URL.Path, "msg: ", errStr)
}

// HandlerResponse is a middleware handler for ghttp.Request.
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits the current handler.
	if r.Response.BufferLength() > 0 {
		g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "response buffer length: ", r.Response.BufferLength())
		return
	}

	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-service-middleware-HandlerResponse")
	r.SetCtx(ctx)
	defer span.End()

	// 设置公共参数
	tracing.SetAttributes(r, span)

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			case http.StatusBadGateway:
				code = gcode.CodeInvalidRequest
			case http.StatusInternalServerError:
				code = gcode.CodeNotSupported
			default:
				code = gcode.CodeUnknown
			}
		} else {
			code = gcode.New(200, "success", nil)
			msg = code.Message()
			g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "HandlerResponse body res:", res)
		}
	}
	r.Response.WriteJson(&model.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Time:    gtime.TimestampMicro(),
		TraceID: span.SpanContext().TraceID().String(),
	})
}

// middlewareResponse intercept the response
func (s *sMiddleware) middlewareResponse(r *ghttp.Request, span *gtrace.Span, resp *model.DefaultHandlerResponse) {
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "middlewareResponse body resp:", resp)
	// 设置公共参数
	tracing.SetAttributes(r, span)
	r.Response.WriteJson(resp)
}

// AuthorizationForAPIKey is a middleware handler for ghttp.Request.
func (s *sMiddleware) AuthorizationForAPIKey(r *ghttp.Request) {
	if s.authorization(r, consts.AuthTypeDefault) {
		r.Middleware.Next()
	}
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "AuthorizationForAPIKey authorization failed")
}

// AuthorizationForPassword is a middleware handler for ghttp.Request.
func (s *sMiddleware) AuthorizationForPassword(r *ghttp.Request) {
	if s.authorization(r, consts.AuthTypeAppKey) {
		g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "AuthorizationForPassword authorization success")
		r.Middleware.Next()
	}
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "AuthorizationForPassword authorization failed")
}

// authorization is a middleware handler for ghttp.Request.
func (s *sMiddleware) authorization(r *ghttp.Request, authType uint) bool {
	ctx, span := gtrace.NewSpan(r.Context(), "tracing-console-service-middleware-authorization")
	r.SetCtx(ctx)
	defer span.End()

	var (
		authHeader = gstr.Trim(r.GetHeader(consts.AuthorizationHeaderKey))
		logger     = g.Log(r.GetCtxVar("logger").String())
		resp       = &model.DefaultHandlerResponse{
			Code:    http.StatusMovedPermanently,
			Message: http.StatusText(http.StatusMovedPermanently),
			Data:    nil,
			Time:    gtime.TimestampMicro(),
			TraceID: span.SpanContext().TraceID().String(),
		}
	)
	logger.Debug(r.GetCtx(), "authorization authHeader:", authHeader)
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		resp.Message = "Invalid authorization Header"
		s.middlewareResponse(r, span, resp)
		return false
	}

	if fields[0] != consts.AuthorizationTypeBearer {
		resp.Message = "Unsupported authorization Type"
		s.middlewareResponse(r, span, resp)
		return false
	}

	var res, err = validateToken(r.GetCtx(), fields[1], authType)
	if err != nil {
		logger.Error(r.GetCtx(), "authorization failed: ", err)
		resp.Message = "authorization failed reason: " + err.Error()
		s.middlewareResponse(r, span, resp)
		return false
	}

	if res == nil {
		logger.Debug(r.GetCtx(), "authorization failed")
		resp.Message = "authorization failed reason"
		s.middlewareResponse(r, span, resp)
		return false
	}

	if res.AuthToken != fields[1] {
		logger.Debug(r.GetCtx(), "authorization token is Refresh new token:", res.AuthToken)
		resp.Code = http.StatusFound
		resp.Message = "token is Refresh"
		resp.Data = g.Map{
			"token": res.AuthToken,
		}
		s.middlewareResponse(r, span, resp)
		return false
	}

	r.SetParam("authAppNo", res.AuthAppNo)
	r.SetParam("authType", res.AuthType)

	r.SetCtxVar("authAppNo", res.AuthAppNo)
	r.SetCtxVar("authType", res.AuthType)

	logger.Debug(r.GetCtx(), "authorization success")
	return true
}

// validateToken is a middleware handler for ghttp.Request.
func validateToken(ctx context.Context, token string, authType uint) (authToken *model.AuthorizationToken, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-console-service-middleware-validateToken")
	defer span.End()

	var (
		redisKey       = cache.CenterAccessTokenKey(ctx, token)
		conn           gredis.Conn
		isAuthPassword = false
	)

	if conn, err = g.Redis(cache.CenterAccessTokenConn(ctx)).Conn(ctx); err != nil {
		return nil, gerror.Wrap(err, "validateToken Redis conn failed")
	}
	defer func() {
		_ = conn.Close(ctx)
	}()
	if authType != consts.AuthTypeDefault {
		isAuthPassword = true
		redisKey = cache.CenterAuthorizationKey(ctx, token)
	}
	var val *gvar.Var
	if val, err = conn.Do(ctx, "GET", redisKey); err != nil {
		return nil, gerror.Wrap(err, "validateToken Redis get failed(001)")
	}

	if val.IsNil() || val.IsEmpty() {
		return nil, gerror.New("validateToken auth token not found")
	}

	if err = val.Scan(&authToken); err != nil {
		return nil, gerror.Wrap(err, "validateToken Redis scan failed")
	}

	if authToken == nil {
		return nil, gerror.New("validateToken Redis get failed(002)")
	}

	var (
		authTime = authToken.AuthTime
		now      = gtime.Now()
	)
	// 刷新 token 过期时间
	if isAuthPassword {
		if now.Unix()-consts.RefreshTokenExpireTime > authTime {
			return nil, gerror.New("validateToken auth token expired")
		}
		logger := g.Log(helper.Helper().Logger(ctx))
		authToken.AuthTime = now.Unix()
		if now.Unix()-consts.PasswordExpireTime > authTime {
			logger.Debug(ctx, "validateToken auth token password expired 2 hours")
			if token, err = helper.Helper().CreateAccessToken(ctx, authToken.AuthAppNo); err != nil {
				return nil, gerror.Wrap(err, "validateToken CreateAccessToken failed")
			}
			authToken.AuthToken = token
			redisKey = cache.CenterAuthorizationKey(ctx, token)
		}
		logger.Debug(ctx, "validateToken auth token authTime:", authTime, "now:", now.Unix(), " authToken:", authToken)
		if val, err = conn.Do(ctx, "SETEX", redisKey, consts.TokenExpireTime, authToken); err != nil {
			return nil, gerror.Wrap(err, "validateToken Redis set failed")
		}
		logger.Debug(ctx, "validateToken auth token set Redis value:", val)
		return authToken, nil
	}
	if now.Unix()-consts.APIKeyExpireTime > authTime {
		return nil, gerror.New("validateToken auth token expired")
	}

	return authToken, nil
}

// RequestLog is a middleware handler for ghttp.Request.
func (s *sMiddleware) RequestLog(r *ghttp.Request) {
	r.SetParam("referer", r.Referer())
	r.SetParam("path", r.URL.Path)
	r.SetParam("requestURI", r.RequestURI)
	r.SetParam("origin", r.GetHeader("Origin"))
	r.Middleware.Next()
}
