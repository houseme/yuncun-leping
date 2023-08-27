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
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/houseme/yuncun-leping/app/front/internal/consts"
	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/internal/tracing"
)

type sMiddleware struct {
}

// init
func init() {

}

// Initializer is a middleware handler for ghttp.Request.
func (s *sMiddleware) Initializer(r *ghttp.Request) {
	r.SetCtxVar("logger", consts.DefaultLoggerName)
	r.Middleware.Next()
}

// ClientIP sets the client ip to the context.
func (s *sMiddleware) ClientIP(r *ghttp.Request) {
	r.SetParam("clientIP", r.GetClientIp())
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
	}
	g.Log(r.GetCtxVar("logger").String()).Info(r.GetCtx(), "status: ", r.Response.Status, "path: ", r.URL.Path, "msg: ", errStr)
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
