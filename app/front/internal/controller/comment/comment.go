/*
 * Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package comment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

// Controller for comment.
type Controller struct {
}

// New comment controller.
func New() *Controller {
	return &Controller{}
}

// Compatible .compatible with history interfaces
func (c *Controller) Compatible(r *ghttp.Request) {
	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-controller-comment-home")
	r.SetCtx(ctx)
	defer span.End()

	out, err := service.Comment().Home(r.GetCtx(), &model.CommentInput{
		ClientIP:   r.GetClientIp(),
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
		Path:       r.URL.Path,
		RequestURI: r.RequestURI,
		Header:     r.Header,
		AuthBase: &model.AuthBase{
			AuthAppNo: 0,
		},
	})

	if err != nil {
		g.Log().Error(r.GetCtx(), "controller comment home failed err:", err)
		r.Response.WriteStatusExit(503, "系统繁忙，请稍后重试")
	}
	g.Log().Debug(r.GetCtx(), "comment home logic success out:", out)
	r.Response.WriteJsonExit(out)
}

// Home for comment.
func (c *Controller) Home(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-home")
	defer span.End()

	res = &v1.Res{}
	if res.CommentOutput, err = service.Comment().Home(ctx, req.CommentInput); err != nil {
		return nil, err
	}
	return
}

// Redirect for comment.
func (c *Controller) Redirect(ctx context.Context, req *v1.RedirectReq) (res *v1.RedirectRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-Redirect")
	defer span.End()

	res = &v1.RedirectRes{}
	if res.RedirectOutput, err = service.Comment().Redirect(ctx, req.RedirectInput); err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(res.RedirectURL)
	return
}

// Counter for comment.
func (c *Controller) Counter(ctx context.Context, req *v1.CounterReq) (res *v1.CounterRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-Counter")
	defer span.End()

	res = &v1.CounterRes{}
	if res.CounterOutput, err = service.Comment().Counter(ctx, req.CounterInput); err != nil {
		return nil, err
	}
	return
}

//
