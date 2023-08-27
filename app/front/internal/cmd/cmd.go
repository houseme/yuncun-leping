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

package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/houseme/yuncun-leping/app/front/internal/controller/authorize"
	"github.com/houseme/yuncun-leping/app/front/internal/controller/comment"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
)

var (
	// Main is the entry for front application.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Initializer, service.Middleware().ClientIP, service.Middleware().Logger, service.Middleware().HandlerResponse)
				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					authorize.New(),
				)
				group.Group("/api.v1/front", func(group *ghttp.RouterGroup) {
					// group.Middleware(service.Middleware().AuthorizationForPassword)
					group.Bind(
						comment.New(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
