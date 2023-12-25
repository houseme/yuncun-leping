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

package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/yuncun-leping/app/front/internal/model"
)

// AuthorizeReq .
type AuthorizeReq struct {
	g.Meta `path:"/authorize" tags:"Authorize" method:"Post" summary:"authorize token"`
	*model.AuthorizeInput
}

// AuthorizeRes .
type AuthorizeRes struct {
	*model.AuthorizeOutput
}
