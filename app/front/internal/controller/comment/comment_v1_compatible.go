/*
 *  Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package comment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/yuncun-leping/app/front/api/comment/v1"
	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/utility/helper"
)

// Compatible .compatible with history interfaces
func (c *ControllerV1) Compatible(ctx context.Context, req *v1.CompatibleReq) (res *v1.CompatibleRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-comment-home")
	defer span.End()

	var (
		r      = g.RequestFromCtx(ctx)
		logger = g.Log(helper.Helper().Logger(ctx))
	)
	logger.Debug(r.GetCtx(), "comment home logic start params:", req)
	if req.ClientIP == "" {
		req.ClientIP = r.GetClientIp()
	}
	if req.UserAgent == "" {
		req.UserAgent = r.UserAgent()
	}
	if req.Referer == "" {
		req.Referer = r.Referer()
	}
	if req.Path == "" {
		req.Path = r.URL.Path
	}
	if req.RequestURI == "" {
		req.RequestURI = r.RequestURI
	}
	if req.Header == nil {
		req.Header = r.Header
	}
	req.AuthBase = &model.AuthBase{
		AuthAppNo: 0,
	}

	if res.HomeOutput, err = service.Comment().Home(ctx, req.HomeInput); err != nil {
		logger.Errorf(r.GetCtx(), "controller comment home failed err:%+v", err)
		r.Response.WriteStatusExit(503, "系统繁忙，请稍后重试")
	}
	logger.Debug(r.GetCtx(), "comment home logic success out:", res.HomeOutput)
	r.Response.WriteJsonExit(res.HomeOutput)
	return
}
