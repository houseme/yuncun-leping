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

package authorize

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
	"github.com/houseme/yuncun-leping/app/front/internal/service"
	"github.com/houseme/yuncun-leping/utility/helper"
)

type sAuthorize struct {
}

// init .
func init() {
	service.RegisterAuthorize(&sAuthorize{})
}

// Authorization app authorization
func (s *sAuthorize) Authorization(ctx context.Context, in *model.AuthorizeInput) (out *model.AuthorizeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-authorize-Authorization")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
	)
	logger.Debug(ctx, "authorize logic start params:", in)

	return
}
