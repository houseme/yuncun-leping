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

// HomeReq struct
type HomeReq struct {
	g.Meta `path:"/comment" tags:"Comment" method:"get" summary:"music comment"`
	*model.HomeInput
}

// HomeRes struct
type HomeRes struct {
	*model.HomeOutput
}

// RedirectReq redirect req
type RedirectReq struct {
	g.Meta `path:"/music/:songID/:contentType" tags:"Comment" method:"get" summary:"music redirect"`
	*model.RedirectInput
}

// RedirectRes redirect res
type RedirectRes struct {
	*model.RedirectOutput
}

// CounterReq counter req
type CounterReq struct {
	g.Meta `path:"/counter" tags:"Comment" method:"get" summary:"music counter"`
	*model.CounterInput
}

// CounterRes counter res
type CounterRes struct {
	*model.CounterOutput
}

// CompatibleReq compatible req
type CompatibleReq struct {
	g.Meta `path:"/" tags:"Comment" method:"get" summary:"music compatible"`
	*model.HomeInput
}

// CompatibleRes compatible res
type CompatibleRes struct {
	*model.HomeOutput
}
