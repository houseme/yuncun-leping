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
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/houseme/yuncun-leping/app/schedule/internal/service"
)

var (
	// Main is the entry for schedule application.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, `cron job start`)
			if _, err = gcron.Add(ctx, "* * * * * *", func(ctx context.Context) {
				g.Log().Debug(ctx, `cron job running`)
			}); err != nil {
				return err
			}

			if _, err = gcron.AddSingleton(ctx, "0 1 0 * * *", func(ctx context.Context) {
				g.Log().Debug(ctx, `cron job top list running start`)
				if err = service.Comment().TopList(ctx); err != nil {
					g.Log().Errorf(ctx, `cron job top list running error: %+v`, err)
				}
				g.Log().Debug(ctx, `cron job top list running end`)
			}); err != nil {
				return err
			}

			// Register shutdown handler.
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Info(ctx, `cron job shutdown`)
			})
			// Block listening the shutdown signal.
			g.Listen()
			return
		},
	}
)
